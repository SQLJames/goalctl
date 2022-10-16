package list

import (
	"github.com/sqljames/goalctl/cli/flags"
	"github.com/sqljames/goalctl/cli/output"
	"github.com/sqljames/goalctl/pkg/actions/list"
	"github.com/urfave/cli/v2"
)

func listNotebookEntries() *cli.Command {
	return &cli.Command{
		Name:   "entry",
		Usage:  "Prints all the entries in a notebook",
		Action: ListEntries,
		Flags: []cli.Flag{
			flags.NameNotebookFlag,
			flags.OutputFormatFlag,
		},
	}
}

func ListEntries(cliContext *cli.Context) error {
	entries := list.ListEntries(cliContext.String(flags.NameFlagName))
	output.Output(cliContext.String(flags.OutputFormatFlagName), entries)
	return nil
}
