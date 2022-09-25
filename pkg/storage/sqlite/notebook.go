package sqlite

import (
	"context"
	"fmt"
	"strings"

	"github.com/sqljames/goalctl/pkg/log"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	sqlite3 "modernc.org/sqlite/lib"
)

func (sl Repository) CreateNotebook(ctx context.Context, name string) (resources.Notebook, error) {
	log.Logger.Trace(fmt.Sprintf("notebook name provided %s", name))
	sqlcNotebook, err := sl.queries.CreateNotebook(ctx, name)
	if err != nil && strings.Contains(err.Error(), fmt.Sprintf("%d", sqlite3.SQLITE_CONSTRAINT_UNIQUE)) {
		return resources.Notebook{}, fmt.Errorf("CreateNotebook, err: %w", err)
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

func (sl Repository) GetNotebooks(ctx context.Context) ([]resources.Notebook, error) {
	sqlcEntries, err := sl.queries.GetNotebooks(ctx)
	if err != nil {
		return nil, err
	}
	var notebooks = make([]resources.Notebook, len(sqlcEntries))
	for index, sqlcEntry := range sqlcEntries {
		notebooks[index] =  resources.Notebook{Notebookid: sqlcEntry.Notebookid, Name: sqlcEntry.Name}
	}
	return notebooks, err
}
