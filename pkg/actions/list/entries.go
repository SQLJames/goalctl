package list

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func Entries(notebookName string) (*resources.Notebook, error) {
	NotebookEntries, err := EntriesForNotebook(notebookName)
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return nil, err
	}

	notebook := resources.Notebook{
		Name:    notebookName,
		Entries: NotebookEntries,
	}

	return &notebook, nil
}

func lookupLogEntry(entries []*resources.LogEntry, logEntryID int) *resources.LogEntry {
	for _, entry := range entries {
		if entry.LogEntryID == int64(logEntryID) {
			return entry
		}
	}

	return nil
}

func LogEntryByLogEntryID(logentryid int) (*resources.LogEntry, error) {
	storagelayer, err := storage.NewVault()
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return nil, err
	}

	return storagelayer.Storage.GetLogEntryByLogEntryID(context.TODO(), int64(logentryid))
}

func EntriesForNotebook(notebookName string) (entries []*resources.LogEntry, err error) {
	storagelayer, err := storage.NewVault()
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return nil, err
	}

	return storagelayer.Storage.GetLogEntryByNotebook(context.TODO(), notebookName)
}
