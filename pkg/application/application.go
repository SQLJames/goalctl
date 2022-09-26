package application

import (
	"os"

	"github.com/sqljames/goalctl/pkg/commands"
	"github.com/sqljames/goalctl/pkg/log"
)

func Run() {
	// Build command tree.
	cmd := commands.NewApp()

	// Execute.
	if err := cmd.Run(os.Args); err != nil {
		log.Logger.ILog.Fatal(err, err.Error())
	}
}
