package matiasdatabase

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/matias/database"
)

type MatiasDatabase struct {
	DB *gorm.DB
}

func NewMatiasDatabase(databasePath string) *MatiasDatabase {
	db, err := database.CreateDBConnection(databasePath)
	if err != nil {
		log.Fatalf("Error while creating dbconnection %v", err)
	}
	return &MatiasDatabase{
		DB: db,
	}
}

func (md *MatiasDatabase) Close() {
	md.DB.Close()
}
