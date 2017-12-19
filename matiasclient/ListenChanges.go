package matiasclient

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	"github.com/koodinikkarit/matias/models"

	"github.com/koodinikkarit/matias/matias_service"
)

func (mc *MatiasClient) ListenChanges(
	lastestConnection *time.Time,
) {
	go func() {
		HostName, _ := os.Hostname()
		stream, err := mc.client.ListenChanges(
			context.Background(),
			&MatiasService.ListenEventsRequest{
				Key:      mc.matiasKey,
				HostName: HostName,
			},
		)
		if err != nil {
			mc.newGrpcError <- grpc.Code(err)
			log.Printf("error in listeting changes %v", err)
			return
		}
		for {
			eventItem, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				mc.newGrpcError <- grpc.Code(err)
				log.Println("Listen recv error %v", err)
				return
			}
			switch event := eventItem.EventMessage.(type) {
			case *MatiasService.EventItem_AcceptedKey:
				mc.clientAccepted <- event.AcceptedKey
			case *MatiasService.EventItem_NewEwDatabase:
				log.Printf("newEwDatabas %v", event.NewEwDatabase)
				mc.newEwDatabase <- models.EwDatabase{
					ServerID:       event.NewEwDatabase.Id,
					FilesystemPath: event.NewEwDatabase.FilesystemPath,
					SongDatabaseID: event.NewEwDatabase.SongDatabaseId,
				}
			case *MatiasService.EventItem_NewSongDatabaseVariation:
				mc.newSongDatabaseVariation <- models.SongDatabaseVariation{
					ServerID:       event.NewSongDatabaseVariation.Id,
					SongDatabaseID: event.NewSongDatabaseVariation.SongDatabaseId,
					VariationID:    event.NewSongDatabaseVariation.VariationId,
				}
			case *MatiasService.EventItem_NewSongDatabaseTag:
				mc.newSongDatabaseTag <- models.SongDatabaseTag{
					ServerID:       event.NewSongDatabaseTag.Id,
					SongDatabaseID: event.NewSongDatabaseTag.SongDatabaseId,
					TagID:          event.NewSongDatabaseTag.TagId,
				}
			case *MatiasService.EventItem_NewTagVariation:
				mc.newTagVariation <- models.TagVariation{
					ServerID:    event.NewTagVariation.Id,
					TagID:       event.NewTagVariation.TagId,
					VariationID: event.NewTagVariation.VariationId,
				}
			case *MatiasService.EventItem_NewVariation:
				//log.Printf("New variation %v", event.NewVariation)
				mc.newVariation <- models.Variation{
					ServerID: event.NewVariation.Id,
					Name:     event.NewVariation.Name,
					Text:     event.NewVariation.Text,
				}
			}
		}
	}()
}
