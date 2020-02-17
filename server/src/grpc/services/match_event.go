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
	clients  []*pb.Get5_MatchEventServer
	stop     chan bool
}

type EventsMap struct {
	Event map[int32]*Events
	mux   sync.Mutex
}

func (e *EventsMap) AddReceiver(key int32, srv *pb.Get5_MatchEventServer) (chan bool, bool, error) { // stopper,not found,error
	if e.Event == nil {
		e.Event = make(map[int32]*Events)
	}
	notfound := false
	e.mux.Lock()
	defer e.mux.Unlock()
	if e.Event[key] == nil {
		notfound = true
		e.Event[key] = &Events{
			Finished: make(chan bool, 1),
			Event:    make(chan *pb.MatchEventReply, 1),
			clients:  make([]*pb.Get5_MatchEventServer, 0),
		}
	}
	e.Event[key].clients = append(e.Event[key].clients, srv)
	sender := *srv
	sender.Send(&pb.MatchEventReply{Event: &pb.MatchEventReply_Initialized{
		Initialized: &pb.MatchEventInitialized{},
	}})
	log.Printf("[gRPC] Added receiver client to key %d. Event has %d clients now\n", key, len(e.Event[key].clients))
	return e.Event[key].stop, notfound, nil
}

func (e *EventsMap) Write(key int32, Event *pb.MatchEventReply, finished bool) error {
	if e.Event == nil {
		e.Event = make(map[int32]*Events)
	}
	e.mux.Lock()
	defer e.mux.Unlock()
	if e.Event[key] == nil {
		e.Event[key] = &Events{
			Finished: make(chan bool, 1),
			Event:    make(chan *pb.MatchEventReply, 1),
			clients:  make([]*pb.Get5_MatchEventServer, 0),
		}
	}
	log.Printf("[gRPC] Sending event for key %d. %d clients\n", key, len(e.Event[key].clients))
	for i := 0; i < len(e.Event[key].clients); i++ {
		sender := *e.Event[key].clients[i]
		err := sender.Send(Event)
		if err != nil {
			e.Event[key].clients = e.RemoveClient(key, i)
		} else {
			log.Printf("[gRPC] Writing for key %d DONE\n", key)
			if finished {
				e.Event[key].stop <- true
			}
		}
	}
	return nil
}

func (e *EventsMap) RemoveClient(key int32, i int) []*pb.Get5_MatchEventServer {
	if i >= len(e.Event[key].clients) {
		return e.Event[key].clients
	}
	return append(e.Event[key].clients[:i], e.Event[key].clients[i+1:]...)
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
	log.Printf("[gRPC] Client connected. matchid : %d\n", matchid)
	// 関数を抜けるまでchannel待機
	stop, notfound, err := MatchesStream.AddReceiver(matchid, &srv)
	if notfound {
		log.Println("Match not found. Awaiting streams...")
	}
	if err != nil {
		return err
	}
	<-stop
	return err
}
