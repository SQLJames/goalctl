package create

import (
	"context"
	"os"
	"time"

	"github.com/sqljames/goalctl/pkg/log"

	"github.com/sqljames/goalctl/pkg/flags"
	"github.com/sqljames/goalctl/pkg/printer"
	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/urfave/cli/v2"
)

type emptryStringErr struct{}

func (empty *emptryStringErr) Error() string {
	return "string can not be empty"
}
func actionCreateNotebook(cliContext *cli.Context) error {
    storagelayer := storage.NewVault()
	if cliContext.String(flags.NameFlagName) == "" {
		log.Logger.Error(&emptryStringErr{}, "Error creating notebook", "function", "actionCreateNotebook")
		return &emptryStringErr{}
	}

	row, err := storagelayer.CreateNotebook(context.TODO(), cliContext.String(flags.NameFlagName))
	if err != nil {
		return err
	}
	err = printer.NewPrinter(cliContext).Write(row, os.Stdout)
	if err != nil {
		log.Logger.Warn("issue Printing the data", "function", "CreateNotebook", "error", err.Error())
	}
	return err
}

func actionCreateLogEntry(cliContext *cli.Context) error {
	storagelayer := storage.NewVault()	
	le := resources.NewLogEntry(cliContext.String(flags.EntryTextFlagName), cliContext.StringSlice(flags.TagsFlagName))

	row, err := storagelayer.CreateLogEntry(context.TODO(), le, cliContext.String(flags.NameFlagName))
	if err != nil {
		return err
	}
	err = printer.NewPrinter(cliContext).Write(row, os.Stdout)
	if err != nil {
		log.Logger.Warn("issue Printing the data", "function", "CreateLogEntry", "error", err.Error())
	}
	return err
}

func actionCreateGoal(cliContext *cli.Context) error {
	storagelayer := storage.NewVault()
	goal := resources.NewGoal(
		cliContext.String(flags.NameFlagName),
		cliContext.Timestamp(flags.DueDateFlagName).UTC().Format(time.RFC3339),
		cliContext.String(flags.EntryTextFlagName),
		cliContext.Int(flags.PriorityFlagName))

	row, err := storagelayer.CreateGoal(context.TODO(), goal)
	if err != nil {
		return err
	}
	err = printer.NewPrinter(cliContext).Write(row, os.Stdout)
	if err != nil {
		log.Logger.Warn("issue Printing the data", "function", "CreateGoal", "error", err.Error())
	}
	return err
}
