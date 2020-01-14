package services

import (
	// "github.com/FlowingSPDG/get5-web-go/server/src/api"
	// "github.com/FlowingSPDG/get5-web-go/server/src/db"
	pb "github.com/FlowingSPDG/get5-web-go/server/src/grpc/proto"
	// "github.com/FlowingSPDG/get5-web-go/server/src/util"
	// "google.golang.org/grpc"
	// "log"
	"fmt"
)

type Events struct {
	Finished bool // Stream is closed or not
	Event    *pb.MatchEventReply
}

var (
	MatchesStream map[int32]*Events // MatchesStream[MatchID].Servers[0].Send()
)

func init() {
	MatchesStream = make(map[int32]*Events)
}

func (s Server) MatchEvent(req *pb.MatchEventRequest, srv pb.Get5_MatchEventServer) error {
	matchid := req.GetMatchid()
	fmt.Printf("MatchEvent. matchid : %d\n", matchid)
	if _, ok := MatchesStream[matchid]; !ok {
		MatchesStream[matchid] = &Events{
			Finished: false,
			Event: &pb.MatchEventReply{
				Event: &pb.MatchEventReply_Initialized{},
			},
		} // initialize
	}

	lastevent := &pb.MatchEventReply{}
	for { //go func(){}() ?
		if !MatchesStream[matchid].Finished && lastevent != MatchesStream[matchid].Event {
			fmt.Println("sending data : %V\n", MatchesStream[matchid].Event)
			srv.Send(MatchesStream[matchid].Event)
			lastevent = MatchesStream[matchid].Event
		}
		if MatchesStream[matchid].Finished {
			return nil // closes stream
		}
	}
}
