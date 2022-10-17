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
		Action: Notebooks,
		Flags: []cli.Flag{
			flags.OutputFormatFlag,
		},
	}
}

func Notebooks(cliContext *cli.Context) error {
	results, err := list.Notebooks()
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return err
	}

	output.Output(cliContext.String(flags.OutputFormatFlagName), results)

	return nil
}
