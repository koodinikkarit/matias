package matiasclient

import (
	"context"
	"log"
	"time"

	"github.com/koodinikkarit/matias/ewdatabases"
	"github.com/koodinikkarit/matias/matias_service"
	"github.com/koodinikkarit/matias/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type MatiasClient struct {
	ctx                      context.Context
	cancel                   context.CancelFunc
	ip                       string
	port                     string
	matiasKey                string
	ewDatabaseInstancces     []ewdatabases.EwDatabaseInstance
	client                   MatiasService.MatiasClient
	newGrpcError             chan codes.Code
	clientAccepted           chan bool
	newEwDatabase            chan models.EwDatabase
	newSongDatabaseVariation chan models.SongDatabaseVariation
	newSongDatabaseTag       chan models.SongDatabaseTag
	newTagVariation          chan models.TagVariation
	newVariation             chan models.Variation
}

func NewMatiasClient(
	ctx context.Context,
	ip string,
	port string,
	matiasKey string,
	latestConnectionDate *time.Time,
	newGrpcError chan codes.Code,
	clientAccepted chan bool,
	newEwDatabase chan models.EwDatabase,
	newSongDatabaseVariation chan models.SongDatabaseVariation,
	newSongDatabaseTag chan models.SongDatabaseTag,
	newTagVariation chan models.TagVariation,
	newVariation chan models.Variation,
) *MatiasClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(ip+":"+port, opts...)
	if err != nil {
		log.Fatalf("Error while creating matiasclient connection %v", err)
	}
	client := MatiasService.NewMatiasClient(conn)

	newCtx, cancel := context.WithCancel(ctx)
	matiasClient := &MatiasClient{
		ctx:                      newCtx,
		cancel:                   cancel,
		ip:                       ip,
		port:                     port,
		matiasKey:                matiasKey,
		client:                   client,
		newGrpcError:             newGrpcError,
		clientAccepted:           clientAccepted,
		newEwDatabase:            newEwDatabase,
		newSongDatabaseVariation: newSongDatabaseVariation,
		newSongDatabaseTag:       newSongDatabaseTag,
		newTagVariation:          newTagVariation,
		newVariation:             newVariation,
	}

	matiasClient.ListenChanges(latestConnectionDate)

	return matiasClient
}

func (mc *MatiasClient) Close() {
	mc.cancel()
}
