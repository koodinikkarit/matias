package program

import "github.com/koodinikkarit/matias/models"

func (p *Program) HandleRemovedSongDatabaseVariation(
	removedSongDatabaseVariation models.SongDatabaseVariation,
) {
	p.matiasDatabase.DB.Where(
		"song_database_id = ? and variation_id = ?",
		removedSongDatabaseVariation.SongDatabaseID,
		removedSongDatabaseVariation.VariationID).
		Delete(&models.SongDatabaseVariation{})

	var ewDatabases []models.EwDatabase
	p.matiasDatabase.DB.
		Where("song_database_id = ?", removedSongDatabaseVariation.SongDatabaseID).
		Find(&ewDatabases)

	var ewDatabaseIds []uint32

	for _, ewDatabase := range ewDatabases {
		ewDatabaseIds = append(
			ewDatabaseIds,
			ewDatabase.ServerID,
		)
	}

	var ewDatabaseLinks []models.EwDatabaseLink
	p.matiasDatabase.DB.
		Where("ew_database_id in (?)", ewDatabaseIds).
		Where("variation_id = ?", removedSongDatabaseVariation.VariationID).
		Find(&ewDatabaseLinks)

	p.matiasDatabase.DB.Where("ew_database_id in (?)", ewDatabaseIds).
		Where("variation_id = ?", removedSongDatabaseVariation.VariationID).
		Delete(&models.EwDatabaseLink{})

	for _, ewDatabaseID := range ewDatabaseIds {
		ewDatabaseInstance := p.ewDatabaseInstances[ewDatabaseID]
		for _, ewDatabaseLink := range ewDatabaseLinks {
			ewDatabaseInstance.DeleteEwSong(ewDatabaseLink.EwSongID)
		}
	}

	var variationsCount uint32
	p.matiasDatabase.DB.Table("song_database_variations").
		Where("variation_id = ?", removedSongDatabaseVariation.VariationID).
		Count(&variationsCount)

	if variationsCount == 0 {
		p.matiasDatabase.DB.Where("server_id = ?", removedSongDatabaseVariation.VariationID).
			Delete(&models.SongDatabaseVariation{})
	}
}
