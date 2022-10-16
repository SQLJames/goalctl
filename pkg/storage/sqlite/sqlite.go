package sqlite

import (
	"bytes"
	"database/sql"
	"embed"
	"errors"
	"io/fs"
	"strings"
	"sync"

	"fmt"

	// go-sqlite required for embedded sqlite database.
	_ "github.com/glebarez/go-sqlite"
	migrate "github.com/rubenv/sql-migrate"
	sqlmigrations "github.com/rubenv/sql-migrate"
	"github.com/sqljames/goalctl/pkg/info"
	"github.com/sqljames/goalctl/pkg/storage/sqlite/sqlc"
	"github.com/sqljames/goalctl/pkg/util"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

var (
	once  sync.Once
	repo  *Repository
	errDB error
)

const (
	timeoutInSeconds int    = 15
	delimiter        string = "||"
)

type database struct {
	Location  string
	Name      string
	Extension string
}

type Repository struct {
	queries sqlc.Queries
}

//go:embed sqlc/schema/*.sql
var schema embed.FS

func newDB() (db *database, err error) {
	location, err := util.MakeApplicationFolder("database")
	if err != nil {
		return nil, err
	}
	return &database{
		Location:  location,
		Name:      info.GetApplicationName(),
		Extension: "db",
	}, nil
}

// getDatabaseFileName gets just the file name and extension of the database.
func (db *database) getDatabaseFileName() (databaseFileName string) {
	return fmt.Sprintf("%s.%s", db.Name, db.Extension)
}

// getDatabasePath returns the full file path to the database file.
func (db *database) getDatabasePath() (databasePath string) {
	return util.JoinPath(db.Location, db.getDatabaseFileName())
}

func NewSQLiteStorage() (storage *Repository, errDB error) {
	once.Do(func() {
		database, errDB := newDB()
		if errDB != nil {
			jlogr.Logger.ILog.Fatal(errDB, errDB.Error())
		}

		sqlDB, errDB := sql.Open("sqlite", fmt.Sprintf("%s?_pragma=busy_timeout(%d000)&_pragma=journal_mode(WAL)", database.getDatabasePath(), timeoutInSeconds))
		if errDB != nil {
			jlogr.Logger.ILog.Fatal(errDB, errDB.Error())
		}
		numberOfMigrations, errDB := runMigrations(sqlDB)
		if errDB != nil {
			jlogr.Logger.ILog.Fatal(errDB, errDB.Error())
		}
		jlogr.Logger.ILog.Info("Applied migrations!", "numberOfMigrations", numberOfMigrations)
		repo = &Repository{queries: *sqlc.New(sqlDB)}
	})

	return repo, errDB
}

func runMigrations(sqlDB *sql.DB) (int, error) {
	MigrationSource := EmbedFileMigrationSource{Filesystem: schema}
	migration, err := MigrationSource.FindMigrations()
	if err != nil {
		return 0, err
	}

	return sqlmigrations.Exec(sqlDB, "sqlite3", &migrate.MemoryMigrationSource{Migrations: migration}, migrate.Up)
}

// EmbedFS interface for supporting embed.FS as injected filesystem and provide possibility to mock.
type EmbedFS interface {
	fs.ReadFileFS
}

// EmbedFileMigrationSource implements MigrationSource and provide migrations from a native embed filesystem.
type EmbedFileMigrationSource struct {
	Filesystem EmbedFS
}

var (
	ErrEmbedWalkFailed    = errors.New(`failed to walk recursive over embed source directory`)
	ErrEmbedReadDirFailed = errors.New(`directory read failed`)
)

// FindMigrations is part of MigrationSource implementation
func (f *EmbedFileMigrationSource) FindMigrations() ([]*sqlmigrations.Migration, error) {
	items := make([]*sqlmigrations.Migration, 0)
	err := fs.WalkDir(f.Filesystem, `.`, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf(`%v: %v`, ErrEmbedReadDirFailed, err)
		}
		// from now we always return nil cause if we send err, WalkDir stop processing
		if entry.IsDir() {
			return nil
		}
		if !strings.HasSuffix(entry.Name(), `.sql`) {
			return nil
		}
		content, err := f.Filesystem.ReadFile(path)
		if err != nil {
			return nil
		}
		migration, err := sqlmigrations.ParseMigration(entry.Name(), bytes.NewReader(content))
		if err != nil {
			return nil
		}
		items = append(items, migration)
		return nil
	})
	if err != nil {
		return items, fmt.Errorf(`%v: %v`, ErrEmbedWalkFailed, err)
	}
	return items, nil
}
