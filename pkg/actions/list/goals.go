package list

import (
	"context"
	"time"

	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/util"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func ListGoals(filter GoalFilter) []*resources.GoalDetail {
	goals := GetGoalDetails()

	if filter.PastDue {
		pastDueGoals := []*resources.GoalDetail{}

		for index := range goals {
			if expired(&goals[index].Goal.Deadline) {
				pastDueGoals = append(pastDueGoals, goals[index])
			}
		}

		goals = pastDueGoals
	}

	return goals
}

func expired(inputDate *string) bool {
	parsedTime, err := util.StringToTime(*inputDate)
	if err != nil {
		jlogr.Logger.ILog.Error(err, "date stored in database is not correct", "DateInDatabase", &inputDate)

		return false
	}

	if parsedTime.Before(time.Now()) {
		return true
	}

	return false
}

func GetGoalDetails() (details []*resources.GoalDetail) {
	storagelayer, err := storage.NewVault()
	if err != nil {
		jlogr.Logger.ILog.Fatal(err, err.Error())
		return
	}
	goals := storagelayer.Storage.GetGoals(context.TODO())
	journal := resources.Journal{}
	allLogEntries := storagelayer.Storage.GetLogEntries(context.TODO())
	allAssociations := storagelayer.Storage.GetAssociations(context.TODO())

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

	return journal.GoalDetails
}

func GetGoalByGoalID(goalID int) *resources.Goal {
	storagelayer, err := storage.NewVault()
	if err != nil {
		jlogr.Logger.ILog.Fatal(err, err.Error())
	}

	return storagelayer.Storage.GetGoalByGoalID(context.TODO(), goalID)
}

func lookupAssociations(entries []*resources.Association, goalID int) (associations []*resources.Association) {
	for _, entry := range entries {
		if entry.GoalID == goalID {
			associations = append(associations, entry)
		}
	}

	return associations
}
