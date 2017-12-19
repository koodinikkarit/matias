package ewdatabases

import (
	"log"
	"path"

	"github.com/jinzhu/gorm"

	"github.com/fsnotify/fsnotify"
	"github.com/koodinikkarit/matias/utility"
)

type EwDatabaseInstance struct {
	filesystemPath        string
	ewDatabaseLockChannel chan EwDatabaseLockEvent
	watcher               *fsnotify.Watcher
	isLocked              bool
	SongsDB               *gorm.DB
	SongsWordsDB          *gorm.DB
}

func NewEwDatabaseInstance(
	ewDatabaseID uint32,
	filesystemPath string,
	ewDatabaseLockChnnel chan EwDatabaseLockEvent,
) *EwDatabaseInstance {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Error while creating lock watcher error: %v", err)
	}

	go watcherFunc(
		ewDatabaseID,
		ewDatabaseLockChnnel,
		watcher,
	)

	locksPath := path.Join(filesystemPath, "/v6.1/Databases/Locks/")
	err = watcher.Add(locksPath)
	if err != nil {
		log.Fatalf(
			"Error while adding path %v to watcher error: %v",
			locksPath,
			err,
		)
	}

	databaseFound := utility.FileExists(
		filesystemPath,
	)

	if databaseFound {
		log.Println("Ewdatabase not found creating one")
	}

	log.Printf("songs path %v", path.Join(filesystemPath, "/v6.1/Databases/Data/Songs.db"))

	songsDB, songsErr := createSongsDB(
		path.Join(filesystemPath, "/v6.1/Databases/Data/Songs.db"),
	)

	if songsErr != nil {
		log.Printf("Error while creating ewdatabaseinstance songs connection error %v", songsErr)
		return nil
	}

	songWordsDB, songWordsErr := createSongWordsDB(
		path.Join(filesystemPath, "/v6.1/Databases/Data/SongWords.db"),
	)

	if songWordsErr != nil {
		log.Printf("Error while creating ewdatabaseinstance songwords connection error %v", songWordsErr)
		return nil
	}

	return &EwDatabaseInstance{
		SongsDB:      songsDB,
		SongsWordsDB: songWordsDB,
	}
}

func (ewi *EwDatabaseInstance) Close() {
	if ewi.watcher != nil {
		ewi.watcher.Close()
	}
}

func (ewi *EwDatabaseInstance) IsLocked() bool {
	return utility.FileExists(ewi.filesystemPath)
}
