package sqlite

import (
	"context"
	"database/sql"

	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/storage/sqlite/sqlc"
)

func (SL SQLiteStorage) CreateGoal(ctx context.Context, arg resources.Goal) (resources.Goal, error) {
	sqlcGoal, err := SL.queries.CreateGoal(ctx, sqlc.CreateGoalParams{
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
		return resources.Goal{}, err
	}
	arg.GoalID = int(sqlcGoal.Goalid)
	return arg, err
}

func (SL SQLiteStorage) GetGoals(ctx context.Context) ([]resources.Goal, error) {
	sqlcGoals, err := SL.queries.GetGoals(ctx)
	if err != nil {
		return nil, err
	}
	return convertSqlcGoalsToResource(sqlcGoals), err
}
