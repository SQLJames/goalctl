package create

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func CreateLogEntry(newLogEntry NewLogEntry) *resources.LogEntry {
	storagelayer, err := storage.NewVault()
	if err != nil {
		jlogr.Logger.ILog.Fatal(err, err.Error())
	}
	le := resources.NewLogEntry(newLogEntry.LogEntry, newLogEntry.Tags)

	return storagelayer.Storage.CreateLogEntry(context.TODO(), le, newLogEntry.NotebookName)
}
