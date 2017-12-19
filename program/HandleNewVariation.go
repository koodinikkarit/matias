package program

import (
	"log"

	"github.com/koodinikkarit/matias/ewdatabase_models"

	"github.com/koodinikkarit/matias/models"
)

func (p *Program) HandleNewVariation(
	newVariation models.Variation,
) {
	log.Printf("New variation %v", newVariation)
	var sameVariation models.Variation
	p.matiasDatabase.DB.
		Where("server_id = ?", newVariation.ServerID).
		First(&sameVariation)
	if sameVariation.ID == 0 {
		p.matiasDatabase.DB.Create(&newVariation)
	} else {

	}
	var songDatabaseVariation models.SongDatabaseVariation
	p.matiasDatabase.DB.Where("server_id = ?", newVariation.ServerID).
		First(&songDatabaseVariation)

	if songDatabaseVariation.ID == 0 {
		return
	}

	var ewDatabases []models.EwDatabase
	p.matiasDatabase.DB.
		Where("song_database_id = ?", songDatabaseVariation.SongDatabaseID).
		Find(ewDatabases)

	for _, ewDatabase := range ewDatabases {
		ewDatabaseInstance := p.GetOrCreateEwDatabaseInstance(
			ewDatabase,
		)

		var ewDatabaseLink models.EwDatabaseLink
		p.matiasDatabase.DB.Where("ew_database_id = ?", ewDatabase.ID).
			Where("variation_id = ?", newVariation.ServerID).
			First(&ewDatabaseLink)

		if ewDatabaseLink.ID == 0 {
			ewSong, _ := ewDatabaseInstance.CreateSong(
				newVariation.Name,
				newVariation.Text,
			)
			newEwDatabaseLink := models.EwDatabaseLink{
				EwDatabaseID: ewDatabase.ID,
				VariationID:  newVariation.ServerID,
				Version:      newVariation.Version,
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
				newVariation.Name,
				newVariation.Text,
			)
			p.matiasDatabase.DB.Model(&ewDatabaseLink).
				Update("ew_song_id", ewSong.Rowid)
			p.matiasDatabase.DB.Model(&ewDatabaseLink).
				Update("version", 1)
			continue
		}
	}
}
