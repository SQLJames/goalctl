package list

import (
	"github.com/sqljames/goalctl/cli/flags"
	"github.com/sqljames/goalctl/cli/output"
	"github.com/sqljames/goalctl/pkg/actions/list"
	"github.com/urfave/cli/v2"
)

func listGoalEntries() *cli.Command {
	return &cli.Command{
		Name:   "goal",
		Usage:  "Prints all the goals you have set",
		Action: ListGoals,
		Flags: []cli.Flag{
			flags.OutputFormatFlag,
			PastDueFlag,
		},
	}
}

func ListGoals(cliContext *cli.Context) error {
	filter := list.GoalFilter{}
	if cliContext.Bool(flags.PastDueFlagName) {
		filter.PastDue = true
	}

	goals := list.ListGoals(filter)

	output.Output(cliContext.String(flags.OutputFormatFlagName), goals)

	return nil
}
