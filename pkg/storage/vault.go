package storage

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/storage/sqlite"
)

type Notebook interface {
	CreateNotebook(ctx context.Context, name string) (resources.Notebook, error)
	GetNotebookIDByName(ctx context.Context, name string) (int64, error)
	GetNotebook(ctx context.Context, name string) (resources.Notebook, error)
	GetNotebooks(ctx context.Context) ([]*resources.Notebook, error)
	GetLogEntryByNotebook(ctx context.Context, name string) ([]*resources.LogEntry, error)
}
type LogEntry interface {
	CreateLogEntry(ctx context.Context, arg *resources.LogEntry, notebookName string) (*resources.LogEntry, error)
	GetLogEntryByCreatedDate(ctx context.Context, createddate string) ([]*resources.LogEntry, error)
	GetLogEntryByLogEntryID(ctx context.Context, logentryid int64) (*resources.LogEntry, error)
}
type Goal interface {
	CreateGoal(ctx context.Context, arg *resources.Goal) (*resources.Goal, error)
	GetGoals(ctx context.Context) ([]*resources.Goal, error)
}
type Associations interface {
	CreateAssociation(ctx context.Context, arg resources.Association) (resources.Association, error)
	GetAssociations(ctx context.Context) ([]*resources.Association, error)
	GetAssociationsByGoalID(ctx context.Context, goalid int) ([]*resources.Association, error)
	GetAssociationsByLogEntryID(ctx context.Context, logentryid int) ([]*resources.Association, error)
}

type Vault interface {
	Notebook
	LogEntry
	Goal
	Associations
}

func NewVault() (vault Vault) {

	return sqlite.NewSQLiteStorage()
}
