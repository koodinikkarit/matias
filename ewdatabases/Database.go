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
		db.AutoMigrate(
			&ewdatabasemodels.Word{},
		)
	}
	return db.Debug(), err
}
