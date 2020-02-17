package services

import (
	"fmt"
	"time"
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
	Finished  chan (bool) // Stream is closed or not
	Event     chan (*pb.MatchEventReply)
	Receivers uint16 // number of receivers
}

type EventsMap struct {
	Event map[int32]*Events
	mux   sync.Mutex
}

func (e *EventsMap) Write(key int32, Event *pb.MatchEventReply, finished bool) error {
	if e.Event == nil {
		e.Event = make(map[int32]*Events)
	}
	//e.mux.Lock()
	//defer e.mux.Unlock()
	if e.Event[key] == nil {
		e.Event[key] = &Events{
			Finished:  make(chan bool, 1),
			Event:     make(chan *pb.MatchEventReply, 1),
			Receivers: 0,
		}
	}
	for i := uint16(0); i < e.Event[key].Receivers; i++ {
		log.Printf("[gRPC] Writing for key %d\n", key)
		e.Event[key].Finished <- finished
		e.Event[key].Event <- Event
		log.Printf("[gRPC] Writing for key %d DONE\n", key)
	}
	return nil
}

func (e *EventsMap) Read(key int32) (*pb.MatchEventReply, bool, error) {
	log.Printf("[gRPC] Reading for key %d\n", key)
	e.mux.Lock()
	defer e.mux.Unlock()
	log.Printf("[gRPC] Received Event on Read(). Event : %v Finished : %v\n", ev, fi)
	if val, ok := e.Event[key]; ok {
		for i := uint16(1); i < val.Receivers; i++ {
			ev := <-val.Event
			fi := <-val.Finished
			log.Printf("[gRPC] Reading for key %d DONE\n", key)
		}
		return <-val.Event, <-val.Finished, nil
	}
	return nil, false, fmt.Errorf("Not Found")
}

func (e *EventsMap) AddReceiver(key int32) error {
	log.Printf("[gRPC] Adding for key %d\n", key)
	e.mux.Lock()
	defer e.mux.Unlock()
	if val, ok := e.Event[key]; ok {
		val.Receivers++
		return nil
	}
	return fmt.Errorf("Not Found")
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
	log.Printf("[gRPC] MatchEvent. matchid : %d\n", matchid)

	initev := &pb.MatchEventReply{
		Event: &pb.MatchEventReply_Initialized{
			Initialized: &pb.MatchEventInitialized{},
		},
	}

	lastevent := initev
	for { //go func(){}() ?
		MatchesStream.AddReceiver(matchid)
		senddata, finished, err := MatchesStream.Read(matchid)
		if err != nil {
			log.Printf("[gRPC] Unknown ERROR : %v\n", err)
			time.Sleep(time.Millisecond * 200)
			continue
		}
		if lastevent != senddata {
			log.Printf("[gRPC] Data Updated! Sending data : %v\n", senddata)
			err = srv.Send(senddata)
			if err != nil {
				log.Println(err)
				return err
			}
			lastevent = senddata
			if finished {
				log.Println("[gRPC] Stream finished")
				MatchesStream.CloseChannels(matchid)
				return nil
			}
		}
	}
}
