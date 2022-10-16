package create

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
)

func CreateLogEntry(newLogEntry NewLogEntry) *resources.LogEntry {
	storagelayer := storage.NewVault()
	le := resources.NewLogEntry(newLogEntry.LogEntry, newLogEntry.Tags)

	return storagelayer.Storage.CreateLogEntry(context.TODO(), le, newLogEntry.NotebookName)
}
