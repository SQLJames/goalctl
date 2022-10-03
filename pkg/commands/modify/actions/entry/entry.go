package entry

import (
	"fmt"
	"os"

	common "github.com/sqljames/goalctl/pkg/actions"
	"github.com/sqljames/goalctl/pkg/commands/modify/actions"
	"github.com/sqljames/goalctl/pkg/flags"
	"github.com/sqljames/goalctl/pkg/printer"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
	"github.com/urfave/cli/v2"
)

func ModifyEntry() *cli.Command {
	return &cli.Command{
		Name:   "entry",
		Usage:  "Allows users to modify a target entry",
		Action: actionModifyEntry,
		Flags: []cli.Flag{
			actions.LogEntryIDFlag,
			actions.EntryFlag,
			flags.TagsFlag,
			actions.NotebookIDFlag,
			flags.ConfirmFlag,
			flags.OutputFormatFlag,
		},
	}
}

func actionModifyEntry(cliContext *cli.Context) error {
	validateParameters(cliContext)

	logEntry := common.GetLogEntryByLogEntryID(cliContext.Int(flags.LogEntryIDFlagName))

	if cliContext.IsSet(flags.NotebookIDFlagName) {
		logEntry.Notebookid = cliContext.Int64(flags.NotebookIDFlagName)
	}

	if cliContext.IsSet(flags.EntryTextFlagName) {
		logEntry.Entry = cliContext.String(flags.EntryTextFlagName)
	}

	if cliContext.IsSet(flags.TagsFlagName) {
		logEntry.Tags = cliContext.StringSlice(flags.TagsFlagName)
	}

	err := printer.NewPrinter(cliContext).Writer.Write(logEntry, os.Stdout)
	if err != nil {
		jlogr.Logger.ILog.Warn("issue Printing the data", "function", "ListEntries", "error", err.Error())
		err = fmt.Errorf("printer: %w", err)
	}

	doWork(cliContext.Bool(flags.ConfirmFlagName), logEntry)

	return nil
}

func validateParameters(cliContext *cli.Context) {
	if !cliContext.IsSet(flags.EntryTextFlagName) &&
		!cliContext.IsSet(flags.NotebookIDFlagName) &&
		!cliContext.IsSet(flags.TagsFlagName) {
		jlogr.Logger.ILog.Fatal(&actions.NoModificationError{}, "please check the command or the help section to see what you can modify.")
	}
}

func doWork(confirm bool, logEntry *resources.LogEntry) {
	if confirm {
		jlogr.Logger.ILog.Warn("Updating")
		common.UpdateLogEntry(logEntry)
	} else {
		jlogr.Logger.ILog.Warn("Not updated, to replace the information, pass in the --confirm flag.")
	}
}
