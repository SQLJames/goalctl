package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/storage/sqlite/sqlc"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
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

func (sl Repository) GetGoals(ctx context.Context) []*resources.Goal {
	sqlcGoals, err := sl.queries.GetGoals(ctx)
	if err != nil {
		jlogr.Logger.ILog.Fatal(err, "error running query")
	}

	return convertSqlcGoalsToResource(sqlcGoals)
}

func (sl Repository) GetGoalByGoalID(ctx context.Context, goalID int) *resources.Goal {
	sqlcGoal, err := sl.queries.GetGoalByGoalID(ctx, int64(goalID))

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		jlogr.Logger.ILog.Fatal(err, fmt.Sprintf("The Goal with the ID of %d, does not exist.", goalID))
	}

	if err != nil {
		jlogr.Logger.ILog.Fatal(err, "error running query")
	}

	return convertSqlcGoalToResource(sqlcGoal)
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
