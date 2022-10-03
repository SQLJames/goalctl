package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func (sl Repository) CreateNotebook(ctx context.Context, name string) resources.Notebook {
	jlogr.Logger.ILog.Trace(fmt.Sprintf("notebook name provided %s", name))

	sqlcNotebook, err := sl.queries.CreateNotebook(ctx, name)
	if err != nil {
		// we can check for unique value error by comparing it with sqlite3.SQLITE_CONSTRAINT_UNIQUE
		jlogr.Logger.ILog.Fatal(err, "error running query")
	}

	return resources.Notebook{Notebookid: sqlcNotebook.Notebookid, Name: sqlcNotebook.Name}
}

func (sl Repository) GetNotebookIDByName(ctx context.Context, name string) int64 {
	notebookID, err := sl.queries.GetNotebookIDByName(ctx, name)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		jlogr.Logger.ILog.Fatal(err, "error running query")
	}

	if errors.Is(err, sql.ErrNoRows) {
		notebookID = 0
	}

	return notebookID
}

func (sl Repository) GetNotebook(ctx context.Context, name string) (notebook resources.Notebook) {
	sqlcNotebook, err := sl.queries.GetNotebook(ctx, name)
	if err != nil {
		jlogr.Logger.ILog.Fatal(err, "error running query")
	}

	return resources.Notebook{Notebookid: sqlcNotebook.Notebookid, Name: sqlcNotebook.Name}
}

func (sl Repository) GetNotebooks(ctx context.Context) []*resources.Notebook {
	sqlcEntries, err := sl.queries.GetNotebooks(ctx)
	if err != nil {
		jlogr.Logger.ILog.Fatal(err, "error running query")
	}

	var notebooks = make([]*resources.Notebook, len(sqlcEntries))

	for index, sqlcEntry := range sqlcEntries {
		notebooks[index] = &resources.Notebook{Notebookid: sqlcEntry.Notebookid, Name: sqlcEntry.Name}
	}

	return notebooks
}
