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

func actionCreateNotebook(c *cli.Context) error {
	storagelayer, err := storage.NewStorageLayer()
	if err != nil {
		return err
	}

	row, err := storagelayer.CreateNotebook(context.TODO(), c.String(flags.NameFlagName))
	if err != nil {
		return err
	}
	printer := printer.NewPrinter(c)
	err = printer.Write(row, os.Stdout)
	if err != nil {
		log.Logger.Warn("issue Printing the data", "function", "CreateNotebook", "error", err.Error())
	}
	return err
}

func actionCreateLogEntry(c *cli.Context) error {
	storagelayer, err := storage.NewStorageLayer()
	if err != nil {
		return err
	}
	le := resources.NewLogEntry(c.String(flags.EntryTextFlagName), c.StringSlice(flags.TagsFlagName))

	row, err := storagelayer.CreateLogEntry(context.TODO(), *le, c.String(flags.NameFlagName))
	if err != nil {
		return err
	}
	printer := printer.NewPrinter(c)
	err = printer.Write(row, os.Stdout)
	if err != nil {
		log.Logger.Warn("issue Printing the data", "function", "CreateLogEntry", "error", err.Error())
	}
	return err
}

func actionCreateGoal(c *cli.Context) error {
	storagelayer, err := storage.NewStorageLayer()
	if err != nil {
		return err
	}

	goal := resources.NewGoal(c.String(flags.NameFlagName), c.Timestamp(flags.DueDateFlagName).UTC().Format(time.RFC3339), c.String(flags.EntryTextFlagName), c.Int(flags.PriorityFlagName))

	row, err := storagelayer.CreateGoal(context.TODO(), *goal)
	if err != nil {
		return err
	}
	printer := printer.NewPrinter(c)
	err = printer.Write(row, os.Stdout)
	if err != nil {
		log.Logger.Warn("issue Printing the data", "function", "CreateGoal", "error", err.Error())
	}
	return err
}
