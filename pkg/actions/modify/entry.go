package modify

import (
	"context"

	"github.com/sqljames/goalctl/pkg/actions/list"
	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func ModifyEntry(confirm bool, targetEntryID int, modOpts EntryModificationOptions) *resources.LogEntry {

	logEntry := list.GetLogEntryByLogEntryID(targetEntryID)

	logEntry = decodeEntryModificationOptions(logEntry, modOpts)

	if confirm {
		UpdateLogEntry(logEntry)
	}
	return logEntry
}

func decodeEntryModificationOptions(logEntry *resources.LogEntry, modOpts EntryModificationOptions) *resources.LogEntry {

	if modOpts.TargetNotebookID != 0 {
		logEntry.Notebookid = modOpts.TargetNotebookID
	}

	if modOpts.EntryDetails != "" {
		logEntry.Entry = modOpts.EntryDetails
	}

	if modOpts.EntryTags != nil {
		logEntry.Tags = modOpts.EntryTags
	}

	return logEntry
}

func UpdateLogEntry(arg *resources.LogEntry) {
	storagelayer, err := storage.NewVault()
	if err != nil {
		jlogr.Logger.ILog.Fatal(err, err.Error())
	}
	storagelayer.Storage.UpdateLogEntry(context.TODO(), arg)
}
