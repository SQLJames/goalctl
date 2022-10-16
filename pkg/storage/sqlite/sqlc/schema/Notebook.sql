-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE Notebook (
  NotebookID   INTEGER PRIMARY KEY AUTOINCREMENT,
  name text NOT NULL,
  UNIQUE(name)
);
-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
DROP TABLE Notebook;
-- +migrate StatementEnd