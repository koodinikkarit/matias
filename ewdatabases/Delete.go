package ewdatabases

import (
	"github.com/koodinikkarit/matias/ewdatabase_models"
)

func (ewi *EwDatabaseInstance) DeleteEwSong(ewSongID uint32) {
	ewi.SongsDB.Where("row_id = ?", ewSongID).Delete(&ewdatabasemodels.Song{})
	ewi.SongsWordsDB.Where("song_id = ?", ewSongID).Delete(&ewdatabasemodels.Word{})
}
