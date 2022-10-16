package modify

import (
	"context"

	"github.com/sqljames/goalctl/pkg/actions/list"
	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/util"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func ModifyGoal(confirm bool, targetGoalID int, modOpts GoalModificationOptions) (*resources.Goal, error) {
	goal, err := list.GetGoalByGoalID(targetGoalID)
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())
		return nil, err
	}
	goal = decodeGoalModificationOptions(goal, modOpts)

	if confirm {
		UpdateGoal(goal)
	}

	return goal, nil
}

func decodeGoalModificationOptions(goal *resources.Goal, modOpts GoalModificationOptions) *resources.Goal {
	if modOpts.GoalDeadline != nil {
		goal.Deadline = util.TimeToString(modOpts.GoalDeadline)
	}

	if modOpts.GoalName != "" {
		goal.Goal = modOpts.GoalName
	}

	if modOpts.GoalDetails != "" {
		goal.Details = modOpts.GoalDetails
	}

	if modOpts.GoalPriority != 0 {
		goal.Priority = modOpts.GoalPriority
	}

	if modOpts.GoalStatus != "" {
		goal.Status = modOpts.GoalStatus
	}

	return goal
}

func UpdateGoal(arg *resources.Goal) error {
	storagelayer, err := storage.NewVault()
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())
		return err
	}
	return storagelayer.Storage.UpdateGoal(context.TODO(), arg)
}
