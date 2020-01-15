package services

import (
	"time"
	// "github.com/FlowingSPDG/get5-web-go/server/src/api"
	// "github.com/FlowingSPDG/get5-web-go/server/src/db"
	pb "github.com/FlowingSPDG/get5-web-go/server/src/grpc/proto"
	// "github.com/FlowingSPDG/get5-web-go/server/src/util"
	// "google.golang.org/grpc"
	// "log"
	"fmt"
	"sync"
)

type Events struct {
	Finished bool // Stream is closed or not
	Event    *pb.MatchEventReply
}

type EventsMap struct {
	Event map[int32]*Events
	mux   sync.Mutex
}

func (e *EventsMap) Write(key int32, Event *pb.MatchEventReply) {
	e.mux.Lock()
	defer e.mux.Unlock()
	if e.Event == nil {
		e.Event = make(map[int32]*Events)
	}
	if e.Event[key] == nil {
		e.Event[key] = &Events{}
	}
	e.Event[key].Event = Event
}

func (e *EventsMap) Read(key int32) *pb.MatchEventReply {
	e.mux.Lock()
	defer e.mux.Unlock()
	if e.Event[key] == nil {
		e.Event[key] = &Events{}
	}
	return e.Event[key].Event
}

var (
	MatchesStream EventsMap
)

func init() {
	MatchesStream = EventsMap{}
}

func (s Server) MatchEvent(req *pb.MatchEventRequest, srv pb.Get5_MatchEventServer) error {
	matchid := req.GetMatchid()
	fmt.Printf("MatchEvent. matchid : %d\n", matchid)
	MatchesStream.Write(matchid, &pb.MatchEventReply{
		Event: &pb.MatchEventReply_Initialized{
			Initialized: &pb.MatchEventInitialized{},
		},
	}) // initialize?
	err := srv.Send(MatchesStream.Read(matchid))
	if err != nil {
		return err
	}

	lastevent := &pb.MatchEventReply{}
	for { //go func(){}() ?
		senddata := MatchesStream.Read(matchid)
		if lastevent != senddata {
			fmt.Printf("Data Updated! Sending data : %v\n", senddata)
			err = srv.Send(senddata)
			if err != nil {
				return err
			}
			lastevent = senddata
		}
	}
}
