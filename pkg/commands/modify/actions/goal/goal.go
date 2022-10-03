package goal

import (
	"fmt"
	"os"

	common "github.com/sqljames/goalctl/pkg/actions"
	"github.com/sqljames/goalctl/pkg/commands/modify/actions"
	"github.com/sqljames/goalctl/pkg/flags"
	"github.com/sqljames/goalctl/pkg/printer"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/util"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
	"github.com/urfave/cli/v2"
)

func ModifyGoal() *cli.Command {
	return &cli.Command{
		Name:   "goal",
		Usage:  "Allows users to modify a target goal",
		Action: actionModifyGoal,
		Flags: []cli.Flag{
			actions.GoalIDFlag,
			flags.PriorityFlag,
			flags.GoalStatusFlag,
			flags.DueDateFlag,
			actions.EntryFlag,
			actions.NameGoalFlag,
			flags.ConfirmFlag,
			flags.OutputFormatFlag,
		},
	}
}

func actionModifyGoal(cliContext *cli.Context) error {
	validateParameters(cliContext)

	goal := common.GetGoalByGoalID(cliContext.Int(flags.GoalIDFlagName))

	if cliContext.IsSet(flags.DueDateFlagName) {
		goal.Deadline = util.TimeToString(cliContext.Timestamp(flags.DueDateFlagName))
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

	err := printer.NewPrinter(cliContext).Writer.Write(goal, os.Stdout)
	if err != nil {
		jlogr.Logger.ILog.Warn("issue Printing the data", "function", "ListEntries", "error", err.Error())
		err = fmt.Errorf("printer: %w", err)
	}

	doWork(cliContext.Bool(flags.ConfirmFlagName), goal)

	return nil
}

func validateParameters(cliContext *cli.Context) {
	if !cliContext.IsSet(flags.DueDateFlagName) &&
		!cliContext.IsSet(flags.NameFlagName) &&
		!cliContext.IsSet(flags.EntryTextFlagName) &&
		!cliContext.IsSet(flags.PriorityFlagName) &&
		!cliContext.IsSet(flags.GoalStatusFlagName) {
		jlogr.Logger.ILog.Fatal(&actions.NoModificationError{}, "please check the command or the help section to see what you can modify.")
	}
}

func doWork(confirm bool, goal *resources.Goal) {
	if confirm {
		jlogr.Logger.ILog.Warn("Updating")
		common.UpdateGoal(goal)
	} else {
		jlogr.Logger.ILog.Warn("Not updated, to replace the information, pass in the --confirm flag.")
	}
}
