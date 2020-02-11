package services

import (
	"fmt"
	// "github.com/FlowingSPDG/get5-web-go/server/src/api"
	// "github.com/FlowingSPDG/get5-web-go/server/src/db"
	pb "github.com/FlowingSPDG/get5-web-go/server/src/grpc/proto"
	// "github.com/FlowingSPDG/get5-web-go/server/src/util"
	// "google.golang.org/grpc"
	// "log"
	"log"
	"sync"
)

type Events struct {
	Finished chan (bool) // Stream is closed or not
	Event    chan (*pb.MatchEventReply)
}

type EventsMap struct {
	Event map[int32]*Events
	mux   sync.Mutex
}

func (e *EventsMap) Write(key int32, Event *pb.MatchEventReply, finished bool) error {
	e.mux.Lock()
	defer e.mux.Unlock()
	if e.Event == nil {
		return fmt.Errorf("Not Found")
	}
	if e.Event[key] == nil {
		return fmt.Errorf("Not Found")
	}
	e.Event[key].Finished <- finished
	e.Event[key].Event <- Event
	return nil
}

func (e *EventsMap) Read(key int32) (*pb.MatchEventReply, bool, error) {
	e.mux.Lock()
	defer e.mux.Unlock()
	if e.Event[key] == nil {
		return nil, false, fmt.Errorf("Not Found")
	}
	return <-e.Event[key].Event, <-e.Event[key].Finished, nil
}

func (e *EventsMap) CloseChannels(key int32) error {
	e.mux.Lock()
	defer e.mux.Unlock()
	if e.Event == nil {
		return fmt.Errorf("Not Found")
	}
	if e.Event[key] == nil {
		return fmt.Errorf("Not Found")
	}
	for _, v := range e.Event {
		close(v.Event)
	}
	return nil
}

var (
	MatchesStream EventsMap
)

func init() {
	MatchesStream = EventsMap{}
}

func (s Server) MatchEvent(req *pb.MatchEventRequest, srv pb.Get5_MatchEventServer) error {
	matchid := req.GetMatchid()
	log.Printf("MatchEvent. matchid : %d\n", matchid)
	MatchesStream.Write(matchid, &pb.MatchEventReply{
		Event: &pb.MatchEventReply_Initialized{
			Initialized: &pb.MatchEventInitialized{},
		},
	}, false) // initialize?
	err := srv.Send(&pb.MatchEventReply{
		Event: &pb.MatchEventReply_Initialized{},
	})
	if err != nil {
		log.Println(err)
		return err
	}

	lastevent := &pb.MatchEventReply{
		Event: &pb.MatchEventReply_Initialized{},
	}
	for { //go func(){}() ?
		senddata, finished, err := MatchesStream.Read(matchid)
		if err != nil {
			log.Printf("Unknown ERROR : %v\n", err)
		}
		if lastevent.Event != senddata.Event {
			log.Printf("Data Updated! Sending data : %v\n", senddata)
			err = srv.Send(senddata)
			if err != nil {
				log.Println(err)
				return err
			}
			lastevent = senddata
			if finished {
				log.Println("Stream finished")
				MatchesStream.CloseChannels(matchid)
				return nil
			}
		}
	}
	return nil
}
