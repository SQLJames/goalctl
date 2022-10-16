package create

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func CreateGoal(newGoal NewGoal) *resources.Goal {
	storagelayer, err := storage.NewVault()
	if err != nil {
		jlogr.Logger.ILog.Fatal(err, err.Error())
	}
	goal := resources.NewGoal(newGoal.Goal, newGoal.DueDate, newGoal.Details, newGoal.Priority)

	results := storagelayer.Storage.CreateGoal(context.TODO(), goal)

	return results
}
