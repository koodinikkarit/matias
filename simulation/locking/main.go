package main

import (
	"log"
	"sync"

	"github.com/koodinikkarit/matias/ewdatabases"
)

func main() {
	var wg sync.WaitGroup
	lockChannel := make(chan ewdatabases.EwDatabaseLockEvent)
	wg.Add(1)
	ewdatabases.NewEwDatabaseInstance(
		5,
		"F:\\ewkanta\\ruksakanta",
		lockChannel,
	)
	wg.Add(1)
	go func() {
		for {
			event := <-lockChannel
			log.Printf(
				"New lock event ewDatabaseID: %v state: %v",
				event.EwDatabaseID,
				event.State,
			)
		}
	}()
	wg.Wait()
}
