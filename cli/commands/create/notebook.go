package create

import (
	"github.com/sqljames/goalctl/cli/flags"
	"github.com/sqljames/goalctl/cli/output"
	"github.com/sqljames/goalctl/pkg/actions/create"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
	"github.com/urfave/cli/v2"
)

func createNotebook() *cli.Command {
	return &cli.Command{
		Name:   "notebook",
		Usage:  "creates a new notebook in your journal.",
		Action: actionCreateNotebook,
		Flags: []cli.Flag{
			flags.NameNotebookFlag,
		},
	}
}

func actionCreateNotebook(cliContext *cli.Context) error {
	result, err := create.Notebook(cliContext.String(flags.NameFlagName))
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return err
	}

	output.Output(cliContext.String(flags.OutputFormatFlagName), result)

	return nil
}
