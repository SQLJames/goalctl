package modify

import (
	"github.com/sqljames/goalctl/pkg/commands/modify/actions/entry"
	"github.com/sqljames/goalctl/pkg/commands/modify/actions/goal"
	"github.com/urfave/cli/v2"
)

func New() *cli.Command {
	return &cli.Command{
		Name:    "modify",
		Usage:   "Allows users to modify objects",
		Aliases: []string{"m"},
		Subcommands: []*cli.Command{
			goal.ModifyGoal(),
			entry.ModifyEntry(),
		},
	}
}
