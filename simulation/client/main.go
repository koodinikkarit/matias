package main

import (
	"io"
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/koodinikkarit/matias/matias_service"
)

func main() {
	ctx := context.Background()
	go clitter(ctx, "kon1", 5000)
	go clitter(ctx, "kon2", 0)
	for {
	}
}

func clitter(ctx context.Context, hostName string, timeout uint32) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:4445", opts...)
	if err != nil {
		log.Fatalf("Error while creating matiasclient connection %v", err)
	}

	client := MatiasService.NewMatiasClient(conn)
	newCtx, cancel := context.WithCancel(ctx)

	stream, err := client.ListenChanges(
		newCtx,
		&MatiasService.ListenEventsRequest{
			Key:      "qwertyu12345",
			HostName: hostName,
		},
	)
	if err != nil {
		//log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
	}
	log.Printf("Timeout is %v", timeout)
	if timeout > 0 {
		timer := time.NewTimer(time.Second * 10)
		go func() {
			<-timer.C
			cancel()
		}()
	}
	for {
		eventItem, err := stream.Recv()
		if err == io.EOF {
			log.Println("EOF")
			break
		}
		if err != nil {
			log.Printf("%v.ListFeatures(_) = _, %v", client, err)
			break
		}
		log.Println(hostName, eventItem)
	}
}
