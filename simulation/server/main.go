package main

import (
	"log"
	"net"
	"time"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/koodinikkarit/matias/matias_service"
)

type ServerType struct {
}

func main() {
	lis, err := net.Listen("tcp", ":4445")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	MatiasService.RegisterMatiasServer(
		server,
		&ServerType{},
	)
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (st *ServerType) ListenChanges(
	req *MatiasService.ListenEventsRequest,
	stream MatiasService.Matias_ListenChangesServer,
) error {
	log.Printf("Uusi kysely key %v hostname %v", req.Key, req.HostName)
	ev := &MatiasService.EventItem{
		EventMessage: &MatiasService.EventItem_AcceptedKey{
			AcceptedKey: true,
		},
	}
	stream.Send(ev)

	for {
		<-time.NewTimer(time.Second * 1).C
		ev = &MatiasService.EventItem{
			EventMessage: &MatiasService.EventItem_NewTagVariation{
				NewTagVariation: &MatiasService.TagVariation{
					TagId:       120,
					VariationId: 752,
				},
			},
		}
		log.Printf("Sending %v to client %v", ev, req.HostName)
		err := stream.Send(ev)
		if err != nil {
			log.Printf("Error stopping %v", err)
			return err
		}
	}
	return nil
}
