
-- name: GetGoals :many
SELECT
  *
FROM
  Goal;



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


-- name: UpdateGoal :exec
UPDATE Goal
SET duedate = ?,
    goal = ?, 
    details = ?, 
    priority = ?,
    status = ?
WHERE 
  GoalID = ?;