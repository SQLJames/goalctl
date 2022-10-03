package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/storage/sqlite/sqlc"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
	sqlcpkg "github.com/yiplee/sqlc"
)

func (sl Repository) CreateGoal(ctx context.Context, arg *resources.Goal) *resources.Goal {
	sqlcGoal, err := sl.queries.CreateGoal(ctx, sqlc.CreateGoalParams{
		Author: sql.NullString{
			String: arg.Author,
			Valid:  true,
		},
		Duedate: sql.NullString{
			String: arg.Deadline,
			Valid:  true,
		},
		Createddate: arg.CreatedDate,
		Goal:        arg.Goal,
		Details:     arg.Details,
		Priority:    int64(arg.Priority),
		Status:      arg.Status,
	})
	if err != nil {
		jlogr.Logger.ILog.Fatal(err, "error running query")
	}

	arg.GoalID = int(sqlcGoal.Goalid)

	return arg
}
func (sl Repository) GetGoals(ctx context.Context, goalFilter *resources.Goal) []*resources.Goal {
	query := sqlc.New(sqlcpkg.Wrap(&sl.db))

	sqlcEntries, err := query.GetGoals(sqlcpkg.Build(ctx, func(builder *sqlcpkg.Builder) {
		if goalFilter != nil && goalFilter.GoalID >= 0 {
			builder.Where("GoalID = $1", goalFilter.GoalID)
		}
		builder.Order("GoalID DESC")
	}))
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		jlogr.Logger.ILog.Fatal(err, fmt.Sprintf("The Goal with the ID of %d, does not exist.", goalFilter.GoalID))
	}

	if err != nil {
		jlogr.Logger.ILog.Fatal(err, "error running query")
	}

	return convertSqlcGoalsToResource(sqlcEntries)
}

func (sl Repository) UpdateGoal(ctx context.Context, arg *resources.Goal) {
	err := sl.queries.UpdateGoal(ctx, sqlc.UpdateGoalParams{
		Goalid: int64(arg.GoalID),
		Duedate: sql.NullString{
			String: arg.Deadline,
			Valid:  true,
		},
		Goal:     arg.Goal,
		Details:  arg.Details,
		Priority: int64(arg.Priority),
		Status:   arg.Status,
	})

	if err != nil {
		jlogr.Logger.ILog.Fatal(err, "error running query")
	}
}
