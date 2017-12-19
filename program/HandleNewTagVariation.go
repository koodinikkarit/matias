package program

import (
	"log"

	"github.com/koodinikkarit/matias/models"
)

func (p *Program) HandleNewTagVariation(
	newTagVariation models.TagVariation,
) {
	log.Printf("New tagvariation %v", newTagVariation)

	var tagVariation models.TagVariation
	p.matiasDatabase.DB.Where("tag_variation.tag_id = ?", newTagVariation.TagID).
		Where("tag_variation.variation_id = ?", newTagVariation.VariationID).
		First(&tagVariation)

	if tagVariation.ID == 0 {
		tagVariation = newTagVariation
		p.matiasDatabase.DB.Create(&tagVariation)
	}

	var variation models.Variation
	p.matiasDatabase.DB.Where("server_id = ?", newTagVariation.VariationID).
		First(&variation)

	if variation.ID == 0 {
		return
	}

	var ewDatabases []models.EwDatabase

	p.matiasDatabase.DB.Table("ew_databases ed").
		Joins("inner join song_databases sd on ed.song_database_id = sd.server_id").
		Joins("inner join song_database_tags sdt on sd.server_id = sdt.song_database_id").
		Where("sdt.tag_id = ?", newTagVariation.TagID)

	for _, ewDatabase := range ewDatabases {
		var sameVariationCount uint32
		p.matiasDatabase.DB.Table("ew_database_links").
			Where("ew_database_id = ?", ewDatabase.ServerID).
			Where("variation_id = ?", variation.ServerID).
			Count(&sameVariationCount)

		if sameVariationCount == 0 {
			ewDatabaseInstance := p.GetOrCreateEwDatabaseInstance(
				ewDatabase,
			)

			if ewDatabaseInstance == nil {
				continue
			}

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
		}
	}
}
