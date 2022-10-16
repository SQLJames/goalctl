package storage

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/storage/sqlite"
)

type Vault struct {
	Storage Repository
}

type Notebook interface {
	CreateNotebook(ctx context.Context, name string) (resources.Notebook, error)
	GetNotebookIDByName(ctx context.Context, name string) (int64, error)
	GetNotebook(ctx context.Context, name string) (resources.Notebook, error)
	GetNotebooks(ctx context.Context) ([]*resources.Notebook, error)
}
type LogEntry interface {
	GetLogEntryByNotebook(ctx context.Context, name string) ([]*resources.LogEntry, error)
	CreateLogEntry(ctx context.Context, arg *resources.LogEntry, notebookName string) (*resources.LogEntry, error)
	GetLogEntries(ctx context.Context) ([]*resources.LogEntry, error)
	GetLogEntryByCreatedDate(ctx context.Context, createddate string) ([]*resources.LogEntry, error)
	GetLogEntryByLogEntryID(ctx context.Context, logentryid int64) (*resources.LogEntry, error)
	UpdateLogEntry(ctx context.Context, arg *resources.LogEntry) error
}
type Goal interface {
	CreateGoal(ctx context.Context, arg *resources.Goal) (*resources.Goal, error)
	GetGoals(ctx context.Context) ([]*resources.Goal, error)
	UpdateGoal(ctx context.Context, arg *resources.Goal) error
	GetGoalByGoalID(ctx context.Context, goalID int) (*resources.Goal, error)
}
type Associations interface {
	CreateAssociation(ctx context.Context, arg resources.Association) (resources.Association, error)
	GetAssociations(ctx context.Context) ([]*resources.Association, error)
	GetAssociationsByGoalID(ctx context.Context, goalid int) ([]*resources.Association, error)
	GetAssociationsByLogEntryID(ctx context.Context, logentryid int) ([]*resources.Association, error)
}

type Repository interface {
	Notebook
	LogEntry
	Goal
	Associations
}

func NewVault() (vault Vault, err error) {
	repo, err := sqlite.NewSQLiteStorage()
	if err != nil {
		return Vault{}, err
	}

	return Vault{Storage: repo}, nil
}
