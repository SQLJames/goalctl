// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: logentry.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createLogEntry = `-- name: CreateLogEntry :one
;

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
  RETURNING logentryid, author, tags, note, createddate, notebookid
`

type CreateLogEntryParams struct {
	Author      sql.NullString `json:"author"`
	Tags        sql.NullString `json:"tags"`
	Note        string         `json:"note"`
	Createddate string         `json:"createddate"`
	Notebookid  int64          `json:"notebookid"`
}

func (q *Queries) CreateLogEntry(ctx context.Context, arg CreateLogEntryParams) (*LogEntry, error) {
	row := q.db.QueryRowContext(ctx, createLogEntry,
		arg.Author,
		arg.Tags,
		arg.Note,
		arg.Createddate,
		arg.Notebookid,
	)
	var i LogEntry
	err := row.Scan(
		&i.Logentryid,
		&i.Author,
		&i.Tags,
		&i.Note,
		&i.Createddate,
		&i.Notebookid,
	)
	return &i, err
}

const getLogEntryByCreatedDate = `-- name: GetLogEntryByCreatedDate :many
SELECT
  logentryid, author, tags, note, createddate, notebookid
FROM
  LogEntry
where
  CreatedDate >= ?
ORDER BY
  LogEntryID
`

func (q *Queries) GetLogEntryByCreatedDate(ctx context.Context, createddate string) ([]*LogEntry, error) {
	rows, err := q.db.QueryContext(ctx, getLogEntryByCreatedDate, createddate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*LogEntry
	for rows.Next() {
		var i LogEntry
		if err := rows.Scan(
			&i.Logentryid,
			&i.Author,
			&i.Tags,
			&i.Note,
			&i.Createddate,
			&i.Notebookid,
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

const getLogEntryByLogEntryID = `-- name: GetLogEntryByLogEntryID :one
SELECT
  logentryid, author, tags, note, createddate, notebookid
FROM
  LogEntry
WHERE
  LogEntryID = ?
`

func (q *Queries) GetLogEntryByLogEntryID(ctx context.Context, logentryid int64) (*LogEntry, error) {
	row := q.db.QueryRowContext(ctx, getLogEntryByLogEntryID, logentryid)
	var i LogEntry
	err := row.Scan(
		&i.Logentryid,
		&i.Author,
		&i.Tags,
		&i.Note,
		&i.Createddate,
		&i.Notebookid,
	)
	return &i, err
}

const getLogEntryByNotebook = `-- name: GetLogEntryByNotebook :many
SELECT
  logentryid, author, tags, note, createddate, notebookid
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
  LogEntryID
`

func (q *Queries) GetLogEntryByNotebook(ctx context.Context, name string) ([]*LogEntry, error) {
	rows, err := q.db.QueryContext(ctx, getLogEntryByNotebook, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*LogEntry
	for rows.Next() {
		var i LogEntry
		if err := rows.Scan(
			&i.Logentryid,
			&i.Author,
			&i.Tags,
			&i.Note,
			&i.Createddate,
			&i.Notebookid,
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
