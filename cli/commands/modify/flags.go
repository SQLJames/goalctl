package modify

import (
	"github.com/sqljames/goalctl/cli/flags"
	"github.com/urfave/cli/v2"
)

var (
	NameGoalFlag *cli.StringFlag = &cli.StringFlag{
		Name:     flags.NameFlagName,
		Usage:    "Name of the Goal.",
		Required: false,
		//Category: "formatting",
		//Aliases: []string{"nb"},
	}
	GoalIDFlag *cli.IntFlag = &cli.IntFlag{
		Name:     flags.GoalIDFlagName,
		Usage:    "ID of the Goal you want to modify",
		Required: true,
		Aliases:  []string{"g"},
	}

	EntryFlag *cli.StringFlag = &cli.StringFlag{
		Name:     flags.EntryTextFlagName,
		Usage:    "Details of your object",
		Required: false,
		//Category: "formatting",
		//Aliases: []string{"nb"},
	}
	NotebookIDFlag *cli.IntFlag = &cli.IntFlag{
		Name:     flags.NotebookIDFlagName,
		Usage:    "ID of the Notebook you want to move the entry too",
		Required: false,
	}
	LogEntryIDFlag *cli.IntFlag = &cli.IntFlag{
		Name:     flags.LogEntryIDFlagName,
		Usage:    "ID of the log entry you want to modify",
		Required: true,
	}
)
