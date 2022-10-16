package export

import (
	"github.com/sqljames/goalctl/pkg/actions/list"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func ExportJournal() (*resources.Journal, error) {
	var journal resources.Journal

	Allnotebooks, err := list.GetNotebooks()
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())
		return nil, err
	}
	for index, NotebookObject := range Allnotebooks {
		entries, err := list.GetEntriesForNotebook(NotebookObject.Name)
		if err != nil {
			jlogr.Logger.ILog.Error(err, err.Error())
			return nil, err
		}
		Allnotebooks[index].Entries = entries
	}

	journal.NoteBooks = Allnotebooks
	journal.GoalDetails, err = list.GetGoalDetails()
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())
		return nil, err
	}
	return &journal, nil
}
