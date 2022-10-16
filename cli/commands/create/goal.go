package create

import (
	"github.com/sqljames/goalctl/cli/flags"
	"github.com/sqljames/goalctl/cli/output"
	"github.com/sqljames/goalctl/pkg/actions/create"
	"github.com/urfave/cli/v2"
)

func createGoalEntry() *cli.Command {
	return &cli.Command{
		Name:   "goal",
		Usage:  "creates a new goal on the Goals table",
		Action: actionCreateGoal,
		Flags: []cli.Flag{
			flags.NameGoalFlag,
			flags.EntryFlag,
			flags.PriorityFlag,
			flags.DueDateFlag,
		},
	}
}

func actionCreateGoal(cliContext *cli.Context) error {
	newGoal := create.NewGoal{
		Goal:     cliContext.String(flags.NameFlagName),
		DueDate:  cliContext.Timestamp(flags.DueDateFlagName),
		Details:  cliContext.String(flags.EntryTextFlagName),
		Priority: cliContext.Int(flags.PriorityFlagName),
	}
	result := create.CreateGoal(newGoal)
	output.Output(cliContext.String(flags.OutputFormatFlagName), result)

	return nil
}
