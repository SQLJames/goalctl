package sqlite

import (
	"context"
	"database/sql"
	"errors"

	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/storage/sqlite/sqlc"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func (sl Repository) CreateGoal(ctx context.Context, arg *resources.Goal) (*resources.Goal, error) {
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
		jlogr.Logger.ILog.Error(err, err.Error())

		return nil, err
	}

	arg.GoalID = int(sqlcGoal.Goalid)

	return arg, err
}

func (sl Repository) GetGoals(ctx context.Context) ([]*resources.Goal, error) {
	sqlcGoals, err := sl.queries.GetGoals(ctx)
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return nil, err
	}

	return convertSqlcGoalsToResource(sqlcGoals), err
}

func (sl Repository) GetGoalByGoalID(ctx context.Context, goalID int) (*resources.Goal, error) {
	sqlcGoal, err := sl.queries.GetGoalByGoalID(ctx, int64(goalID))

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		jlogr.Logger.ILog.Error(err, err.Error())
		
		return nil, ErrNoRows
	}

	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return nil, err
	}

	return convertSqlcGoalToResource(sqlcGoal), err
}

func (sl Repository) UpdateGoal(ctx context.Context, arg *resources.Goal) error {
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

	return err
}
