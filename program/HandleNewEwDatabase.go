package program

import (
	"log"

	"github.com/koodinikkarit/matias/ewdatabase_models"
	"github.com/koodinikkarit/matias/models"
)

func (p *Program) HandleNewEwDatabase(
	newEwDatabase models.EwDatabase,
) {
	log.Printf("New ewdatabase %v", newEwDatabase)
	var sameEwDatabase models.EwDatabase
	p.matiasDatabase.DB.
		Where("server_id = ?", newEwDatabase.ServerID).
		First(&sameEwDatabase)

	if sameEwDatabase.ID == 0 {
		p.matiasDatabase.DB.Create(&newEwDatabase)
	}

	ewDatabaseInstance := p.GetOrCreateEwDatabaseInstance(newEwDatabase)

	var variations []models.Variation
	p.matiasDatabase.DB.Table("variations as v").
		Select("*").
		Joins("left join song_database_variations as sdv on v.server_id = sdv.variation_id").
		Find(&variations)

	for _, variation := range variations {
		var ewDatabaseLink models.EwDatabaseLink
		p.matiasDatabase.DB.Where("ew_database_id = ?", newEwDatabase.ServerID).
			Where("variation_id = ?", variation.ServerID).
			First(&ewDatabaseLink)

		if ewDatabaseLink.ID == 0 {
			ewSong, _ := ewDatabaseInstance.CreateSong(
				variation.Name,
				variation.Text,
			)
			newEwDatabaseLink := models.EwDatabaseLink{
				EwDatabaseID: newEwDatabase.ServerID,
				VariationID:  variation.ServerID,
				Version:      variation.Version,
				EwSongID:     ewSong.Rowid,
			}
			p.matiasDatabase.DB.Create(&newEwDatabaseLink)
			continue
		}

		var ewSong ewdatabasemodels.Song
		ewDatabaseInstance.SongsDB.
			Where("rowid = ?", ewDatabaseLink.EwSongID).
			First(&ewSong)

		if ewSong.Rowid == 0 {
			ewSong, _ := ewDatabaseInstance.CreateSong(
				variation.Name,
				variation.Text,
			)
			p.matiasDatabase.DB.Model(&ewDatabaseLink).
				Update("ew_song_id", ewSong.Rowid)
			p.matiasDatabase.DB.Model(&ewDatabaseLink).
				Update("version", 1)
			continue
		}
	}
}
