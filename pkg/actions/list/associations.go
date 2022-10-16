package list

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func AssociationsByGoalID(goalid int) []*resources.Association {
	storagelayer, err := storage.NewVault()
	if err != nil {
		jlogr.Logger.ILog.Fatal(err, err.Error())
	}

	return storagelayer.Storage.GetAssociationsByGoalID(context.TODO(), goalid)
}
