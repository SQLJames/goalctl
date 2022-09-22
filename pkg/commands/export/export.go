package export

import (
	"github.com/sqljames/goalctl/pkg/flags"
	"github.com/urfave/cli/v2"
)

func New() *cli.Command {
	return &cli.Command{
		Name:    "export",
		Usage:   "Allows users to export objects",
		Aliases: []string{"e"},
		Subcommands: []*cli.Command{
			exportNotebook(),
		},
	}
}

func exportNotebook() *cli.Command {
	return &cli.Command{
		Name:   "journal",
		Usage:  "exports the specified whole Journal to the specific format",
		Action: actionExportJournal,
		Flags: []cli.Flag{
			flags.OutputFormatFlag,
		},
	}
}
