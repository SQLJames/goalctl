package create

import (
	"github.com/sqljames/goalctl/cli/flags"
	"github.com/sqljames/goalctl/cli/output"
	"github.com/sqljames/goalctl/pkg/actions/create"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
	"github.com/urfave/cli/v2"
)

func createLogEntry() *cli.Command {
	return &cli.Command{
		Name:   "entry",
		Usage:  "creates a new entry in a specific notebook.",
		Action: actionCreateLogEntry,
		Flags: []cli.Flag{
			flags.NameNotebookFlag,
			flags.EntryFlag,
			flags.TagsFlag,
		},
	}
}

func actionCreateLogEntry(cliContext *cli.Context) error {
	logEntry := create.NewLogEntry{
		LogEntry:     cliContext.String(flags.EntryTextFlagName),
		NotebookName: cliContext.String(flags.NameFlagName),
		Tags:         cliContext.StringSlice(flags.TagsFlagName),
	}
	
	result, err := create.LogEntry(logEntry)
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return err
	}

	output.Output(cliContext.String(flags.OutputFormatFlagName), result)

	return nil
}
