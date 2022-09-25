-- name: GetNotebook :one
SELECT * FROM Notebook
WHERE name = ?;

-- name: GetNotebooks :many
SELECT * FROM Notebook
ORDER BY name;

-- name: GetNotebookIDByName :one
SELECT NotebookID FROM Notebook
WHERE name = ?;

-- name: CreateNotebook :one
INSERT INTO Notebook (
  name
) VALUES (
  ?
)
  RETURNING *;
