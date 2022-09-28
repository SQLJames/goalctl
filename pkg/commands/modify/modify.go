package modify

import (
	"github.com/sqljames/goalctl/pkg/flags"
	"github.com/urfave/cli/v2"
)

func New() *cli.Command {
	return &cli.Command{
		Name:    "modify",
		Usage:   "Allows users to modify objects",
		Aliases: []string{"m"},
		Subcommands: []*cli.Command{
			modifyGoal(),
		},
	}
}

var (
	NameGoalFlag *cli.StringFlag = &cli.StringFlag{
		Name:     flags.NameFlagName,
		Usage:    "Name of the Goal.",
		Required: false,
		//Category: "formatting",
		//Aliases: []string{"nb"},
	}
	GoalIDFlag *cli.StringFlag = &cli.StringFlag{
		Name:     flags.GoalIDFlagName,
		Usage:    "ID of the Goal you want to link with the log entry",
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
)

func modifyGoal() *cli.Command {
	return &cli.Command{
		Name:   "goal",
		Usage:  "Allows users to modify a target goal",
		Action: actionModifyGoal,
		Flags: []cli.Flag{
			GoalIDFlag,
			flags.PriorityFlag,
			flags.GoalStatusFlag,
			flags.DueDateFlag,
			EntryFlag,
			NameGoalFlag,
		},
	} 
}
