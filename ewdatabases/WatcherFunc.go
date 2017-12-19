package ewdatabases

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func watcherFunc(
	ewDatabaseID uint32,
	lockChannel chan EwDatabaseLockEvent,
	watcher *fsnotify.Watcher,
) {
	for {
		select {
		case event := <-watcher.Events:
			if strings.Contains(event.Name, "Client (1).ulck") {
				log.Printf("Lock event channel %v", lockChannel)
				if event.Op.String() == "CREATE" {
					lockChannel <- EwDatabaseLockEvent{
						EwDatabaseID: ewDatabaseID,
						State:        true,
					}
				} else {
					lockChannel <- EwDatabaseLockEvent{
						EwDatabaseID: ewDatabaseID,
						State:        false,
					}
				}
			}
			// log.Printf("Watcher event name %v op %v", event.Name, event.Op.String())
		case err := <-watcher.Errors:
			log.Printf("Watcher error %v", err)
		}
	}
}
