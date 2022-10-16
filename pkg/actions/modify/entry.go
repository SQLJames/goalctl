package modify

import (
	"context"

	"github.com/sqljames/goalctl/pkg/actions/list"
	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func ModifyEntry(confirm bool, targetEntryID int, modOpts EntryModificationOptions) (*resources.LogEntry, error) {

	logEntry, err := list.GetLogEntryByLogEntryID(targetEntryID)
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())
		return nil, err
	}
	logEntry = decodeEntryModificationOptions(logEntry, modOpts)

	if confirm {
		UpdateLogEntry(logEntry)
	}
	return logEntry, nil
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

func UpdateLogEntry(arg *resources.LogEntry) error {
	storagelayer, err := storage.NewVault()
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())
		return err
	}
	return storagelayer.Storage.UpdateLogEntry(context.TODO(), arg)
}
