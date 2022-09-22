package flags

import (
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func GenerateGlobalFlags() []cli.Flag {
	return []cli.Flag{
		altsrc.NewIntFlag(
			&cli.IntFlag{
				Name:    "verbosity",
				Usage:   "Larger numbers entail more logging information.",
				Aliases: []string{"v"},
				Value:   0,
			},
		),
	}
}
