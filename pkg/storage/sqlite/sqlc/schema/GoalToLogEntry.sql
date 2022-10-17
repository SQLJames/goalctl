-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE GoalToLogEntry (
  GoalID INTEGER NOT NULL,
  LogEntryID INTEGER NOT NULL,
  UNIQUE(GoalID,LogEntryID)
);


-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
DROP TABLE GoalToLogEntry;
-- +migrate StatementEnd