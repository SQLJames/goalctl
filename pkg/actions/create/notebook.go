package create

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func Notebook(notebookName string) (data *resources.Notebook, err error) {
	if notebookName == "" {
		jlogr.Logger.ILog.Error(&EmptryStringError{}, "Error creating notebook", "function", "actionCreateNotebook")

		return nil, &EmptryStringError{}
	}

	storagelayer, err := storage.NewVault()
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return nil, err
	}
	
	results, err := storagelayer.Storage.CreateNotebook(context.TODO(), notebookName)
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return nil, err
	}
	
	return &results, err
}
