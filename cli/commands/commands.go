package commands

import (
	"github.com/sqljames/goalctl/cli/commands/create"
	"github.com/sqljames/goalctl/cli/commands/export"
	"github.com/sqljames/goalctl/cli/commands/link"
	"github.com/sqljames/goalctl/cli/commands/list"
	"github.com/sqljames/goalctl/cli/commands/modify"
	"github.com/sqljames/goalctl/cli/flags"
	"github.com/sqljames/goalctl/pkg/info"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
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
		Authors:   makeAuthors(),
		Copyright: info.Copyright,
		Suggest:   true,
		Commands: []*cli.Command{
			create.New(),
			list.New(),
			export.New(),
			link.New(),
			modify.New(),
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
		jlogr.Logger.ILog.Error(err, err.Error())
	}

	// The second config map will not overwrite the first
	second := altsrc.InitInputSourceWithContext(cliContext.App.Flags, altsrc.NewYamlSourceFromFlagFunc("global-config"))

	err = second(cliContext)
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())
	}

	jlogr.InitializeLogger(cliContext.Int("verbosity"))

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

func makeAuthors() (authors []*cli.Author) {
	for _, author := range info.Authors {
		authors = append(authors, &cli.Author{
			Name:  author.Name,
			Email: author.Email,
		})
	}
	return authors
}
