package flags

import (
	"github.com/urfave/cli/v2"
)

const (
	NameFlagName         string = "name"
	EntryTextFlagName    string = "entry"
	OutputFormatFlagName string = "outputformat"
	TagsFlagName         string = "tag"
	DueDateFlagName      string = "duedate"
	PriorityFlagName     string = "priority"
	GoalIDFlagName       string = "goalid"
	LogEntryIDFlagName   string = "logentryid"
	GoalStatusFlagName   string = "status"
	ConfirmFlagName      string = "confirm"
	NotebookIDFlagName   string = "notebookid"
)

var (
	NameNotebookFlag *cli.StringFlag = &cli.StringFlag{
		Name:     NameFlagName,
		Usage:    "Name of the notebook.",
		Required: true,
		//Category: "formatting",
		//Aliases: []string{"nb"},
	}
	NameGoalFlag *cli.StringFlag = &cli.StringFlag{
		Name:     NameFlagName,
		Usage:    "Name of the Goal.",
		Required: true,
		//Category: "formatting",
		//Aliases: []string{"nb"},
	}

	PriorityFlag *cli.IntFlag = &cli.IntFlag{

		Name:     PriorityFlagName,
		Usage:    "Priority Ranking of the object, 1-127",
		Required: false,
		//Category: "formatting",
		//Aliases: []string{"nb"},
	}
	DueDateFlag *cli.TimestampFlag = &cli.TimestampFlag{
		Name:   DueDateFlagName,
		Usage:  "Due Date for your Goal, YYYY-MM-DD Format",
		Layout: "2006-01-02",
		//Required: true,
		//Category: "formatting",
		//Aliases: []string{"nb"},
	}
	EntryFlag *cli.StringFlag = &cli.StringFlag{
		Name:     EntryTextFlagName,
		Usage:    "Details of your object",
		Required: true,
		//Category: "formatting",
		//Aliases: []string{"nb"},
	}
	OutputFormatFlag *cli.StringFlag = &cli.StringFlag{
		Name:     OutputFormatFlagName,
		Usage:    "output format of the journal",
		Required: false,
		//Category: "formatting",
		Aliases: []string{"out"},
	}
	TagsFlag *cli.StringSliceFlag = &cli.StringSliceFlag{
		Name:    TagsFlagName,
		Aliases: []string{"t"},
	}
	GoalIDFlag *cli.StringSliceFlag = &cli.StringSliceFlag{
		Name:     GoalIDFlagName,
		Usage:    "ID of the Goal you want to link with the log entry",
		Required: true,
		Aliases:  []string{"g"},
	}
	LogEntryIDFlag *cli.StringSliceFlag = &cli.StringSliceFlag{
		Name:     LogEntryIDFlagName,
		Usage:    "ID of the log entry that you want to link with a goal id",
		Required: true,
		Aliases:  []string{"le"},
	}
	GoalStatusFlag *cli.StringSliceFlag = &cli.StringSliceFlag{
		Name:     GoalStatusFlagName,
		Usage:    "The status you would like to assign to the goal",
		Required: false,
		//Aliases:  []string{"le"},
	}
	ConfirmFlag *cli.BoolFlag = &cli.BoolFlag{
		Name:     ConfirmFlagName,
		Usage:    "Flag used to confirm that you want to make the modification",
		Required: false,
		Value:    false,
	}
)
