package sqlite

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/sqljames/goalctl/pkg/log"

	"github.com/sqljames/goalctl/pkg/storage/resources"
)

func (sl Repository) CreateNotebook(ctx context.Context, name string) (resources.Notebook, error) {
	log.Logger.Trace(fmt.Sprintf("notebook name provided %s", name))
	if name == "" {
		err := errors.New("notebook name can not be empty")
		log.Logger.Error(err, "I have no idea how you did this either")
	}
	sqlcNotebook, err := sl.queries.CreateNotebook(ctx, name)
	if err != nil && strings.Contains(err.Error(), "2067") {

		return resources.Notebook{}, fmt.Errorf("a notebook resource with that name already exists, err: %s", err.Error())
	}
	if err != nil {
		return resources.Notebook{}, err
	}
	return resources.Notebook{Notebookid: sqlcNotebook.Notebookid, Name: sqlcNotebook.Name}, err
}

func (sl Repository) GetNotebookIDByName(ctx context.Context, name string) (int64, error) {
	return sl.queries.GetNotebookIDByName(ctx, name)
}

func (sl Repository) GetNotebook(ctx context.Context, name string) (notebook resources.Notebook, err error) {
	sqlcNotebook, err := sl.queries.GetNotebook(ctx, name)
	if err != nil {
		return resources.Notebook{}, err
	}
	return resources.Notebook{Notebookid: sqlcNotebook.Notebookid, Name: sqlcNotebook.Name}, err
}

func (sl Repository) GetNotebooks(ctx context.Context) (notebooks []resources.Notebook, err error) {
	sqlcNotebooks, err := sl.queries.GetNotebooks(ctx)
	if err != nil {
		return nil, err
	}
	for _, sqlcNotebook := range sqlcNotebooks {
		notebooks = append(notebooks, resources.Notebook{Notebookid: sqlcNotebook.Notebookid, Name: sqlcNotebook.Name})
	}
	return notebooks, err
}
