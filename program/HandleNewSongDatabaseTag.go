package program

import (
	"log"

	"github.com/koodinikkarit/matias/models"
)

func (p *Program) HandleNewSongDatabaseTag(
	newSongDatabaseTag models.SongDatabaseTag,
) {
	log.Printf("New songdatabasetag %v", newSongDatabaseTag)
	var songDatabaseTag models.SongDatabaseTag
	p.matiasDatabase.DB.
		Where("song_database_id = ?", newSongDatabaseTag.SongDatabaseID).
		Where("tag_id = ?", newSongDatabaseTag.TagID).
		Or("server_id = ?", newSongDatabaseTag.ServerID)

	if songDatabaseTag.ID == 0 {
		songDatabaseTag = newSongDatabaseTag
		p.matiasDatabase.DB.Create(&songDatabaseTag)
	}

	var variations []models.Variation

	p.matiasDatabase.DB.Table("variations v").
		Joins("inner join tag_variations tg on v.server_id = tg.variation_id").
		Where("tg.tag_id = ?", newSongDatabaseTag.TagID).
		Find(&variations)

	var ewDatabases []models.EwDatabase

	p.matiasDatabase.DB.Table("ew_databases ed").
		Where("ed.song_database_id = ?", songDatabaseTag.SongDatabaseID).
		Find(&ewDatabases)

	for _, ewDatabase := range ewDatabases {
		var ewDatabaseLinks []models.EwDatabaseLink
		p.matiasDatabase.DB.Where("ew_database_id = ?", ewDatabase.ServerID).
			Find(&ewDatabaseLinks)

		ewDatabaseInstance := p.ewDatabaseInstances[ewDatabase.ServerID]

		if ewDatabaseInstance == nil {
			continue
		}

		for _, variation := range variations {
			found := false
			for _, ewDatabaseLink := range ewDatabaseLinks {
				if ewDatabaseLink.VariationID != variation.ServerID {
					continue
				}
				found = true
			}
			if found == true {
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
