package sqlite

import (
	"context"
	"fmt"

	"github.com/sqljames/goalctl/pkg/log"
	"github.com/sqljames/goalctl/pkg/storage/resources"
)

func (sl Repository) CreateNotebook(ctx context.Context, name string) resources.Notebook {
	log.Logger.Trace(fmt.Sprintf("notebook name provided %s", name))

	sqlcNotebook, err := sl.queries.CreateNotebook(ctx, name)
	if err != nil {
		// we can check for unique value error by comparing it with sqlite3.SQLITE_CONSTRAINT_UNIQUE
		log.Logger.Fatal(err, "error running query")
	}

	return resources.Notebook{Notebookid: sqlcNotebook.Notebookid, Name: sqlcNotebook.Name}
}

func (sl Repository) GetNotebookIDByName(ctx context.Context, name string) int64 {
	id, err := sl.queries.GetNotebookIDByName(ctx, name)
	if err != nil {
		log.Logger.Fatal(err, "error running query")
	}
	return id
}

func (sl Repository) GetNotebook(ctx context.Context, name string) (notebook resources.Notebook) {
	sqlcNotebook, err := sl.queries.GetNotebook(ctx, name)
	if err != nil {
		log.Logger.Fatal(err, "error running query")
	}
	return resources.Notebook{Notebookid: sqlcNotebook.Notebookid, Name: sqlcNotebook.Name}
}

func (sl Repository) GetNotebooks(ctx context.Context) []*resources.Notebook {
	sqlcEntries, err := sl.queries.GetNotebooks(ctx)
	if err != nil {
		log.Logger.Fatal(err, "error running query")
	}

	var notebooks = make([]*resources.Notebook, len(sqlcEntries))

	for index, sqlcEntry := range sqlcEntries {
		notebooks[index] = &resources.Notebook{Notebookid: sqlcEntry.Notebookid, Name: sqlcEntry.Name}
	}
	return notebooks
}
