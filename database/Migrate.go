package database

import "github.com/jinzhu/gorm"
import "github.com/koodinikkarit/matias/models"

func Migrate(
	db *gorm.DB,
) {
	db.AutoMigrate(
		&models.Author{},
		&models.Copyright{},
		&models.EwDatabase{},
		&models.EwDatabaseLink{},
		&models.SongDatabase{},
		&models.SongDatabaseTag{},
		&models.SongDatabaseVariation{},
		&models.Variation{},
		&models.MatiasEventConnection{},
	)
}
