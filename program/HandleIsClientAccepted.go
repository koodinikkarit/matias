package program

import (
	"log"
)

func (p *Program) HandleIsCLientAccepted(
	isClientAccepted bool,
) {
	log.Printf("Is clientaccepted %v", isClientAccepted)
	if isClientAccepted == false {
		p.matiasDatabase.ClearMatiasDatabase()
	}
}
