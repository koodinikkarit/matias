package program

import (
	"github.com/koodinikkarit/matias/ewdatabases"
	"github.com/koodinikkarit/matias/models"
)

func (p *Program) GetOrCreateEwDatabaseInstance(
	ewDatabase models.EwDatabase,
) *ewdatabases.EwDatabaseInstance {
	ewDatabaseInstance := p.ewDatabaseInstances[ewDatabase.ServerID]
	if ewDatabaseInstance == nil {
		ewDatabaseInstance := CreateEwDatabaseInstance(
			p.matiasDatabase,
			ewDatabase.ID,
			ewDatabase.FilesystemPath,
			p.ewDatabaseLockStateChangedChannel,
		)
		p.ewDatabaseInstances[ewDatabase.ServerID] = ewDatabaseInstance
	}
	return ewDatabaseInstance
}
