package ewdatabases

import (
	"github.com/koodinikkarit/matias/ewdatabase_models"
	"github.com/satori/go.uuid"
)

func (ewi *EwDatabaseInstance) CreateSong(
	title string,
	text string,
) (
	ewdatabasemodels.Song,
	ewdatabasemodels.Word,
) {
	u1 := uuid.Must(uuid.NewV4())
	song := ewdatabasemodels.Song{
		Title:       title,
		SongItemUID: u1.String(),
		SongUID:     u1.String(),
	}
	ewi.SongsDB.Create(&song)

	songWords := ewdatabasemodels.Word{
		SongID: song.Rowid,
		Words:  PrepareWords(text),
	}
	ewi.SongsWordsDB.Create(&songWords)
	return song, ewdatabasemodels.Word{}
}
