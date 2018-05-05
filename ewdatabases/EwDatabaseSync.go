package ewdatabases

import (
	"log"

	"github.com/koodinikkarit/matias/matiasdatabase"
)

func (ewi *EwDatabaseInstance) SyncEwDatabase(
	matiasDatabase *matiasdatabase.MatiasDatabase,
) {
	//var songs []ewdatabasemodels.Song
	//var words []ewdatabasemodels.Word

	//ewi.SongsDB.Find(&songs)
	//ewi.SongsWordsDB.Find(&words)
	log.Println("fetched")
}
