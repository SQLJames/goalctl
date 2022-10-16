package create

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func CreateNotebook(notebookName string) (data *resources.Notebook, err error) {
	if notebookName == "" {
		jlogr.Logger.ILog.Error(&EmptryStringError{}, "Error creating notebook", "function", "actionCreateNotebook")

		return nil, &EmptryStringError{}
	}

	storagelayer := storage.NewVault()
	row := storagelayer.Storage.CreateNotebook(context.TODO(), notebookName)

	return &row, err
}
