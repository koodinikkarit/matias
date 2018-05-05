package ewdatabases

import (
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/matias/ewdatabase_models"
)

func createSongsDB(
	path string,
) (
	*gorm.DB,
	error,
) {
	db, err := gorm.Open("sqlite3", path)
	if err == nil {
		db.Debug().AutoMigrate(
			&ewdatabasemodels.Song{},
		)
	}
	return db.Debug(), err
}

func createSongWordsDB(
	path string,
) (
	*gorm.DB,
	error,
) {
	db, err := gorm.Open("sqlite3", path)
	if err == nil {
		// db.Exec(`CREATE TABLE if not exists word (rowid	integer NOT NULL UNIQUE, song_id integer, words	rtf, slide_uids	text, slide_layout_revisions int64a, slide_revisions int64a, RIMARY KEY(rowid))`)
		db.AutoMigrate(
			&ewdatabasemodels.Word{},
		)
	}
	return db.Debug(), err
}
