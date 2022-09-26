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

	for _, goal := range goals {
		var associations []*resources.Association

		var logEntries = make([]*resources.LogEntry, len(associations))

		associations = ListAssociationsByGoalID(goal.GoalID)

		for index, association := range associations {
			entry := GetLogEntryByLogEntryID(association.LogEntryID)

			logEntries[index] = entry
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
