package storage

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/storage/sqlite"
)

type Notebook interface {
	CreateNotebook(ctx context.Context, name string) (resources.Notebook, error)
	GetNotebookIdByName(ctx context.Context, name string) (int64, error)
	GetNotebook(ctx context.Context, name string) (resources.Notebook, error)
	GetNotebooks(ctx context.Context) ([]resources.Notebook, error)
	GetLogEntryByNotebook(ctx context.Context, name string) ([]resources.LogEntry, error)
}
type LogEntry interface {
	CreateLogEntry(ctx context.Context, arg *resources.LogEntry, notebookName string) (*resources.LogEntry, error)
	GetLogEntryByCreatedDate(ctx context.Context, createddate string) ([]resources.LogEntry, error)
	GetLogEntryByLogEntryID(ctx context.Context, logentryid int64) (resources.LogEntry, error)
}
type Goal interface {
	CreateGoal(ctx context.Context, arg *resources.Goal) (*resources.Goal, error)
	GetGoals(ctx context.Context) ([]resources.Goal, error)
}
type Associations interface {
	CreateAssociation(ctx context.Context, arg resources.Association) (resources.Association, error)
	GetAssociations(ctx context.Context) ([]resources.Association, error)
	GetAssociationsByGoalID(ctx context.Context, goalid int) ([]resources.Association, error)
	GetAssociationsByLogEntryID(ctx context.Context, logentryid int) ([]resources.Association, error)
}

type StorageLayer interface {
	Notebook
	LogEntry
	Goal
	Associations
}

func NewStorageLayer() (storageLayer StorageLayer, err error) {
	storageLayer, err = sqlite.NewSQLiteStorage()
	if err != nil {
		return nil, err
	}
	return storageLayer, nil
}
