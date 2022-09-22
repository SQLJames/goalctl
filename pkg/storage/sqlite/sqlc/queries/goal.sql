
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