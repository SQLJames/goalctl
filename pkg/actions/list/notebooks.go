package list

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func Notebooks() (notebookList []*resources.Notebook, err error) {
	storagelayer, err := storage.NewVault()
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return nil, err
	}

	notebooks, err := storagelayer.Storage.GetNotebooks(context.TODO())
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return nil, err
	}

	return notebooks, err
}
