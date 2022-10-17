package modify

import (
	"github.com/sqljames/goalctl/pkg/util/jlogr"
	"github.com/urfave/cli/v2"
)

func New() *cli.Command {
	return &cli.Command{
		Name:    "modify",
		Usage:   "Allows users to modify objects",
		Aliases: []string{"m"},
		Subcommands: []*cli.Command{
			modifyGoal(),
			modifyEntry(),
		},
	}
}

func confirmationWarning(confirmation bool) {
	if !confirmation {
		jlogr.Logger.ILog.Warn("Not updated, to replace the information, pass in the --confirm flag.")
	}
}
