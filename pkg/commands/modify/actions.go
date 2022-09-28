package modify

import (
	"fmt"
	"time"

	"github.com/sqljames/goalctl/pkg/actions"
	"github.com/sqljames/goalctl/pkg/flags"
	"github.com/sqljames/goalctl/pkg/log"
	"github.com/urfave/cli/v2"
)

func actionModifyGoal(cliContext *cli.Context) error {
	log.Logger.ILog.Debug("data", "goalid", cliContext.Int(flags.GoalIDFlagName))

	if !cliContext.IsSet(flags.DueDateFlagName) &&
		!cliContext.IsSet(flags.NameFlagName) &&
		!cliContext.IsSet(flags.EntryTextFlagName) &&
		!cliContext.IsSet(flags.PriorityFlagName) &&
		!cliContext.IsSet(flags.GoalStatusFlagName) {
		log.Logger.ILog.Fatal(fmt.Errorf("you must modify something"), "You didn't pass in any parameters you wanted modified")
	}

	goal := actions.GetGoalByGoalID(cliContext.Int(flags.GoalIDFlagName))

	if cliContext.IsSet(flags.DueDateFlagName) {
		goal.Deadline = cliContext.Timestamp(flags.DueDateFlagName).UTC().Format(time.RFC3339)
	}

	if cliContext.IsSet(flags.NameFlagName) {
		goal.Goal = cliContext.String(flags.NameFlagName)
	}

	if cliContext.IsSet(flags.EntryTextFlagName) {
		goal.Details = cliContext.String(flags.EntryTextFlagName)
	}

	if cliContext.IsSet(flags.PriorityFlagName) {
		goal.Priority = cliContext.Int(flags.PriorityFlagName)
	}

	if cliContext.IsSet(flags.GoalStatusFlagName) {
		goal.Status = cliContext.String(flags.GoalStatusFlagName)
	}

	actions.UpdateGoal(goal)

	return nil
}
