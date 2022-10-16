package modify

import (
	"github.com/sqljames/goalctl/cli/flags"
	"github.com/sqljames/goalctl/cli/output"
	"github.com/sqljames/goalctl/pkg/actions/modify"
	"github.com/urfave/cli/v2"
)

func modifyEntry() *cli.Command {
	return &cli.Command{
		Name:   "entry",
		Usage:  "Allows users to modify a target entry",
		Action: actionModifyEntry,
		Flags: []cli.Flag{
			LogEntryIDFlag,
			EntryFlag,
			flags.TagsFlag,
			NotebookIDFlag,
			flags.ConfirmFlag,
			flags.OutputFormatFlag,
		},
	}
}

func actionModifyEntry(cliContext *cli.Context) error {
	modificationDetails := decodeModificationRequest(cliContext)
	logEntry := modify.ModifyEntry(cliContext.Bool(flags.ConfirmFlagName), cliContext.Int(flags.LogEntryIDFlagName), modificationDetails)

	output.Output(cliContext.String(flags.OutputFormatFlagName), logEntry)

	confirmationWarning(cliContext.Bool(flags.ConfirmFlagName))
	return nil
}

func decodeModificationRequest(cliContext *cli.Context) modify.EntryModificationOptions {
	modificationDetails := modify.EntryModificationOptions{}
	if cliContext.IsSet(flags.NotebookIDFlagName) {
		modificationDetails.TargetNotebookID = cliContext.Int64(flags.NotebookIDFlagName)
	}

	if cliContext.IsSet(flags.EntryTextFlagName) {
		modificationDetails.EntryDetails = cliContext.String(flags.EntryTextFlagName)
	}

	if cliContext.IsSet(flags.TagsFlagName) {
		modificationDetails.EntryTags = cliContext.StringSlice(flags.TagsFlagName)
	}
	return modificationDetails
}
