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
	notebookList, err := actions.GetNotebooks()
	if err != nil {
		log.Logger.Error(err, "Error getting notebooks")
	}
	err = printer.NewPrinter(cliContext).Write(notebookList, os.Stdout)
	if err != nil {
		log.Logger.Warn("issue Printing the data", "function", "ListNotebooks", "error", err.Error())
	}
	return err
}

func actionListEntries(cliContext *cli.Context) error {
	NotebookEntries, err := actions.GetEntriesForNotebook(cliContext.String(flags.NameFlagName))
	if err != nil {
		log.Logger.Error(err, err.Error())
		return err
	}

	notebook := resources.Notebook{
		Name:    cliContext.String(flags.NameFlagName),
		Entries: NotebookEntries,
	}
	log.Logger.Trace("Created notebook resource")
	err = printer.NewPrinter(cliContext).Write(notebook, os.Stdout)
	if err != nil {
		log.Logger.Warn("issue Printing the data", "function", "ListEntries", "error", err.Error())
	}
	return err
}

func actionListGoals(cliContext *cli.Context) error {
	goals, err := actions.GetGoalDetails()
	if err != nil {
		log.Logger.Error(err, err.Error())
		return err
	}
	err = printer.NewPrinter(cliContext).Write(goals, os.Stdout)
	if err != nil {
		log.Logger.Warn("issue Printing the data", "function", "ListGoals", "error", err.Error())
	}
	return err
}
