package modify

import (
	"github.com/sqljames/goalctl/pkg/commands/modify/actions"
	"github.com/urfave/cli/v2"
)

func New() *cli.Command {
	return &cli.Command{
		Name:    "modify",
		Usage:   "Allows users to modify objects",
		Aliases: []string{"m"},
		Subcommands: []*cli.Command{
			actions.ModifyGoal(),
		},
	}
}
