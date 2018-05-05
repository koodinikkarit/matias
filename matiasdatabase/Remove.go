package matiasdatabase

import (
	"github.com/koodinikkarit/matias/models"
)

func (md *MatiasDatabase) ClearMatiasDatabase() {
	md.DB.Delete(&models.Author{})
	md.DB.Delete(&models.Copyright{})
	md.DB.Delete(&models.EwDatabase{})
	md.DB.Delete(&models.EwDatabaseLink{})
	md.DB.Delete(&models.SongDatabase{})
	md.DB.Delete(&models.SongDatabaseTag{})
	md.DB.Delete(&models.SongDatabaseVariation{})
	md.DB.Delete(&models.TagVariation{})
	md.DB.Delete(&models.Variation{})
}
