package program

import (
	"log"

	"google.golang.org/grpc/codes"
)

func (p *Program) HandleNewGrpcError(
	errorCode codes.Code,
) {
	log.Printf("New matiasclient grpc error %v", uint32(errorCode))
}
