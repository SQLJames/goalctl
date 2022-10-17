package create

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func LogEntry(newLogEntry NewLogEntry) (*resources.LogEntry, error) {
	storagelayer, err := storage.NewVault()
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return nil, err
	}

	logEntry := resources.NewLogEntry(newLogEntry.LogEntry, newLogEntry.Tags)

	logEntry, err = storagelayer.Storage.CreateLogEntry(context.TODO(), logEntry, newLogEntry.NotebookName)
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return nil, err
	}

	return logEntry, nil
}
