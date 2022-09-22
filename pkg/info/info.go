package info

import (
	"github.com/urfave/cli/v2"
)

const (
	Description     string = "A Cli tool for managing your Daily Journal entries."
	Copyright string = "Database Ally, LLC, 2022"
)

var (
	applicationName = "unknown"
	Authors         = []*cli.Author{
		{
			Name:  "James Rhoat",
			Email: "James@Rhoat.com",
		},
	}
)

func GetApplicationName() string {
	return applicationName
}
