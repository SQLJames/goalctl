package sqlite

import (
	"database/sql"
	"errors"
	"sync"

	"fmt"

	// go-sqlite required for embedded sqlite database.
	_ "github.com/glebarez/go-sqlite"

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

var (
	ErrNoRows = errors.New("that item doesn't exist")
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

func newDB() (db *database, err error) {
	location, err := util.MakeApplicationFolder("database")
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

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
			return
		}

		sqlDB, errDB := sql.Open("sqlite", fmt.Sprintf("%s?_pragma=busy_timeout(%d000)&_pragma=journal_mode(WAL)", database.getDatabasePath(), timeoutInSeconds))
		if errDB != nil {
			return
		}
		_, errDB = runMigrations(sqlDB)
		if errDB != nil {
			return
		}

		repo = &Repository{queries: *sqlc.New(sqlDB)}
	})

	return repo, errDB
}
