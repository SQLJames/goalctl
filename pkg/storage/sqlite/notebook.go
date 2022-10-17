package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func (sl Repository) CreateNotebook(ctx context.Context, name string) (resources.Notebook, error) {
	jlogr.Logger.ILog.Trace(fmt.Sprintf("notebook name provided %s", name))

	sqlcNotebook, err := sl.queries.CreateNotebook(ctx, name)
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		// we can check for unique value error by comparing it with sqlite3.SQLITE_CONSTRAINT_UNIQUE
		return resources.Notebook{}, err
	}

	return resources.Notebook{Notebookid: sqlcNotebook.Notebookid, Name: sqlcNotebook.Name}, err
}

func (sl Repository) GetNotebookIDByName(ctx context.Context, name string) (int64, error) {
	notebookID, err := sl.queries.GetNotebookIDByName(ctx, name)

	if errors.Is(err, sql.ErrNoRows) {
		return 0, ErrNoRows
	}

	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return -1, err
	}

	return notebookID, nil
}

func (sl Repository) GetNotebook(ctx context.Context, name string) (notebook resources.Notebook, err error) {
	sqlcNotebook, err := sl.queries.GetNotebook(ctx, name)
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return resources.Notebook{}, err
	}

	return resources.Notebook{Notebookid: sqlcNotebook.Notebookid, Name: sqlcNotebook.Name}, nil
}

func (sl Repository) GetNotebooks(ctx context.Context) ([]*resources.Notebook, error) {
	sqlcEntries, err := sl.queries.GetNotebooks(ctx)
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())
		
		return nil, err
	}

	var notebooks = make([]*resources.Notebook, len(sqlcEntries))

	for index, sqlcEntry := range sqlcEntries {
		notebooks[index] = &resources.Notebook{Notebookid: sqlcEntry.Notebookid, Name: sqlcEntry.Name}
	}

	return notebooks, nil
}
