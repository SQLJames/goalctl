package list

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
)

func ListEntries(notebookName string) resources.Notebook {
	NotebookEntries := GetEntriesForNotebook(notebookName)

	notebook := resources.Notebook{
		Name:    notebookName,
		Entries: NotebookEntries,
	}

	return notebook
}

func lookupLogEntry(entries []*resources.LogEntry, logEntryID int) *resources.LogEntry {
	for _, entry := range entries {
		if entry.LogEntryID == int64(logEntryID) {
			return entry
		}
	}

	return nil
}

func GetLogEntryByLogEntryID(logentryid int) *resources.LogEntry {
	storagelayer := storage.NewVault()

	return storagelayer.Storage.GetLogEntryByLogEntryID(context.TODO(), int64(logentryid))
}
