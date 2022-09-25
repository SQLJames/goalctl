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

func convertSqlcLogEntriesToResource(sqlcEntries []sqlc.LogEntry) (logEntries []resources.LogEntry) {
	for _, sqlcLogEntry := range sqlcEntries {
		logEntries = append(logEntries, convertSqlcLogEntryToResource(&sqlcLogEntry))
	}
	return logEntries
}

func convertSqlcLogEntryToResource(sqlcEntry *sqlc.LogEntry) (logEnty resources.LogEntry) {
	return resources.LogEntry{
		LogEntryID:  sqlcEntry.Logentryid,
		Author:      sqlcEntry.Author.String,
		Tags:        convertStringToSlice(sqlcEntry.Tags.String),
		Entry:       sqlcEntry.Note,
		CreatedDate: sqlcEntry.Createddate,
		Notebookid:  sqlcEntry.Notebookid,
	}
}

func convertSqlcGoalsToResource(sqlcEntries []sqlc.Goal) (goals []resources.Goal) {
	for index := range sqlcEntries {
		goals = append(goals, convertSqlcGoalToResource(&sqlcEntries[index]))
	}
	return goals
}

func convertSqlcGoalToResource(sqlcEntry *sqlc.Goal) (goal resources.Goal) {
	return resources.Goal{
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

func convertSqlcGoalToLogEntriesToResource(sqlcEntries []sqlc.GoalToLogEntry) (associations []resources.Association) {
	for _, sqlc := range sqlcEntries {
		associations = append(associations, convertSqlcGoalToLogEntryToResource(sqlc))
	}
	return associations
}

func convertSqlcGoalToLogEntryToResource(sqlcEntry sqlc.GoalToLogEntry) (association resources.Association) {

	return resources.Association{
		GoalID:     int(sqlcEntry.Goalid),
		LogEntryID: int(sqlcEntry.Logentryid),
	}
}
