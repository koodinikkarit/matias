package database

import "github.com/jinzhu/gorm"
import "log"

func CreateDBConnection(
	databasePath string,
) (
	*gorm.DB,
	error,
) {
	db, err := gorm.Open("sqlite3", databasePath)
	if err == nil {
		log.Println("Migrating database")
		Migrate(db.Debug())
	}
	return db.Debug(), err
}
