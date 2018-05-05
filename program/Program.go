package program

import (
	"context"
	"fmt"
	"log"

	"github.com/koodinikkarit/matias/configuration"
	"github.com/koodinikkarit/matias/ewdatabases"
	"github.com/koodinikkarit/matias/matiasclient"
	"github.com/koodinikkarit/matias/matiasdatabase"
	"github.com/koodinikkarit/matias/models"
	"google.golang.org/grpc/codes"
)

type Program struct {
	ctx                               context.Context
	ewDatabaseInstances               map[uint32]*ewdatabases.EwDatabaseInstance
	clientAccepted                    bool
	clientConnected                   bool
	matiasDatabase                    *matiasdatabase.MatiasDatabase
	matiasClient                      *matiasclient.MatiasClient
	configFile                        configuration.ConfigFile
	newMatiasClientGrpcError          chan codes.Code
	isClientAcceptedChannel           chan bool
	newEwDatabase                     chan models.EwDatabase
	newSongDatabaseVariation          chan models.SongDatabaseVariation
	removedSongDatabaseVariation      chan models.SongDatabaseVariation
	newSongDatabaseTag                chan models.SongDatabaseTag
	newTagVariation                   chan models.TagVariation
	newVariation                      chan models.Variation
	ewDatabaseLockStateChangedChannel chan ewdatabases.EwDatabaseLockEvent
	reconnectTimer                    chan bool
	matiasServiceIP                   string
	matiasServicePort                 string
	matiasKey                         string
}

func (p *Program) Start() {
	p.parseConfig()
	for {
		select {
		case errorCode := <-p.newMatiasClientGrpcError:
			p.HandleNewGrpcError(errorCode)
		case <-p.reconnectTimer:
			fmt.Println("reconnectTiemr")
			p.createMatiasClient()
		case isClientAccepted := <-p.isClientAcceptedChannel:
			p.HandleIsCLientAccepted(isClientAccepted)
		case newEwDatabase := <-p.newEwDatabase:
			p.HandleNewEwDatabase(
				newEwDatabase,
			)
		case newSongDatabaseVariation := <-p.newSongDatabaseVariation:
			p.HandleNewSongDatabaseVariation(
				newSongDatabaseVariation,
			)
		case removedSongDatabaseVariation := <-p.removedSongDatabaseVariation:
			p.HandleRemovedSongDatabaseVariation(
				removedSongDatabaseVariation,
			)
		case newSongDatabaseTag := <-p.newSongDatabaseTag:
			p.HandleNewSongDatabaseTag(
				newSongDatabaseTag,
			)
		case newTagVariation := <-p.newTagVariation:
			p.HandleNewTagVariation(
				newTagVariation,
			)
		case newVariation := <-p.newVariation:
			fmt.Println("newVariation", newVariation)
			p.HandleNewVariation(
				newVariation,
			)
		case lockEvent := <-p.ewDatabaseLockStateChangedChannel:
			log.Printf(
				"EwDatabaseLock ewDatabaseID: %v state: %v",
				lockEvent.EwDatabaseID,
				lockEvent.State,
			)
		}
	}
}

func (p *Program) Stop() {
	if p.matiasClient != nil {
		p.matiasClient.Close()
	}
}

func NewProgram(ctx context.Context) *Program {
	return &Program{
		ctx: ctx,
		newMatiasClientGrpcError:          make(chan codes.Code),
		isClientAcceptedChannel:           make(chan bool),
		newEwDatabase:                     make(chan models.EwDatabase),
		newSongDatabaseVariation:          make(chan models.SongDatabaseVariation),
		removedSongDatabaseVariation:      make(chan models.SongDatabaseVariation),
		newSongDatabaseTag:                make(chan models.SongDatabaseTag),
		newTagVariation:                   make(chan models.TagVariation),
		newVariation:                      make(chan models.Variation),
		ewDatabaseLockStateChangedChannel: make(chan ewdatabases.EwDatabaseLockEvent),
		ewDatabaseInstances:               make(map[uint32]*ewdatabases.EwDatabaseInstance),
		reconnectTimer:                    make(chan bool),
	}
}
