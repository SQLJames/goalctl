package sqlite

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/sqljames/goalctl/pkg/log"

	"github.com/sqljames/goalctl/pkg/storage/resources"
)

func (SL SQLiteStorage) CreateNotebook(ctx context.Context, name string) (resources.Notebook, error) {
	log.Logger.Trace(fmt.Sprintf("notebook name provided %s", name))
	if name == "" {
		err := errors.New("notebook name can not be empty")
		log.Logger.Error(err, "I have no idea how you did this either")
	}
	sqlcNotebook, err := SL.queries.CreateNotebook(ctx, name)
	if err != nil && strings.Contains(err.Error(), "2067") {

		return resources.Notebook{}, fmt.Errorf("a notebook resource with that name already exists, err: %s", err.Error())
	}
	if err != nil {
		return resources.Notebook{}, err
	}
	return resources.Notebook{Notebookid: sqlcNotebook.Notebookid, Name: sqlcNotebook.Name}, err
}

func (SL SQLiteStorage) GetNotebookIdByName(ctx context.Context, name string) (int64, error) {
	return SL.queries.GetNotebookIdByName(ctx, name)
}

func (SL SQLiteStorage) GetNotebook(ctx context.Context, name string) (Notebook resources.Notebook, err error) {
	sqlcNotebook, err := SL.queries.GetNotebook(ctx, name)
	if err != nil {
		return resources.Notebook{}, err
	}
	return resources.Notebook{Notebookid: sqlcNotebook.Notebookid, Name: sqlcNotebook.Name}, err
}

func (SL SQLiteStorage) GetNotebooks(ctx context.Context) (Notebooks []resources.Notebook, err error) {
	sqlcNotebooks, err := SL.queries.GetNotebooks(ctx)
	if err != nil {
		return nil, err
	}
	for _, sqlcNotebook := range sqlcNotebooks {
		Notebooks = append(Notebooks, resources.Notebook{Notebookid: sqlcNotebook.Notebookid, Name: sqlcNotebook.Name})
	}
	return Notebooks, err
}
