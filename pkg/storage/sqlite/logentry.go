package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/storage/sqlite/sqlc"
)

func (sl Repository) CreateLogEntry(ctx context.Context, arg *resources.LogEntry, notebookName string) (*resources.LogEntry, error) {
	NotebookID, err := sl.GetNotebookIDByName(ctx, notebookName)
	if NotebookID == 0 {
		return nil, fmt.Errorf("CreateLogEntry: notebook doesn't exist, please create it first. notebookName: %s", notebookName)
	}
	if err != nil {
		return nil, err
	}
	Entry, err := sl.queries.CreateLogEntry(ctx, sqlc.CreateLogEntryParams{
		Author: sql.NullString{
			String: arg.Author,
			Valid:  true,
		},
		Tags: sql.NullString{
			String: convertSliceToString(arg.Tags),
			Valid:  true,
		},
		Note:        arg.Entry,
		Createddate: arg.CreatedDate,
		Notebookid:  NotebookID,
	})
	if err != nil {
		return nil, err
	}
	arg.LogEntryID = Entry.Logentryid
	arg.Notebookid = Entry.Notebookid

	return arg, err
}

func (sl Repository) GetLogEntryByCreatedDate(ctx context.Context, createddate string) (logEntries *[]resources.LogEntry, err error) {
	sqlcLogEntries, err := sl.queries.GetLogEntryByCreatedDate(ctx, createddate)
	if err != nil {
		return nil, err
	}
	return convertSqlcLogEntriesToResource(&sqlcLogEntries), err
}

func (sl Repository) GetLogEntryByNotebook(ctx context.Context, name string) (*[]resources.LogEntry, error) {
	sqlcLogEntries, err := sl.queries.GetLogEntryByNotebook(ctx, name)
	if err != nil {
		return nil, err
	}
	return convertSqlcLogEntriesToResource(&sqlcLogEntries), err
}

func (sl Repository) GetLogEntryByLogEntryID(ctx context.Context, logentryid int64) (resources.LogEntry, error) {
	sqlcLogEntry, err := sl.queries.GetLogEntryByLogEntryID(ctx, logentryid)
	if err != nil {
		return resources.LogEntry{}, err
	}
	return convertSqlcLogEntryToResource(&sqlcLogEntry), err
}
