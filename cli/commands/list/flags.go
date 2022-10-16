package list

import (
	"github.com/urfave/cli/v2"

	"github.com/sqljames/goalctl/cli/flags"
)

var (
	PastDueFlag *cli.BoolFlag = &cli.BoolFlag{
		Name:     flags.PastDueFlagName,
		Usage:    "Flag used to filter expired Goals",
		Required: false,
		Value:    false,
	}
)
