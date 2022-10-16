-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE LogEntry (
  LogEntryID INTEGER PRIMARY KEY AUTOINCREMENT,
  Author text,
  Tags text,
  Note text NOT NULL,
  CreatedDate text NOT NULL,
  NotebookID INTEGER NOT NULL,
  FOREIGN KEY(NotebookID) REFERENCES Notebook(NotebookID)
);

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
DROP TABLE LogEntry;
-- +migrate StatementEnd