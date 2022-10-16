package sqlite

import (
	"context"
	"database/sql"

	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/storage/sqlite/sqlc"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func (sl Repository) GetLogEntries(ctx context.Context) (logEntries []*resources.LogEntry, err error) {
	sqlcLogEntries, err := sl.queries.GetLogEntries(ctx)
	if err != nil {
		return nil, err
	}

	return convertSqlcLogEntriesToResource(sqlcLogEntries), err
}

func (sl Repository) CreateLogEntry(ctx context.Context, arg *resources.LogEntry, notebookName string) (*resources.LogEntry, error) {
	NotebookID, err := sl.GetNotebookIDByName(ctx, notebookName)
	if err != nil {
		return nil, err
	}
	if NotebookID == 0 {
		jlogr.Logger.ILog.Debug("No results for notebook, Attempting to create a new notebook")

		notebook, err := sl.CreateNotebook(ctx, notebookName)
		if err != nil {
			return nil, err
		}
		NotebookID = notebook.Notebookid
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

func (sl Repository) GetLogEntryByCreatedDate(ctx context.Context, createddate string) (logEntries []*resources.LogEntry, err error) {
	sqlcLogEntries, err := sl.queries.GetLogEntryByCreatedDate(ctx, createddate)
	if err != nil {
		return nil, err
	}

	return convertSqlcLogEntriesToResource(sqlcLogEntries), err
}

func (sl Repository) GetLogEntryByNotebook(ctx context.Context, name string) ([]*resources.LogEntry, error) {
	sqlcLogEntries, err := sl.queries.GetLogEntryByNotebook(ctx, name)
	if err != nil {
		return nil, err
	}

	return convertSqlcLogEntriesToResource(sqlcLogEntries), err
}

func (sl Repository) GetLogEntryByLogEntryID(ctx context.Context, logentryid int64) (*resources.LogEntry, error) {
	sqlcLogEntry, err := sl.queries.GetLogEntryByLogEntryID(ctx, logentryid)
	if err != nil {
		return nil, err
	}

	return convertSqlcLogEntryToResource(sqlcLogEntry), err
}

func (sl Repository) UpdateLogEntry(ctx context.Context, arg *resources.LogEntry) error {
	err := sl.queries.UpdateLogEntry(ctx, sqlc.UpdateLogEntryParams{
		Tags: sql.NullString{
			String: convertSliceToString(arg.Tags),
			Valid:  true,
		},
		Note:       arg.Entry,
		Notebookid: arg.Notebookid,
		Logentryid: arg.LogEntryID,
	})

	return err
}
