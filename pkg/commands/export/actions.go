package export

import (
	"fmt"
	"os"

	"github.com/sqljames/goalctl/pkg/actions"
	"github.com/sqljames/goalctl/pkg/filter"

	"github.com/sqljames/goalctl/pkg/flags"
	"github.com/sqljames/goalctl/pkg/log"
	"github.com/sqljames/goalctl/pkg/printer"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/urfave/cli/v2"
)

func actionExportJournal(cliContext *cli.Context) error {
	var filtered interface{}

	book := GenerateBook()

	if cliContext.IsSet(flags.FilterFlagName) {
		log.Logger.ILog.Warn("filter", "function", "CreateGoal", "filter", cliContext.String(flags.FilterFlagName))
		filtered = filter.Filter(cliContext.String(flags.FilterFlagName), book)
	} else {
		filtered = book
	}

	err := printer.NewPrinter(cliContext).Writer.Write(filtered, os.Stdout)
	if err != nil {
		log.Logger.ILog.Warn("issue Printing the data", "function", "CreateGoal", "error", err.Error())
		err = fmt.Errorf("printer: %w", err)
	}

	return err
}

func GenerateBook() resources.Book {
	var journal resources.Journal

	Allnotebooks := actions.GetNotebooks()

	for index, NotebookObject := range Allnotebooks {
		entries := actions.GetEntriesForNotebook(NotebookObject.Name)

		Allnotebooks[index].Entries = entries
	}

	GoalDetails := actions.GetGoalDetails()

	journal.NoteBooks = Allnotebooks
	journal.GoalDetails = GoalDetails

	return resources.Book{Journal: journal}
}
