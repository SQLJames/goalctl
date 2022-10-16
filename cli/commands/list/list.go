package list

import (
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
