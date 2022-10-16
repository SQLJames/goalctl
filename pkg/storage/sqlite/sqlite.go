package sqlite

import (
	"database/sql"
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
		runMigrations(sqlDB)
		if errDB != nil {
			jlogr.Logger.ILog.Fatal(errDB, errDB.Error())
		}

		repo = &Repository{queries: *sqlc.New(sqlDB)}
	})

	return repo, errDB
}
