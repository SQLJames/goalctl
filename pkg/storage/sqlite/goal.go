package sqlite

import (
	"context"
	"database/sql"

	"github.com/sqljames/goalctl/pkg/log"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/storage/sqlite/sqlc"
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
		log.Logger.Fatal(err, "error running query")
	}

	arg.GoalID = int(sqlcGoal.Goalid)

	return arg
}

func (sl Repository) GetGoals(ctx context.Context) []*resources.Goal {
	sqlcGoals, err := sl.queries.GetGoals(ctx)
	if err != nil {
		log.Logger.Fatal(err, "error running query")
	}
	
	return convertSqlcGoalsToResource(sqlcGoals)
}
