package export

import (
	"fmt"
	"os"

	"github.com/sqljames/goalctl/pkg/actions"
	"github.com/sqljames/goalctl/pkg/log"
	"github.com/sqljames/goalctl/pkg/printer"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/urfave/cli/v2"
)

func actionExportJournal(cliContext *cli.Context) error {
	var journal resources.Journal

	Allnotebooks := actions.GetNotebooks()

	for index, NotebookObject := range Allnotebooks {
		entries := actions.GetEntriesForNotebook(NotebookObject.Name)

		Allnotebooks[index].Entries = entries
	}
	
	GoalDetails := actions.GetGoalDetails()

	journal.NoteBooks = Allnotebooks
	journal.GoalDetails = GoalDetails

	err := printer.NewPrinter(cliContext).Writer.Write(resources.Book{Journal: journal}, os.Stdout)
	if err != nil {
		log.Logger.ILog.Warn("issue Printing the data", "function", "CreateGoal", "error", err.Error())
	}

	return fmt.Errorf("printer: %w",err )
}
