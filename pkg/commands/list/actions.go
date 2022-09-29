package list

import (
	"fmt"
	"os"
	"time"

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
		err = fmt.Errorf("printer: %w", err)
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
		err = fmt.Errorf("printer: %w", err)
	}

	return err
}

func actionListGoals(cliContext *cli.Context) error {
	goals := actions.GetGoalDetails()

	if cliContext.Bool(flags.PastDueFlagName) {
		tempGoalDetails := []*resources.GoalDetail{}

		for index := range goals {
			if expired(&goals[index].Goal.Deadline) {
				tempGoalDetails = append(tempGoalDetails, goals[index])
			}
		}

		goals = tempGoalDetails
	}

	err := printer.NewPrinter(cliContext).Writer.Write(goals, os.Stdout)
	if err != nil {
		log.Logger.ILog.Warn("issue Printing the data", "function", "ListGoals", "error", err.Error())
		err = fmt.Errorf("printer: %w", err)
	}

	return err
}

func expired(inputDate *string) bool {
	parsedTime, err := time.Parse(time.RFC3339, *inputDate)
	if err != nil {
		log.Logger.ILog.Error(err, "date stored in database is not correct", "DateInDatabase", inputDate)

		return false
	}

	if parsedTime.Before(time.Now()) {
		return true
	}

	return false
}

func removeItemFromSlice(slice []*resources.GoalDetail, index int) []*resources.GoalDetail {
	return append(slice[:index], slice[index+1:]...)
}
