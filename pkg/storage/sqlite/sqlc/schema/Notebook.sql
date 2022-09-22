CREATE TABLE Notebook (
  NotebookID   INTEGER PRIMARY KEY AUTOINCREMENT,
  name text NOT NULL,
  UNIQUE(name)
);
