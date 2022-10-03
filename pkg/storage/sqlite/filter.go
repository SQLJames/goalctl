package sqlite

import (
	"context"

	"github.com/sqljames/goalctl/pkg/log"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	compiledsqlc "github.com/sqljames/goalctl/pkg/storage/sqlite/sqlc"
	"github.com/yiplee/sqlc"
)

func (sl Repository) FilterLogEntries(ctx context.Context, arg *resources.LogEntry) (logEntries []*resources.LogEntry) {
	query := compiledsqlc.New(sqlc.Wrap(&sl.db))
	sqlcLogEntries, err := query.GetLogEntries(sqlc.Build(ctx, func(builder *sqlc.Builder) {
		if arg.Notebookid != 0 {
			builder.Where("Notebookid = $1", arg.Notebookid)
		}

		builder.Order("LogEntryID DESC")
	}))

	if err != nil {
		log.Logger.ILog.Fatal(err, "error running query")
	}

	return convertSqlcLogEntriesToResource(sqlcLogEntries)
}

func (sl Repository) FilterGetGoals(ctx context.Context,arg *resources.Goal)[]*resources.Goal {
	query := compiledsqlc.New(sqlc.Wrap(&sl.db))
	sqlcLogEntries, err := query.GetGoals(sqlc.Build(ctx, func(builder *sqlc.Builder) {
		if arg.Deadline != "" {
			builder.Where("duedate < $1", arg.Deadline)
		}

		builder.Order("GoalID DESC")
	}))
	
	if err != nil {
		log.Logger.ILog.Fatal(err, "error running query")
	}

	return convertSqlcGoalsToResource(sqlcLogEntries)
}