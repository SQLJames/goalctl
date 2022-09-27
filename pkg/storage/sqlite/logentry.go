package sqlite

import (
	"context"
	"database/sql"

	"github.com/sqljames/goalctl/pkg/log"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/storage/sqlite/sqlc"
)

func (sl Repository) GetLogEntries(ctx context.Context) (logEntries []*resources.LogEntry) {
	sqlcLogEntries, err := sl.queries.GetLogEntries(ctx)
	if err != nil {
		log.Logger.ILog.Fatal(err, "error running query")
	}

	return convertSqlcLogEntriesToResource(sqlcLogEntries)
}

func (sl Repository) CreateLogEntry(ctx context.Context, arg *resources.LogEntry, notebookName string) *resources.LogEntry {
	NotebookID := sl.GetNotebookIDByName(ctx, notebookName)
	if NotebookID == 0 {
		log.Logger.ILog.Warn("No results for notebook, Attempting to create a new notebook")

		notebook := sl.CreateNotebook(ctx, notebookName)
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
		log.Logger.ILog.Fatal(err, "error running query")
	}

	arg.LogEntryID = Entry.Logentryid
	arg.Notebookid = Entry.Notebookid

	return arg
}

func (sl Repository) GetLogEntryByCreatedDate(ctx context.Context, createddate string) (logEntries []*resources.LogEntry) {
	sqlcLogEntries, err := sl.queries.GetLogEntryByCreatedDate(ctx, createddate)
	if err != nil {
		log.Logger.ILog.Fatal(err, "error running query")
	}

	return convertSqlcLogEntriesToResource(sqlcLogEntries)
}

func (sl Repository) GetLogEntryByNotebook(ctx context.Context, name string) []*resources.LogEntry {
	sqlcLogEntries, err := sl.queries.GetLogEntryByNotebook(ctx, name)
	if err != nil {
		log.Logger.ILog.Fatal(err, "error running query")
	}

	return convertSqlcLogEntriesToResource(sqlcLogEntries)
}

func (sl Repository) GetLogEntryByLogEntryID(ctx context.Context, logentryid int64) *resources.LogEntry {
	sqlcLogEntry, err := sl.queries.GetLogEntryByLogEntryID(ctx, logentryid)
	if err != nil {
		log.Logger.ILog.Fatal(err, "error running query")
	}

	return convertSqlcLogEntryToResource(sqlcLogEntry)
}
