package application

import (
	"os"

	"github.com/sqljames/goalctl/cli/commands"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func Run() {
	// Build command tree.
	cmd := commands.NewApp()

	// Execute.
	if err := cmd.Run(os.Args); err != nil {
		jlogr.Logger.ILog.Fatal(err, err.Error())
	}
}
