package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/storage/sqlite/sqlc"
)

func (SL SQLiteStorage) CreateLogEntry(ctx context.Context, arg resources.LogEntry, notebookName string) (resources.LogEntry, error) {
	NotebookId, err := SL.GetNotebookIdByName(ctx, notebookName)
	if NotebookId == 0 {
		return resources.LogEntry{}, fmt.Errorf("notebook not created yet")
	}
	if err != nil {
		return resources.LogEntry{}, err
	}
	Entry, err := SL.queries.CreateLogEntry(ctx, sqlc.CreateLogEntryParams{
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
		Notebookid:  NotebookId,
	})
	if err != nil {
		return resources.LogEntry{}, err
	}
	arg.LogEntryID = Entry.Logentryid
	arg.Notebookid = Entry.Notebookid

	return arg, err
}

func (SL SQLiteStorage) GetLogEntryByCreatedDate(ctx context.Context, createddate string) (LogEntries []resources.LogEntry, err error) {
	sqlcLogEntries, err := SL.queries.GetLogEntryByCreatedDate(ctx, createddate)
	if err != nil {
		return nil, err
	}
	return convertSqlcLogEntriesToResource(sqlcLogEntries), err
}

func (SL SQLiteStorage) GetLogEntryByNotebook(ctx context.Context, name string) ([]resources.LogEntry, error) {
	sqlcLogEntries, err := SL.queries.GetLogEntryByNotebook(ctx, name)
	if err != nil {
		return nil, err
	}
	return convertSqlcLogEntriesToResource(sqlcLogEntries), err
}

func (SL SQLiteStorage) GetLogEntryByLogEntryID(ctx context.Context, logentryid int64) (resources.LogEntry, error) {
	sqlcLogEntry, err := SL.queries.GetLogEntryByLogEntryID(ctx, logentryid)
	if err != nil {
		return resources.LogEntry{}, err
	}
	return convertSqlcLogEntryToResource(sqlcLogEntry), err
}
