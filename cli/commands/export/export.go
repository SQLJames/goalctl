package export

import (
	"github.com/sqljames/goalctl/cli/flags"
	"github.com/sqljames/goalctl/cli/output"
	"github.com/sqljames/goalctl/pkg/actions/export"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
	"github.com/urfave/cli/v2"
)

func New() *cli.Command {
	return &cli.Command{
		Name:    "export",
		Usage:   "Allows users to export objects",
		Aliases: []string{"e"},
		Subcommands: []*cli.Command{
			exportJournal(),
		},
	}
}

func exportJournal() *cli.Command {
	return &cli.Command{
		Name:   "journal",
		Usage:  "exports the specified whole Journal to the specific format",
		Action: actionExportJournal,
		Flags: []cli.Flag{
			flags.OutputFormatFlag,
		},
	}
}

func actionExportJournal(cliContext *cli.Context) error {
	journal, err := export.Journal()
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return err
	}

	output.Output(cliContext.String(flags.OutputFormatFlagName), journal)

	return nil
}
