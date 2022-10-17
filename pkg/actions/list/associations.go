package list

import (
	"context"

	"github.com/sqljames/goalctl/pkg/storage"
	"github.com/sqljames/goalctl/pkg/storage/resources"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func AssociationsByGoalID(goalid int) ([]*resources.Association, error) {
	storagelayer, err := storage.NewVault()
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return nil, err
	}

	return storagelayer.Storage.GetAssociationsByGoalID(context.TODO(), goalid)
}
