package modify

import (
	"github.com/sqljames/goalctl/cli/flags"
	"github.com/sqljames/goalctl/cli/output"
	"github.com/sqljames/goalctl/pkg/actions/modify"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
	"github.com/urfave/cli/v2"
)

func modifyGoal() *cli.Command {
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
	modificationDetails := modify.GoalModificationOptions{}
	if cliContext.IsSet(flags.DueDateFlagName) {
		modificationDetails.GoalDeadline = cliContext.Timestamp(flags.DueDateFlagName)
	}

	if cliContext.IsSet(flags.NameFlagName) {
		modificationDetails.GoalName = cliContext.String(flags.NameFlagName)
	}

	if cliContext.IsSet(flags.EntryTextFlagName) {
		modificationDetails.GoalDetails = cliContext.String(flags.EntryTextFlagName)
	}

	if cliContext.IsSet(flags.PriorityFlagName) {
		modificationDetails.GoalPriority = cliContext.Int(flags.PriorityFlagName)
	}

	if cliContext.IsSet(flags.GoalStatusFlagName) {
		modificationDetails.GoalStatus = cliContext.String(flags.GoalStatusFlagName)
	}

	goal, err := modify.Goal(cliContext.Bool(flags.ConfirmFlagName), cliContext.Int(flags.GoalIDFlagName), modificationDetails)
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return err
	}

	output.Output(cliContext.String(flags.OutputFormatFlagName), goal)

	confirmationWarning(cliContext.Bool(flags.ConfirmFlagName))

	return nil
}
