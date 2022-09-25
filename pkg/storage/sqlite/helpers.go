package sqlite

import (
	"strings"

	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/storage/sqlite/sqlc"
)

func convertSliceToString(slice []string) string {
	return strings.Join(slice, delimiter)
}

func convertStringToSlice(delimitedString string) []string {
	return strings.Split(delimitedString, delimiter)
}

func convertSqlcLogEntriesToResource(sqlcEntries []*sqlc.LogEntry) (logEntries []*resources.LogEntry) {
	var entries = make([]*resources.LogEntry, len(sqlcEntries))
	for index, entry := range sqlcEntries {
		entries[index] = convertSqlcLogEntryToResource(entry)
	}

	return entries
}

func convertSqlcLogEntryToResource(sqlcEntry *sqlc.LogEntry) (logEnty *resources.LogEntry) {
	return &resources.LogEntry{
		LogEntryID:  sqlcEntry.Logentryid,
		Author:      sqlcEntry.Author.String,
		Tags:        convertStringToSlice(sqlcEntry.Tags.String),
		Entry:       sqlcEntry.Note,
		CreatedDate: sqlcEntry.Createddate,
		Notebookid:  sqlcEntry.Notebookid,
	}
}

func convertSqlcGoalsToResource(sqlcEntries []*sqlc.Goal) []*resources.Goal {
	var goals = make([]*resources.Goal, len(sqlcEntries))
	for index, entry := range sqlcEntries {
		goals[index] = convertSqlcGoalToResource(entry)
	}
	return goals
}

func convertSqlcGoalToResource(sqlcEntry *sqlc.Goal) (goal *resources.Goal) {
	return &resources.Goal{
		GoalID:      int(sqlcEntry.Goalid),
		Author:      sqlcEntry.Author.String,
		Deadline:    sqlcEntry.Duedate.String,
		CreatedDate: sqlcEntry.Createddate,
		Goal:        sqlcEntry.Goal,
		Details:     sqlcEntry.Details,
		Priority:    int(sqlcEntry.Priority),
		Status:      sqlcEntry.Status,
	}
}

func convertSqlcGoalToLogEntriesToResource(sqlcEntries []*sqlc.GoalToLogEntry) []*resources.Association {
	var associations = make([]*resources.Association, len(sqlcEntries))
	for index, entry := range sqlcEntries {
		associations[index] = convertSqlcGoalToLogEntryToResource(entry)
	}
	return associations
}

func convertSqlcGoalToLogEntryToResource(sqlcEntry *sqlc.GoalToLogEntry) (association *resources.Association) {
	return &resources.Association{
		GoalID:     int(sqlcEntry.Goalid),
		LogEntryID: int(sqlcEntry.Logentryid),
	}
}
