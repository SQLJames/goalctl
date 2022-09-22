package link

import (
	"github.com/sqljames/goalctl/pkg/flags"
	"github.com/urfave/cli/v2"
)

func New() *cli.Command {
	return &cli.Command{
		Name:    "link",
		Usage:   "Allows users to link goals to log entries",
		Aliases: []string{"l"},
		Action:  actionLink,
		Flags: []cli.Flag{
			flags.GoalIDFlag,
			flags.LogEntryIDFlag,
		},
	}
}
