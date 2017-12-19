package service

import (
	"log"

	"github.com/koodinikkarit/matias/matias_service"
)

type MatiasServer struct {
	EwDatabases []*MatiasService.EwDatabase
}

func (st *MatiasServer) ListenChanges(
	req *MatiasService.ListenEventsRequest,
	stream MatiasService.Matias_ListenChangesServer,
) error {
	log.Printf("Listen start hostname %v key %v", req.HostName, req.Key)
	newVariation1 := &MatiasService.Variation{
		Id:   56,
		Name: "laulu1",
		Text: "Sisalto1",
	}
	// newVariation2 := MatiasService.Variation{
	// 	Name: "laulu2",
	// 	Text: "Sisalto2",
	// }

	stream.Send(&MatiasService.EventItem{
		EventMessage: &MatiasService.EventItem_NewVariation{
			NewVariation: newVariation1,
		},
	})
	// for {
	// 	<-time.NewTimer(time.Second * 2).C
	// 	for _, ewDatabase := range st.EwDatabases {
	// 		log.Printf("Sending ewdatabase %v", ewDatabase)
	// 		err := stream.Send(&MatiasService.EventItem{
	// 			EventMessage: &MatiasService.EventItem_NewEwDatabase{
	// 				NewEwDatabase: ewDatabase,
	// 			},
	// 		})

	// 		if err != nil {
	// 			log.Printf("Sending error %v", err)
	// 			return err
	// 		}
	// 	}
	// }
	return nil
}
