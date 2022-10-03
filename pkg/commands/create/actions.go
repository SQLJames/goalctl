package create

import (
	"context"
	"fmt"
	"os"

	"github.com/sqljames/goalctl/pkg/util"
	"github.com/sqljames/goalctl/pkg/util/jlogr"

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
		jlogr.Logger.ILog.Error(&emptryStringError{}, "Error creating notebook", "function", "actionCreateNotebook")

		return &emptryStringError{}
	}

	storagelayer := storage.NewVault()
	row := storagelayer.Storage.CreateNotebook(context.TODO(), cliContext.String(flags.NameFlagName))

	err := printer.NewPrinter(cliContext).Writer.Write(row, os.Stdout)
	if err != nil {
		jlogr.Logger.ILog.Warn("issue Printing the data", "function", "CreateNotebook", "error", err.Error())
		err = fmt.Errorf("printer: %w", err)
	}

	return err
}

func actionCreateLogEntry(cliContext *cli.Context) error {
	storagelayer := storage.NewVault()
	le := resources.NewLogEntry(cliContext.String(flags.EntryTextFlagName), cliContext.StringSlice(flags.TagsFlagName))

	row := storagelayer.Storage.CreateLogEntry(context.TODO(), le, cliContext.String(flags.NameFlagName))

	err := printer.NewPrinter(cliContext).Writer.Write(row, os.Stdout)
	if err != nil {
		jlogr.Logger.ILog.Warn("issue Printing the data", "function", "CreateLogEntry", "error", err.Error())
		err = fmt.Errorf("printer: %w", err)
	}

	return err
}

func actionCreateGoal(cliContext *cli.Context) error {
	storagelayer := storage.NewVault()
	goal := resources.NewGoal(
		cliContext.String(flags.NameFlagName),
		util.TimeToString(cliContext.Timestamp(flags.DueDateFlagName)),
		cliContext.String(flags.EntryTextFlagName),
		cliContext.Int(flags.PriorityFlagName))

	row := storagelayer.Storage.CreateGoal(context.TODO(), goal)

	err := printer.NewPrinter(cliContext).Writer.Write(row, os.Stdout)
	if err != nil {
		jlogr.Logger.ILog.Warn("issue Printing the data", "function", "CreateGoal", "error", err.Error())
		err = fmt.Errorf("printer: %w", err)
	}

	return err
}
