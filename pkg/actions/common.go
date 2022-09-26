package actions

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
)

func GetNotebooks() (notebookList []*resources.Notebook) {
	storagelayer := storage.NewVault()
	notebooks := storagelayer.Storage.GetNotebooks(context.TODO())

	return notebooks
}

func GetEntriesForNotebook(notebookName string) (entries []*resources.LogEntry) {
	storagelayer := storage.NewVault()

	return storagelayer.Storage.GetLogEntryByNotebook(context.TODO(), notebookName)
}

func GetGoalDetails() (details []resources.GoalDetail) {
	storagelayer := storage.NewVault()
	goals := storagelayer.Storage.GetGoals(context.TODO())
	journal := resources.Journal{}
	allLogEntries := storagelayer.Storage.GetLogEntries(context.TODO())
	allAssociations := storagelayer.Storage.GetAssociations(context.TODO())

	for _, goal := range goals {
		var associations = lookupAssociations(allAssociations,goal.GoalID)

		var logEntries = make([]*resources.LogEntry, len(associations))

		for index, association := range associations {
			logEntries[index] = lookupLogEntry(allLogEntries, association.LogEntryID)
		}

		journal.GoalDetails = append(journal.GoalDetails, resources.GoalDetail{
			Goal:    *goal,
			Entries: logEntries,
		})
	}

	return journal.GoalDetails
}

func ListAssociationsByGoalID(goalid int) []*resources.Association {
	storagelayer := storage.NewVault()

	return storagelayer.Storage.GetAssociationsByGoalID(context.TODO(), goalid)
}

func GetLogEntryByLogEntryID(logentryid int) *resources.LogEntry {
	storagelayer := storage.NewVault()

	return storagelayer.Storage.GetLogEntryByLogEntryID(context.TODO(), int64(logentryid))
}

func lookupLogEntry(entries []*resources.LogEntry, logEntryID int) *resources.LogEntry {
	for _, entry := range entries {
		if entry.LogEntryID == int64(logEntryID) {
			return entry
		}
	}

	return nil
}

func lookupAssociations(entries []*resources.Association, goalID int) (associations []*resources.Association) {
	for _, entry := range entries {
		if entry.GoalID == goalID {
			associations = append(associations, entry)
		}
	}

	return associations
}
