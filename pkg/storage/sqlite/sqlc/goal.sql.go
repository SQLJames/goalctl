// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: goal.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createGoal = `-- name: CreateGoal :one
INSERT INTO Goal(
    author,
    duedate,
    createddate,
    goal, 
    details,
    priority,
    status
  )
VALUES(
  ?,
  ?, 
  ?, 
  ?, 
  ?, 
  ?,
  ?
  )
  RETURNING goalid, duedate, author, createddate, goal, details, priority, status
`

type CreateGoalParams struct {
	Author      sql.NullString `json:"author"`
	Duedate     sql.NullString `json:"duedate"`
	Createddate string         `json:"createddate"`
	Goal        string         `json:"goal"`
	Details     string         `json:"details"`
	Priority    int64          `json:"priority"`
	Status      string         `json:"status"`
}

func (q *Queries) CreateGoal(ctx context.Context, arg CreateGoalParams) (*Goal, error) {
	row := q.db.QueryRowContext(ctx, createGoal,
		arg.Author,
		arg.Duedate,
		arg.Createddate,
		arg.Goal,
		arg.Details,
		arg.Priority,
		arg.Status,
	)
	var i Goal
	err := row.Scan(
		&i.Goalid,
		&i.Duedate,
		&i.Author,
		&i.Createddate,
		&i.Goal,
		&i.Details,
		&i.Priority,
		&i.Status,
	)
	return &i, err
}

const getGoals = `-- name: GetGoals :many
SELECT
  goalid, duedate, author, createddate, goal, details, priority, status
FROM
  Goal
`

func (q *Queries) GetGoals(ctx context.Context) ([]*Goal, error) {
	rows, err := q.db.QueryContext(ctx, getGoals)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Goal
	for rows.Next() {
		var i Goal
		if err := rows.Scan(
			&i.Goalid,
			&i.Duedate,
			&i.Author,
			&i.Createddate,
			&i.Goal,
			&i.Details,
			&i.Priority,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateGoal = `-- name: UpdateGoal :exec
UPDATE Goal
SET duedate = ?,
    goal = ?, 
    details = ?, 
    priority = ?,
    status = ?
WHERE 
  GoalID = ?
`

type UpdateGoalParams struct {
	Duedate  sql.NullString `json:"duedate"`
	Goal     string         `json:"goal"`
	Details  string         `json:"details"`
	Priority int64          `json:"priority"`
	Status   string         `json:"status"`
	Goalid   int64          `json:"goalid"`
}

func (q *Queries) UpdateGoal(ctx context.Context, arg UpdateGoalParams) error {
	_, err := q.db.ExecContext(ctx, updateGoal,
		arg.Duedate,
		arg.Goal,
		arg.Details,
		arg.Priority,
		arg.Status,
		arg.Goalid,
	)
	return err
}
