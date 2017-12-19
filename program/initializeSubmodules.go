package program

import (
	"log"

	"github.com/koodinikkarit/matias/matiasclient"
)

func (p *Program) createMatiasClient(
	matiasServiceIP string,
	matiasServicePort string,
	matiasKey string,
) {
	lastestConnectionDate := p.matiasDatabase.FindLastMatiasConnectionDate()
	log.Println("Creating matiasdatabase connection")
	p.matiasClient = matiasclient.NewMatiasClient(
		p.ctx,
		matiasServiceIP,
		matiasServicePort,
		matiasKey,
		lastestConnectionDate,
		p.newMatiasClientGrpcError,
		p.isClientAcceptedChannel,
		p.newEwDatabase,
		p.newSongDatabaseVariation,
		p.newSongDatabaseTag,
		p.newTagVariation,
		p.newVariation,
	)
}
