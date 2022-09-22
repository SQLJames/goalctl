
-- name: GetAssociationsByGoalID :many
SELECT
  *
FROM
  GoalToLogEntry
WHERE
  GoalID = ?
ORDER BY
  GoalID;


-- name: GetAssociationsByLogEntryID :many
SELECT
  *
FROM
  GoalToLogEntry
WHERE
  LogEntryID = ?
ORDER BY
  GoalID;

-- name: GetAssociations :many
SELECT
  *
FROM
  GoalToLogEntry
ORDER BY
  GoalID;


-- name: CreateAssociation :one
INSERT INTO GoalToLogEntry (
    goalid ,
    logentryid
  )
VALUES(
  ?,
  ?
  )
  RETURNING *;