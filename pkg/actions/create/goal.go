package create

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
)

func CreateGoal(newGoal NewGoal) *resources.Goal {
	storagelayer := storage.NewVault()
	goal := resources.NewGoal(newGoal.Goal, newGoal.DueDate, newGoal.Details, newGoal.Priority)

	results := storagelayer.Storage.CreateGoal(context.TODO(), goal)

	return results
}
