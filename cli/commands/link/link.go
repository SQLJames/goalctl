package link

import (
	"github.com/sqljames/goalctl/cli/flags"
	"github.com/sqljames/goalctl/pkg/actions/link"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
	"github.com/urfave/cli/v2"
)

func New() *cli.Command {
	return &cli.Command{
		Name:    "link",
		Usage:   "Allows users to link goals to log entries",
		Aliases: []string{"l"},
		Action:  linkAction,
		Flags: []cli.Flag{
			flags.GoalIDFlag,
			flags.LogEntryIDFlag,
		},
	}
}

func linkAction(cliContext *cli.Context) error {
	err := link.Link(cliContext.StringSlice(flags.LogEntryIDFlagName), cliContext.StringSlice(flags.GoalIDFlagName))
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())
		return err
	}

	return nil
}
