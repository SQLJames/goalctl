package list

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
)

func GetNotebooks() (notebookList []*resources.Notebook) {
	storagelayer := storage.NewVault()
	notebooks := storagelayer.Storage.GetNotebooks(context.TODO())

	return notebooks
}

func GetEntriesForNotebook(notebookName string) (entries []*resources.LogEntry) {
	storagelayer := storage.NewVault()

	return storagelayer.Storage.GetLogEntryByNotebook(context.TODO(), notebookName)
}
