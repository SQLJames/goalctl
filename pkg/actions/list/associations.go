package list

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
)

func AssociationsByGoalID(goalid int) []*resources.Association {
	storagelayer := storage.NewVault()

	return storagelayer.Storage.GetAssociationsByGoalID(context.TODO(), goalid)
}
