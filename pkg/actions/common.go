package actions

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
)

func GetNotebooks() (notebookList []*resources.Notebook, err error) {
	storagelayer := storage.NewVault()
	notebooks, err := storagelayer.GetNotebooks(context.TODO())
	if err != nil {
		return nil, err
	}

	return notebooks, nil
}

func GetEntriesForNotebook(notebookName string) (entries []*resources.LogEntry, err error) {
	storagelayer := storage.NewVault()
	return storagelayer.GetLogEntryByNotebook(context.TODO(), notebookName)
}

func GetGoalDetails() (details []resources.GoalDetail, err error) {
	storagelayer := storage.NewVault()
	goals, err := storagelayer.GetGoals(context.TODO())
	var journal resources.Journal
	for _, goal := range goals {
		var associations []*resources.Association
		associations, err = ListAssociationsByGoalID(goal.GoalID)
		if err != nil {
			return nil, err
		}

		var logEntries = make([]*resources.LogEntry, len(associations))
		for _, association := range associations {
			var entry *resources.LogEntry
			entry, err = GetLogEntryByLogEntryID(association.LogEntryID)
			if err != nil {
				return nil, err
			}
			logEntries = append(logEntries, entry)
		}
		journal.GoalDetails = append(journal.GoalDetails, resources.GoalDetail{
			Goal:    *goal,
			Entries: logEntries,
		})
	}
	return journal.GoalDetails, err
}

func ListAssociationsByGoalID(goalid int) ([]*resources.Association, error) {
	storagelayer := storage.NewVault()
	return storagelayer.GetAssociationsByGoalID(context.TODO(), goalid)
}

func GetLogEntryByLogEntryID(logentryid int) (*resources.LogEntry, error) {
	storagelayer := storage.NewVault()
	return storagelayer.GetLogEntryByLogEntryID(context.TODO(), int64(logentryid))
}
