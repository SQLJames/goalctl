package actions

import (
	"fmt"
	"os"
	"time"

	"github.com/sqljames/goalctl/pkg/actions"
	"github.com/sqljames/goalctl/pkg/flags"
	"github.com/sqljames/goalctl/pkg/log"
	"github.com/sqljames/goalctl/pkg/printer"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/urfave/cli/v2"
)

var (
	NameGoalFlag *cli.StringFlag = &cli.StringFlag{
		Name:     flags.NameFlagName,
		Usage:    "Name of the Goal.",
		Required: false,
		//Category: "formatting",
		//Aliases: []string{"nb"},
	}
	GoalIDFlag *cli.StringFlag = &cli.StringFlag{
		Name:     flags.GoalIDFlagName,
		Usage:    "ID of the Goal you want to link with the log entry",
		Required: true,
		Aliases:  []string{"g"},
	}

	EntryFlag *cli.StringFlag = &cli.StringFlag{
		Name:     flags.EntryTextFlagName,
		Usage:    "Details of your object",
		Required: false,
		//Category: "formatting",
		//Aliases: []string{"nb"},
	}
)

func ModifyGoal() *cli.Command {
	return &cli.Command{
		Name:   "goal",
		Usage:  "Allows users to modify a target goal",
		Action: actionModifyGoal,
		Flags: []cli.Flag{
			GoalIDFlag,
			flags.PriorityFlag,
			flags.GoalStatusFlag,
			flags.DueDateFlag,
			EntryFlag,
			NameGoalFlag,
			flags.ConfirmFlag,
			flags.OutputFormatFlag,
		},
	}
}

func actionModifyGoal(cliContext *cli.Context) error {
	validateParameters(cliContext)

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

	err := printer.NewPrinter(cliContext).Writer.Write(goal, os.Stdout)
	if err != nil {
		log.Logger.ILog.Warn("issue Printing the data", "function", "ListEntries", "error", err.Error())
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
		log.Logger.ILog.Fatal(&NoModificationError{}, "please check the command or the help section to see what you can modify.")
	}
}

func doWork(confirm bool, goal *resources.Goal) {
	if confirm {
		log.Logger.ILog.Warn("Updating Goal")
		actions.UpdateGoal(goal)
	} else {
		log.Logger.ILog.Warn("Goal Not updated, to replace the information, pass in the --confirm flag.")
	}
}
