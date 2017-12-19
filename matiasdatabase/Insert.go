package matiasdatabase

import (
	"github.com/koodinikkarit/matias/models"
)

func (md *MatiasDatabase) InsertEwDatabase(ewDatabase *models.EwDatabase) {
	md.DB.Create(&ewDatabase)
}
