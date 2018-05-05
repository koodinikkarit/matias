package program

import (
	"log"

	"github.com/koodinikkarit/matias/matiasclient"
)

func (p *Program) createMatiasClient() {
	lastestConnectionDate := p.matiasDatabase.FindLastMatiasConnectionDate()
	log.Println("Creating matiasdatabase connection")
	p.matiasClient = matiasclient.NewMatiasClient(
		p.ctx,
		p.matiasServiceIP,
		p.matiasServicePort,
		p.matiasKey,
		lastestConnectionDate,
		p.newMatiasClientGrpcError,
		p.isClientAcceptedChannel,
		p.newEwDatabase,
		p.newSongDatabaseVariation,
		p.removedSongDatabaseVariation,
		p.newSongDatabaseTag,
		p.newTagVariation,
		p.newVariation,
	)
}
