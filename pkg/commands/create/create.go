package create

import (
	"github.com/sqljames/goalctl/pkg/flags"
	"github.com/urfave/cli/v2"
)

func New() *cli.Command {
	return &cli.Command{
		Name:    "create",
		Usage:   "Allows users to create objects",
		Aliases: []string{"c"},
		Subcommands: []*cli.Command{
			createNotebook(),
			createLogEntry(),
			createGoalEntry(),
		},
	}
}

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

func createLogEntry() *cli.Command {
	return &cli.Command{
		Name:   "entry",
		Usage:  "creates a new entry in a specific notebook.",
		Action: actionCreateLogEntry,
		Flags: []cli.Flag{
			flags.NameNotebookFlag,
			flags.EntryFlag,
			flags.TagsFlag,
		},
	}
}

func createGoalEntry() *cli.Command {
	return &cli.Command{
		Name:   "goal",
		Usage:  "creates a new goal on the Goals table",
		Action: actionCreateGoal,
		Flags: []cli.Flag{
			flags.NameGoalFlag,
			flags.EntryFlag,
			flags.PriorityFlag,
			flags.DueDateFlag,
		},
	}
}
