package commands

import (
	"github.com/sqljames/goalctl/pkg/commands/create"
	"github.com/sqljames/goalctl/pkg/commands/export"
	"github.com/sqljames/goalctl/pkg/commands/link"
	"github.com/sqljames/goalctl/pkg/commands/list"
	"github.com/sqljames/goalctl/pkg/flags"
	"github.com/sqljames/goalctl/pkg/info"
	"github.com/sqljames/goalctl/pkg/log"
	"github.com/sqljames/goalctl/pkg/version"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func NewApp() *cli.App {
	modifyCLIDefaultVersion()

	app := &cli.App{
		Name:      info.GetApplicationName(),
		Usage:     info.Description,
		Flags:     flags.GenerateGlobalFlags(),
		Version:   version.Version.String(),
		Before:    beforeTasks,
		Authors:   info.Authors,
		Copyright: info.Copyright,
		Suggest:   true,
		Commands: []*cli.Command{
			create.New(),
			list.New(),
			export.New(),
			link.New(),
		},
	}

	return app
}

func beforeTasks(cliContext *cli.Context) error {
	taskList := []func(*cli.Context) error{
		instrumentLoggingFlags,
	}

	for _, t := range taskList {
		if err := t(cliContext); err != nil {
			return err
		}
	}

	return nil
}

// Care of: https://github.com/physcat/klog-cli/blob/main/main.go
func instrumentLoggingFlags(cliContext *cli.Context) error {
	// Command line flags always overwrite configuration files
	first := altsrc.InitInputSourceWithContext(cliContext.App.Flags, altsrc.NewYamlSourceFromFlagFunc("config"))

	err := first(cliContext)
	if err != nil {
		log.Logger.ILog.Error(err, err.Error())
	}

	// The second config map will not overwrite the first
	second := altsrc.InitInputSourceWithContext(cliContext.App.Flags, altsrc.NewYamlSourceFromFlagFunc("global-config"))
	
	err = second(cliContext)
	if err != nil {
		log.Logger.ILog.Error(err, err.Error())
	}

	log.InitializeLogger(cliContext.Int("verbosity"))

	return err
}

func modifyCLIDefaultVersion() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:  "version",
		Usage: "Prints the version",
		Value: false,
	}
	cli.VersionPrinter = func(cCtx *cli.Context) {
		version.Version.CliPrinter()
	}
}
