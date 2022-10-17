package list

import (
	"context"
	"time"

	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/util"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func Goals(filter GoalFilter) ([]*resources.GoalDetail, error) {
	goals, err := GoalDetails()
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return nil, err
	}

	if filter.PastDue {
		pastDueGoals := []*resources.GoalDetail{}

		for index := range goals {
			old, err := expired(&goals[index].Goal.Deadline)
			if err != nil {
				jlogr.Logger.ILog.Error(err, err.Error())

				return nil, err
			}

			if old {
				pastDueGoals = append(pastDueGoals, goals[index])
			}
		}

		goals = pastDueGoals
	}

	return goals, nil
}

func expired(inputDate *string) (bool, error) {
	parsedTime, err := util.StringToTime(*inputDate)
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return false, err
	}

	if parsedTime.Before(time.Now()) {
		return true, nil
	}

	return false, nil
}

func GoalDetails() (details []*resources.GoalDetail, err error) {
	journal := resources.Journal{}

	storagelayer, err := storage.NewVault()
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return nil, err
	}

	goals, err := storagelayer.Storage.GetGoals(context.TODO())
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return nil, err
	}

	allLogEntries, err := storagelayer.Storage.GetLogEntries(context.TODO())
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return nil, err
	}

	allAssociations, err := storagelayer.Storage.GetAssociations(context.TODO())
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return nil, err
	}

	for _, goal := range goals {
		var associations = lookupAssociations(allAssociations, goal.GoalID)

		var logEntries = make([]*resources.LogEntry, len(associations))

		for index, association := range associations {
			logEntries[index] = lookupLogEntry(allLogEntries, association.LogEntryID)
		}

		journal.GoalDetails = append(journal.GoalDetails, &resources.GoalDetail{
			Goal:    *goal,
			Entries: logEntries,
		})
	}

	return journal.GoalDetails, nil
}

func GoalByGoalID(goalID int) (*resources.Goal, error) {
	storagelayer, err := storage.NewVault()
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return nil, err
	}
	
	goal, err := storagelayer.Storage.GetGoalByGoalID(context.TODO(), goalID)
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return nil, err
	}

	return goal, nil
}

func lookupAssociations(entries []*resources.Association, goalID int) (associations []*resources.Association) {
	for _, entry := range entries {
		if entry.GoalID == goalID {
			associations = append(associations, entry)
		}
	}

	return associations
}
