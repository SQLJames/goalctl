package export

import (
	"github.com/sqljames/goalctl/pkg/actions/list"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func Journal() (*resources.Journal, error) {
	var journal resources.Journal

	allNotebooks, err := list.Notebooks()
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return nil, err
	}

	for index, NotebookObject := range allNotebooks {
		entries, err := list.EntriesForNotebook(NotebookObject.Name)
		if err != nil {
			jlogr.Logger.ILog.Error(err, err.Error())

			return nil, err
		}

		allNotebooks[index].Entries = entries
	}

	journal.NoteBooks = allNotebooks

	journal.GoalDetails, err = list.GoalDetails()
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return nil, err
	}
	
	return &journal, nil
}
