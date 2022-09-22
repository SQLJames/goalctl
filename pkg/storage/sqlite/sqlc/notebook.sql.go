// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: notebook.sql

package sqlc

import (
	"context"
)

const createNotebook = `-- name: CreateNotebook :one
INSERT INTO Notebook (
  name
) VALUES (
  ?
)
  RETURNING notebookid, name
`

func (q *Queries) CreateNotebook(ctx context.Context, name string) (Notebook, error) {
	row := q.db.QueryRowContext(ctx, createNotebook, name)
	var i Notebook
	err := row.Scan(&i.Notebookid, &i.Name)
	return i, err
}

const getNotebook = `-- name: GetNotebook :one
SELECT notebookid, name FROM Notebook
WHERE name = ?
`

func (q *Queries) GetNotebook(ctx context.Context, name string) (Notebook, error) {
	row := q.db.QueryRowContext(ctx, getNotebook, name)
	var i Notebook
	err := row.Scan(&i.Notebookid, &i.Name)
	return i, err
}

const getNotebookIdByName = `-- name: GetNotebookIdByName :one
SELECT NotebookID FROM Notebook
WHERE name = ?
`

func (q *Queries) GetNotebookIdByName(ctx context.Context, name string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getNotebookIdByName, name)
	var notebookid int64
	err := row.Scan(&notebookid)
	return notebookid, err
}

const getNotebooks = `-- name: GetNotebooks :many
SELECT notebookid, name FROM Notebook
ORDER BY name
`

func (q *Queries) GetNotebooks(ctx context.Context) ([]Notebook, error) {
	rows, err := q.db.QueryContext(ctx, getNotebooks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Notebook
	for rows.Next() {
		var i Notebook
		if err := rows.Scan(&i.Notebookid, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
