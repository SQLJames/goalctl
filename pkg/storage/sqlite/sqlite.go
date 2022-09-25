package sqlite

import (
	"context"
	"database/sql"
	// embed used for the dll statements on db creation
	_ "embed"
	"fmt"

	// go-sqlite required for embedded sqlite database
	_ "github.com/glebarez/go-sqlite"
	"github.com/sqljames/goalctl/pkg/info"
	"github.com/sqljames/goalctl/pkg/storage/sqlite/sqlc"
	"github.com/sqljames/goalctl/pkg/util"
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

//go:embed sqlc/schema/GoalToLogEntry.sql
var goalToLogEntryddl string

//go:embed sqlc/schema/Goal.sql
var goalddl string

//go:embed sqlc/schema/LogEntry.sql
var logentryddl string

//go:embed sqlc/schema/Notebook.sql
var notebookddl string

var ddls []string = []string{logentryddl, notebookddl, goalddl, goalToLogEntryddl}

func newDB() (db *database, err error) {
	location, err := util.MakeStorageLocation()
	if err != nil {
		return nil, err
	}
	return &database{
		Location:  location,
		Name:      info.GetApplicationName(),
		Extension: "db",
	}, nil
}

// getDatabaseFileName gets just the file name and extension of the database
func (db *database) getDatabaseFileName() (databaseFileName string) {
	return fmt.Sprintf("%s.%s", db.Name, db.Extension)
}

// getDatabasePath returns the full file path to the database file
func (db *database) getDatabasePath() (databasePath string) {
	return util.JoinPath(db.Location, db.getDatabaseFileName())
}

func NewSQLiteStorage() (storage *Repository, err error) {
	database, err := newDB()
	if err != nil {
		return nil, err
	}
	CreateSchema := !util.FileExists(database.getDatabasePath())

	sqlDB, err := sql.Open("sqlite", fmt.Sprintf("%s?_pragma=busy_timeout(%d000)&_pragma=journal_mode(WAL)", database.getDatabasePath(), timeoutInSeconds))
	if err != nil {
		return nil, err
	}
	if CreateSchema {
		// create tables
		for _, ddl := range ddls {
			if _, err := sqlDB.ExecContext(context.Background(), ddl); err != nil {
				return nil, err
			}
		}
	}

	return &Repository{queries: *sqlc.New(sqlDB)}, nil
}
