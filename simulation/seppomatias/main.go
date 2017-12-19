package main

import (
	"log"
	"net"

	"github.com/koodinikkarit/matias/matias_service"
	"github.com/koodinikkarit/matias/simulation/seppomatias/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":4445")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	MatiasService.RegisterMatiasServer(
		server,
		&service.MatiasServer{
			EwDatabases: []*MatiasService.EwDatabase{
				&MatiasService.EwDatabase{
					FilesystemPath: "F:\\ewkanta\\ruksakanta",
					SongDatabaseId: 5,
				},
			},
		},
	)
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
