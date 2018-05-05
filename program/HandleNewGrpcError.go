package program

import (
	"log"
	"time"

	"google.golang.org/grpc/codes"
)

func (p *Program) HandleNewGrpcError(
	errorCode codes.Code,
) {
	log.Printf("New matiasclient grpc error %v", uint32(errorCode))
	timer := time.NewTimer(time.Second * 5)
	go func() {
		<-timer.C
		p.reconnectTimer <- true
	}()
}
