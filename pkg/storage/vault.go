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
	CreateNotebook(ctx context.Context, name string) resources.Notebook
	GetNotebookIDByName(ctx context.Context, name string) int64
	GetNotebook(ctx context.Context, name string) resources.Notebook
	GetNotebooks(ctx context.Context) []*resources.Notebook
	GetLogEntryByNotebook(ctx context.Context, name string) []*resources.LogEntry
}
type LogEntry interface {
	CreateLogEntry(ctx context.Context, arg *resources.LogEntry, notebookName string) *resources.LogEntry
	GetLogEntries(ctx context.Context) []*resources.LogEntry
	GetLogEntryByCreatedDate(ctx context.Context, createddate string) []*resources.LogEntry
	GetLogEntryByLogEntryID(ctx context.Context, logentryid int64) *resources.LogEntry
}
type Goal interface {
	CreateGoal(ctx context.Context, arg *resources.Goal) *resources.Goal
	GetGoals(ctx context.Context) []*resources.Goal
}
type Associations interface {
	CreateAssociation(ctx context.Context, arg resources.Association) resources.Association
	GetAssociations(ctx context.Context) []*resources.Association
	GetAssociationsByGoalID(ctx context.Context, goalid int) []*resources.Association
	GetAssociationsByLogEntryID(ctx context.Context, logentryid int) []*resources.Association
}

type Repository interface {
	Notebook
	LogEntry
	Goal
	Associations
}

func NewVault() (vault Vault) {
	return Vault{Storage: sqlite.NewSQLiteStorage()}
}
