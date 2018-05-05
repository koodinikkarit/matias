package program

import (
	"log"

	"github.com/koodinikkarit/matias/models"
)

func (p *Program) HandleNewSongDatabaseVariation(
	newSongDatabaseVariation models.SongDatabaseVariation,
) {
	log.Printf("New songdatabasevariation %v", newSongDatabaseVariation)
	var songDatabaseVariation models.SongDatabaseVariation
	p.matiasDatabase.DB.
		Where("song_database_id = ?", newSongDatabaseVariation.SongDatabaseID).
		Where("variation_id = ?", newSongDatabaseVariation.VariationID).
		First(&songDatabaseVariation)

	if songDatabaseVariation.ID == 0 {
		songDatabaseVariation = newSongDatabaseVariation
		p.matiasDatabase.DB.Create(&songDatabaseVariation)
	}

	var variation models.Variation
	p.matiasDatabase.DB.Where("server_id = ?", newSongDatabaseVariation.VariationID).
		First(&variation)

	if variation.ID == 0 {
		return
	}

	var ewDatabases []models.EwDatabase
	p.matiasDatabase.DB.
		Where("song_database_id = ?", newSongDatabaseVariation.SongDatabaseID).
		Find(&ewDatabases)

	for _, ewDatabase := range ewDatabases {
		ewDatabaseInstance := p.GetOrCreateEwDatabaseInstance(
			ewDatabase,
		)

		if ewDatabaseInstance == nil {
			continue
		}

		var ewDatabaseLink models.EwDatabaseLink
		p.matiasDatabase.DB.
			Where("ew_database_id = ?", ewDatabase.ServerID).
			Where("variation_id = ?", newSongDatabaseVariation.VariationID).
			Find(&ewDatabaseLink)

		if ewDatabaseLink.ID == 0 {
			ewSong, _ := ewDatabaseInstance.CreateSong(
				variation.Name,
				variation.Text,
			)
			ewDatabaseLink := models.EwDatabaseLink{
				EwDatabaseID: ewDatabase.ServerID,
				VariationID:  variation.ServerID,
				EwSongID:     ewSong.Rowid,
				Version:      variation.Version,
			}
			p.matiasDatabase.DB.Create(&ewDatabaseLink)
			continue
		}
	}
}
