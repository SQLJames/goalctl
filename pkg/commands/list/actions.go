package list

import (
	"os"

	"github.com/sqljames/goalctl/pkg/actions"
	"github.com/sqljames/goalctl/pkg/flags"
	"github.com/sqljames/goalctl/pkg/log"
	"github.com/sqljames/goalctl/pkg/printer"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/urfave/cli/v2"
)

func actionListNotebooks(cliContext *cli.Context) error {
	notebookList := actions.GetNotebooks()

	err := printer.NewPrinter(cliContext).Writer.Write(notebookList, os.Stdout)
	if err != nil {
		log.Logger.ILog.Warn("issue Printing the data", "function", "ListNotebooks", "error", err.Error())
	}
	return err
}

func actionListEntries(cliContext *cli.Context) error {
	NotebookEntries := actions.GetEntriesForNotebook(cliContext.String(flags.NameFlagName))

	notebook := resources.Notebook{
		Name:    cliContext.String(flags.NameFlagName),
		Entries: NotebookEntries,
	}

	err := printer.NewPrinter(cliContext).Writer.Write(notebook, os.Stdout)
	if err != nil {
		log.Logger.ILog.Warn("issue Printing the data", "function", "ListEntries", "error", err.Error())
	}
	return err
}

func actionListGoals(cliContext *cli.Context) error {
	goals := actions.GetGoalDetails()
	err := printer.NewPrinter(cliContext).Writer.Write(goals, os.Stdout)
	if err != nil {
		log.Logger.ILog.Warn("issue Printing the data", "function", "ListGoals", "error", err.Error())
	}
	return err
}
