package list

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func GetNotebooks() (notebookList []*resources.Notebook) {
	storagelayer, err := storage.NewVault()
	if err != nil {
		jlogr.Logger.ILog.Fatal(err, err.Error())
	}
	notebooks := storagelayer.Storage.GetNotebooks(context.TODO())

	return notebooks
}

func GetEntriesForNotebook(notebookName string) (entries []*resources.LogEntry) {
	storagelayer, err := storage.NewVault()
	if err != nil {
		jlogr.Logger.ILog.Fatal(err, err.Error())
	}

	return storagelayer.Storage.GetLogEntryByNotebook(context.TODO(), notebookName)
}
