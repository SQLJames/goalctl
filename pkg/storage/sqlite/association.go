package sqlite

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/storage/sqlite/sqlc"
)

func (sl Repository) CreateAssociation(ctx context.Context, arg resources.Association) (resources.Association, error) {
	_, err := sl.queries.CreateAssociation(ctx, sqlc.CreateAssociationParams{
		Goalid:     int64(arg.GoalID),
		Logentryid: int64(arg.LogEntryID),
	})

	return arg, err
}

func (sl Repository) GetAssociations(ctx context.Context) ([]*resources.Association, error) {
	associations, err := sl.queries.GetAssociations(ctx)
	
	return convertSqlcGoalToLogEntriesToResource(associations), err
}

func (sl Repository) GetAssociationsByGoalID(ctx context.Context, goalid int) ([]*resources.Association, error) {
	associations, err := sl.queries.GetAssociationsByGoalID(ctx, int64(goalid))

	return convertSqlcGoalToLogEntriesToResource(associations), err
}

func (sl Repository) GetAssociationsByLogEntryID(ctx context.Context, logentryid int) ([]*resources.Association, error) {
	associations, err := sl.queries.GetAssociationsByLogEntryID(ctx, int64(logentryid))

	return convertSqlcGoalToLogEntriesToResource(associations), err
}
