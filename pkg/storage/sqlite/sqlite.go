package sqlite

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"

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

type SQLiteStorage struct {
	queries sqlc.Queries
}

//go:embed sqlc/schema/GoalToLogEntry.sql
var GoalToLogEntryddl string

//go:embed sqlc/schema/Goal.sql
var goalddl string

//go:embed sqlc/schema/LogEntry.sql
var logentryddl string

//go:embed sqlc/schema/Notebook.sql
var notebookddl string

var ddls []string = []string{logentryddl, notebookddl, goalddl, GoalToLogEntryddl}

func newDB() (DB *database, err error) {
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
func (DB *database) getDatabaseFileName() (databaseFileName string) {
	return fmt.Sprintf("%s.%s", DB.Name, DB.Extension)
}

// getDatabasePath returns the full file path to the database file
func (DB *database) getDatabasePath() (databasePath string) {
	return util.JoinPath(DB.Location, DB.getDatabaseFileName())
}

func NewSQLiteStorage() (storage *SQLiteStorage, err error) {
	database, err := newDB()
	if err != nil {
		return nil, err
	}
	CreateSchema := !util.FileExists(database.getDatabasePath())

	db, err := sql.Open("sqlite", fmt.Sprintf("%s?_pragma=busy_timeout(%d000)&_pragma=journal_mode(WAL)", database.getDatabasePath(), timeoutInSeconds))
	if err != nil {
		return nil, err
	}
	if CreateSchema {
		// create tables
		for _, ddl := range ddls {
			if _, err := db.ExecContext(context.Background(), ddl); err != nil {
				return nil, err
			}
		}
	}

	return &SQLiteStorage{queries: *sqlc.New(db)}, nil
}
