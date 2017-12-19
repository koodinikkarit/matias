package program

import (
	"log"

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
	p.matiasDatabase.DB.Table("variations v").
		Joins("left join song_database_variations sdv on v.server_id = sdv.variation_id").
		Find(&variations)

	for _, variation := range variations {
		ewSong, _ := ewDatabaseInstance.CreateSong(
			variation.Name,
			variation.Text,
		)
		ewDatabaseLink := models.EwDatabaseLink{
			EwDatabaseID: newEwDatabase.ServerID,
			VariationID:  variation.ServerID,
			EwSongID:     ewSong.Rowid,
			Version:      variation.Version,
		}
		p.matiasDatabase.DB.Create(&ewDatabaseLink)
	}
}
