package export

import (
	"github.com/sqljames/goalctl/pkg/actions/list"
	"github.com/sqljames/goalctl/pkg/storage/resources"
)

func ExportJournal() resources.Journal {
	var journal resources.Journal

	Allnotebooks := list.GetNotebooks()

	for index, NotebookObject := range Allnotebooks {
		entries := list.GetEntriesForNotebook(NotebookObject.Name)
		Allnotebooks[index].Entries = entries
	}

	journal.NoteBooks = Allnotebooks
	journal.GoalDetails = list.GetGoalDetails()

	return journal
}
