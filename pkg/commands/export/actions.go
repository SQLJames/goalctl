package export

import (
	"os"

	"github.com/sqljames/goalctl/pkg/actions"
	"github.com/sqljames/goalctl/pkg/log"
	"github.com/sqljames/goalctl/pkg/printer"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/urfave/cli/v2"
)

func actionExportJournal(c *cli.Context) error {
	var journal resources.Journal
	Allnotebooks, err := actions.GetNotebooks()
	if err != nil {
		return err
	}
	for index, NotebookObject := range Allnotebooks {
		entries, err := actions.GetEntriesForNotebook(NotebookObject.Name)
		if err != nil {
			return err
		}
		Allnotebooks[index].Entries = entries
	}
	GoalDetails, err := actions.GetGoalDetails()
	if err != nil {
		return err
	}
	journal.NoteBooks = Allnotebooks
	journal.GoalDetails = GoalDetails

	printer := printer.NewPrinter(c)
	err = printer.Write(resources.Book{Journal: journal}, os.Stdout)
	if err != nil {
		log.Logger.Warn("issue Printing the data", "function", "CreateGoal", "error", err.Error())
	}
	return err
}
