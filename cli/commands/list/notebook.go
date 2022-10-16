package list

import (
	"github.com/sqljames/goalctl/cli/flags"
	"github.com/sqljames/goalctl/cli/output"
	"github.com/sqljames/goalctl/pkg/actions/list"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
	"github.com/urfave/cli/v2"
)

func listNotebooks() *cli.Command {
	return &cli.Command{
		Name:   "notebook",
		Usage:  "Prints all the notebooks in your journal",
		Action: ListNotebooks,
		Flags: []cli.Flag{
			flags.OutputFormatFlag,
		},
	}
}

func ListNotebooks(cliContext *cli.Context) error {
	notebookList, err := list.GetNotebooks()
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())
		return err
	}
	output.Output(cliContext.String(flags.OutputFormatFlagName), notebookList)

	return nil
}
