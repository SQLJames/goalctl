package create

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/sqljames/goalctl/pkg/log"

	"github.com/sqljames/goalctl/pkg/flags"
	"github.com/sqljames/goalctl/pkg/printer"
	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/urfave/cli/v2"
)

type emptryStringError struct{}

func (empty *emptryStringError) Error() string {
	return "string can not be empty"
}
func actionCreateNotebook(cliContext *cli.Context) error {
	if cliContext.String(flags.NameFlagName) == "" {
		log.Logger.ILog.Error(&emptryStringError{}, "Error creating notebook", "function", "actionCreateNotebook")

		return &emptryStringError{}
	}

	storagelayer := storage.NewVault()
	row := storagelayer.Storage.CreateNotebook(context.TODO(), cliContext.String(flags.NameFlagName))

	err := printer.NewPrinter(cliContext).Writer.Write(row, os.Stdout)
	if err != nil {
		log.Logger.ILog.Warn("issue Printing the data", "function", "CreateNotebook", "error", err.Error())
	}
	
	return fmt.Errorf("printer: %w", err)
}

func actionCreateLogEntry(cliContext *cli.Context) error {
	storagelayer := storage.NewVault()
	le := resources.NewLogEntry(cliContext.String(flags.EntryTextFlagName), cliContext.StringSlice(flags.TagsFlagName))

	row := storagelayer.Storage.CreateLogEntry(context.TODO(), le, cliContext.String(flags.NameFlagName))

	err := printer.NewPrinter(cliContext).Writer.Write(row, os.Stdout)
	if err != nil {
		log.Logger.ILog.Warn("issue Printing the data", "function", "CreateLogEntry", "error", err.Error())
	}

	return fmt.Errorf("printer: %w", err)
}

func actionCreateGoal(cliContext *cli.Context) error {
	storagelayer := storage.NewVault()
	goal := resources.NewGoal(
		cliContext.String(flags.NameFlagName),
		cliContext.Timestamp(flags.DueDateFlagName).UTC().Format(time.RFC3339),
		cliContext.String(flags.EntryTextFlagName),
		cliContext.Int(flags.PriorityFlagName))

	row := storagelayer.Storage.CreateGoal(context.TODO(), goal)

	err := printer.NewPrinter(cliContext).Writer.Write(row, os.Stdout)
	if err != nil {
		log.Logger.ILog.Warn("issue Printing the data", "function", "CreateGoal", "error", err.Error())
	}

	return fmt.Errorf("printer: %w", err)
}
