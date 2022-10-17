-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE Goal (
  GoalID INTEGER PRIMARY KEY AUTOINCREMENT,
  DueDate text,
  Author text,
  CreatedDate text NOT NULL,
  Goal text NOT NULL,--Goal is what you would like to accomplish
  Details text NOT NULL,--Details of the goal
  Priority INTEGER NOT NULL,
  Status text NOT NULL
);
-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
DROP TABLE Goal;
-- +migrate StatementEnd