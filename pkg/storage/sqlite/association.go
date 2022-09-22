package sqlite

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/storage/sqlite/sqlc"
)

func (SL SQLiteStorage) CreateAssociation(ctx context.Context, arg resources.Association) (resources.Association, error) {
	_, err := SL.queries.CreateAssociation(ctx, sqlc.CreateAssociationParams{
		Goalid:     int64(arg.GoalID),
		Logentryid: int64(arg.LogEntryID),
	})
	return arg, err
}

func (SL SQLiteStorage) GetAssociations(ctx context.Context) ([]resources.Association, error) {
	associations, err := SL.queries.GetAssociations(ctx)
	return convertSqlcGoalToLogEntriesToResource(associations), err
}

func (SL SQLiteStorage) GetAssociationsByGoalID(ctx context.Context, goalid int) ([]resources.Association, error) {
	associations, err := SL.queries.GetAssociationsByGoalID(ctx, int64(goalid))
	return convertSqlcGoalToLogEntriesToResource(associations), err
}

func (SL SQLiteStorage) GetAssociationsByLogEntryID(ctx context.Context, logentryid int) ([]resources.Association, error) {
	associations, err := SL.queries.GetAssociationsByLogEntryID(ctx, int64(logentryid))
	return convertSqlcGoalToLogEntriesToResource(associations), err
}
