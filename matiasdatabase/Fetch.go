package matiasdatabase

import "github.com/koodinikkarit/matias/models"
import "time"

func (md *MatiasDatabase) FindLastMatiasConnectionDate() *time.Time {
	var lastestMatiasEventConnection models.MatiasEventConnection
	md.DB.Order("connection_end").First(&lastestMatiasEventConnection)
	if lastestMatiasEventConnection.ID == 0 {
		return nil
	}
	return lastestMatiasEventConnection.ConnectionEnd
}

func (md *MatiasDatabase) GetEwDatabases() []models.EwDatabase {
	var ewDatabases []models.EwDatabase
	md.DB.Find(&ewDatabases)
	return ewDatabases
}
