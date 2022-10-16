package create

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func CreateGoal(newGoal NewGoal) (*resources.Goal, error) {
	storagelayer, err := storage.NewVault()
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())
		return nil, err
	}
	goal := resources.NewGoal(newGoal.Goal, newGoal.DueDate, newGoal.Details, newGoal.Priority)

	results, err := storagelayer.Storage.CreateGoal(context.TODO(), goal)
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())
		return nil, err
	}
	return results, nil
}
