package services

import (
	// "github.com/FlowingSPDG/get5-web-go/server/src/api"
	// "github.com/FlowingSPDG/get5-web-go/server/src/db"
	pb "github.com/FlowingSPDG/get5-web-go/server/src/grpc/proto"
	// "github.com/FlowingSPDG/get5-web-go/server/src/util"
	// "context"
	// "google.golang.org/grpc"
	// "log"
	"fmt"
)

func (s Server) MatchEvent(req *pb.MatchEventRequest, srv pb.Get5_MatchEventServer) error {
	fmt.Printf("MatchEvent. matchid : %d\n", req.Matchid)
	return nil
}
