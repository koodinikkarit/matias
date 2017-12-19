package ewdatabases

import (
	"github.com/koodinikkarit/matias/ewdatabase_models"
)

func (ewi *EwDatabaseInstance) CreateSong(
	title string,
	text string,
) (
	ewdatabasemodels.Song,
	ewdatabasemodels.Word,
) {
	song := ewdatabasemodels.Song{
		Title: title,
	}
	ewi.SongsDB.Create(&song)
	songWords := ewdatabasemodels.Word{
		SongId: song.Rowid,
		Words:  PrepareWords(text),
	}
	ewi.SongsWordsDB.Create(&songWords)
	return song, songWords
}
