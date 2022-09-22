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

func actionListNotebooks(c *cli.Context) error {
	NotebookList, err := actions.GetNotebooks()
	if err != nil {
		log.Logger.Error(err, "Error getting notebooks")
	}
	printer := printer.NewPrinter(c)
	err = printer.Write(NotebookList, os.Stdout)
	if err != nil {
		log.Logger.Warn("issue Printing the data", "function", "ListNotebooks", "error", err.Error())
	}
	return err
}

func actionListEntries(c *cli.Context) error {
	NotebookEntries, err := actions.GetEntriesForNotebook(c.String(flags.NameFlagName))
	if err != nil {
		log.Logger.Error(err, err.Error())
		return err
	}

	NB := resources.Notebook{
		Name:    c.String(flags.NameFlagName),
		Entries: NotebookEntries,
	}
	log.Logger.Trace("Created notebook resource")
	printer := printer.NewPrinter(c)
	err = printer.Write(NB, os.Stdout)
	if err != nil {
		log.Logger.Warn("issue Printing the data", "function", "ListEntries", "error", err.Error())
	}
	return err
}

func actionListGoals(c *cli.Context) error {
	goals, err := actions.GetGoalDetails()
	if err != nil {
		log.Logger.Error(err, err.Error())
		return err
	}
	printer := printer.NewPrinter(c)
	err = printer.Write(goals, os.Stdout)
	if err != nil {
		log.Logger.Warn("issue Printing the data", "function", "ListGoals", "error", err.Error())
	}
	return err
}
