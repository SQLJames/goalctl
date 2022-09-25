package actions

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
)

func GetNotebooks() (notebookList []resources.Notebook, err error) {
	storagelayer, err := storage.NewStorageLayer()
	if err != nil {
		return nil, err
	}
	notebooks, err := storagelayer.GetNotebooks(context.TODO())
	if err != nil {
		return nil, err
	}

	return notebooks, nil
}

func GetEntriesForNotebook(notebookName string) (entries []resources.LogEntry, err error) {
	storagelayer, err := storage.NewStorageLayer()
	if err != nil {
		return nil, err
	}
	return storagelayer.GetLogEntryByNotebook(context.TODO(), notebookName)

}

func GetGoalDetails() (details []resources.GoalDetail, err error) {
	storagelayer, err := storage.NewStorageLayer()
	if err != nil {
		return nil, err
	}
	goals, err := storagelayer.GetGoals(context.TODO())
	var journal resources.Journal
	for _, goal := range goals {
		logEntries := []resources.LogEntry{}

		associations, err := ListAssociationsByGoalId(goal.GoalID)
		if err != nil {
			return nil, err
		}
		for _, association := range associations {
			entry, err := GetLogEntryByLogEntryID(association.LogEntryID)
			if err != nil {
				return nil, err
			}
			logEntries = append(logEntries, entry)
		}
		journal.GoalDetails = append(journal.GoalDetails, resources.GoalDetail{
			Goal:    goal,
			Entries: logEntries,
		})
	}
	return journal.GoalDetails, err
}

func ListAssociationsByGoalId(goalid int) ([]resources.Association, error) {
	storagelayer, err := storage.NewStorageLayer()
	if err != nil {
		return nil, err
	}
	return storagelayer.GetAssociationsByGoalID(context.TODO(), goalid)
}

func GetLogEntryByLogEntryID(logentryid int) (resources.LogEntry, error) {
	storagelayer, err := storage.NewStorageLayer()
	if err != nil {
		return resources.LogEntry{}, err
	}
	return storagelayer.GetLogEntryByLogEntryID(context.TODO(), int64(logentryid))
}
