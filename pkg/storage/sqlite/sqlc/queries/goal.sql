
-- name: GetGoals :many
SELECT
  *
FROM
  Goal
ORDER BY
  GoalID;


-- name: CreateGoal :one
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
  RETURNING *;

-- name: GetGoalByGoalID :one
SELECT
  *
FROM
  Goal
WHERE
  GoalID = ?;


-- name: UpdateGoal :exec
UPDATE Goal
SET duedate = ?,
    goal = ?, 
    details = ?, 
    priority = ?,
    status = ?
WHERE 
  GoalID = ?;