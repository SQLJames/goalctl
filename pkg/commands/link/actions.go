package link

import (
	"context"
	"errors"
	"strconv"

	"github.com/sqljames/goalctl/pkg/flags"
	"github.com/sqljames/goalctl/pkg/log"
	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/urfave/cli/v2"
)

func actionLink(cliContext *cli.Context) error {
	logentryIds := removeDuplicate(cliContext.StringSlice(flags.LogEntryIDFlagName))
	goalIds := removeDuplicate(cliContext.StringSlice(flags.GoalIDFlagName))
	links := []resources.Association{}

	for _, logentryId := range logentryIds {
		for _, goalid := range goalIds {
			goalidInt, err := strconv.Atoi(goalid)
			if err != nil {
				return errors.New("converting Goalid to int")
			}
			logentryIdInt, err := strconv.Atoi(logentryId)
			if err != nil {
				return errors.New("onverting LogEntryId to int")
			}
			link := resources.Association{
				GoalID:     goalidInt,
				LogEntryID: logentryIdInt,
			}
			links = append(links, link)
		}
	}
	storagelayer, err := storage.NewStorageLayer()
	if err != nil {
		return err
	}

	for _, entry := range links {

		_, err := storagelayer.CreateAssociation(context.TODO(), entry)
		if err != nil {
			log.Logger.Warn("unable to create association", "error", err.Error())
		}
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
