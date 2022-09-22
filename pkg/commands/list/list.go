package list

import (
	"github.com/sqljames/goalctl/pkg/flags"
	"github.com/urfave/cli/v2"
)

func New() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Usage:   "Allows users to list objects",
		Aliases: []string{"l"},
		Subcommands: []*cli.Command{
			listNotebooks(),
			listNotebookEntries(),
			listGoalEntries(),
		},
	}
}

func listNotebooks() *cli.Command {
	return &cli.Command{
		Name:   "notebook",
		Usage:  "Prints all the notebooks in your journal",
		Action: actionListNotebooks,
		Flags: []cli.Flag{
			flags.OutputFormatFlag,
		},
	}
}

func listNotebookEntries() *cli.Command {
	return &cli.Command{
		Name:   "entry",
		Usage:  "Prints all the entries in a notebook",
		Action: actionListEntries,
		Flags: []cli.Flag{
			flags.NameNotebookFlag,
			flags.OutputFormatFlag,
		},
	}
}

func listGoalEntries() *cli.Command {
	return &cli.Command{
		Name:   "goal",
		Usage:  "Prints all the goals you have set",
		Action: actionListGoals,
		Flags: []cli.Flag{
			flags.OutputFormatFlag,
		},
	}
}
