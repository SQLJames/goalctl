
-- name: GetLogEntryByCreatedDate :many
SELECT
  *
FROM
  LogEntry
where
  CreatedDate >= ?
ORDER BY
  LogEntryID;

-- name: GetLogEntryByNotebook :many
SELECT
  *
FROM
  LogEntry
WHERE
  NotebookID = (
    SELECT
      NotebookID
    FROM
      Notebook
    WHERE
      name = ?
  )
ORDER BY
  LogEntryID;

-- name: GetLogEntryByLogEntryID :one
SELECT
  *
FROM
  LogEntry
WHERE
  LogEntryID = ?
;

-- name: CreateLogEntry :one
INSERT INTO LogEntry (
    author, tags, note, createddate, notebookid
  )
VALUES(
  ?,
  ?, 
  ?, 
  ?, 
  ?
  )
  RETURNING *;

-- name: GetLogEntries :many
SELECT
  *
FROM
  LogEntry
ORDER BY
  LogEntryID;