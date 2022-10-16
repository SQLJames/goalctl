package link

import (
	"context"
	"fmt"
	"strconv"

	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func Link(logEntryID, goalID []string) error {
	logentryIds := removeDuplicate(logEntryID)
	goalIds := removeDuplicate(goalID)
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

			links = append(links, resources.Association{
				GoalID:     goalidInt,
				LogEntryID: logentryIDInt,
			})
		}
	}

	storagelayer, err := storage.NewVault()
	if err != nil {
		jlogr.Logger.ILog.Fatal(err, err.Error())
	}
	for _, entry := range links {
		storagelayer.Storage.CreateAssociation(context.TODO(), entry)
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
