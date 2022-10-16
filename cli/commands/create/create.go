package create

import (
	"github.com/urfave/cli/v2"
)

func New() *cli.Command {
	return &cli.Command{
		Name:    "create",
		Usage:   "Allows users to create objects",
		Aliases: []string{"c"},
		Subcommands: []*cli.Command{
			createNotebook(),
			createLogEntry(),
			createGoalEntry(),
		},
	}
}
