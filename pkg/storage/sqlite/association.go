package sqlite

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/storage/sqlite/sqlc"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func (sl Repository) CreateAssociation(ctx context.Context, arg resources.Association) resources.Association {
	_, err := sl.queries.CreateAssociation(ctx, sqlc.CreateAssociationParams{
		Goalid:     int64(arg.GoalID),
		Logentryid: int64(arg.LogEntryID),
	})
	if err != nil {
		jlogr.Logger.ILog.Fatal(err, "error running query")
	}

	return arg
}

func (sl Repository) GetAssociations(ctx context.Context) []*resources.Association {
	associations, err := sl.queries.GetAssociations(ctx)
	if err != nil {
		jlogr.Logger.ILog.Fatal(err, "error running query")
	}

	return convertSqlcGoalToLogEntriesToResource(associations)
}

func (sl Repository) GetAssociationsByGoalID(ctx context.Context, goalid int) []*resources.Association {
	associations, err := sl.queries.GetAssociationsByGoalID(ctx, int64(goalid))
	if err != nil {
		jlogr.Logger.ILog.Fatal(err, "error running query")
	}

	return convertSqlcGoalToLogEntriesToResource(associations)
}

func (sl Repository) GetAssociationsByLogEntryID(ctx context.Context, logentryid int) []*resources.Association {
	associations, err := sl.queries.GetAssociationsByLogEntryID(ctx, int64(logentryid))
	if err != nil {
		jlogr.Logger.ILog.Fatal(err, "error running query")
	}

	return convertSqlcGoalToLogEntriesToResource(associations)
}
