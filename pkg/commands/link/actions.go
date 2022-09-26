package link

import (
	"context"
	"fmt"
	"strconv"

	"github.com/sqljames/goalctl/pkg/flags"
	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/urfave/cli/v2"
)

func actionLink(cliContext *cli.Context) error {
	logentryIds := removeDuplicate(cliContext.StringSlice(flags.LogEntryIDFlagName))
	goalIds := removeDuplicate(cliContext.StringSlice(flags.GoalIDFlagName))
	links := []resources.Association{}

	for _, logentryID := range logentryIds {
		for _, goalid := range goalIds {
			goalidInt, err := strconv.Atoi(goalid)
			if err != nil {
				return fmt.Errorf("converting GoalID: %w", err)
			}
			logentryIDInt, err := strconv.Atoi(logentryID)
			if err != nil {
				return fmt.Errorf("converting logentryID: %w", err)
			}
			link := resources.Association{
				GoalID:     goalidInt,
				LogEntryID: logentryIDInt,
			}
			links = append(links, link)
		}
	}

	storagelayer := storage.NewVault()
	for _, entry := range links {
		storagelayer.CreateAssociation(context.TODO(), entry)
	}

	return nil
}

func removeDuplicate[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true

			list = append(list, item)
		}
	}
	return list
}
