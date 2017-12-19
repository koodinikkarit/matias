package program

import (
	"github.com/koodinikkarit/matias/ewdatabases"
	"github.com/koodinikkarit/matias/matiasdatabase"
)

func CreateEwDatabaseInstance(
	matiasDatabase *matiasdatabase.MatiasDatabase,
	ewDatabaseID uint32,
	filesytemPath string,
	ewDatabaseLockStateChangedChannel chan ewdatabases.EwDatabaseLockEvent,
) *ewdatabases.EwDatabaseInstance {
	newInstance := ewdatabases.NewEwDatabaseInstance(
		ewDatabaseID,
		filesytemPath,
		ewDatabaseLockStateChangedChannel,
	)

	if newInstance.IsLocked() == false {
		newInstance.SyncEwDatabase(
			matiasDatabase,
		)
	}

	return newInstance
}
