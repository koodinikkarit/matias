package matiasclient

import "github.com/koodinikkarit/matias/models"

type Channels struct {
	clientAccepted           chan bool
	newEwDatabase            chan models.EwDatabase
	newSongDatabaseVariation chan models.SongDatabaseVariation
	newSongDatabaseTag       chan models.SongDatabaseTag
	newTagVariation          chan models.TagVariation
}
