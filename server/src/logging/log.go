package logging

import (
	"github.com/FlowingSPDG/csgo-log"
	pb "github.com/FlowingSPDG/get5-web-go/server/src/grpc/proto"
	pbservices "github.com/FlowingSPDG/get5-web-go/server/src/grpc/services"
	"log"
	"strconv"
)

// MessageHandler handles message from CSGO Server and Gin middleware
func MessageHandler(msg csgolog.Message) {
	switch m := msg.(type) {
	case csgolog.Get5Event:
		log.Printf("Get5Event : [%v]\n", m)
		var event pb.Get5Event
		matchid, err := strconv.Atoi(m.Matchid)
		if err != nil {
			log.Printf("ERR : %v\n", err)
		}
		switch csgolog.Get5Events(m.Event) {
		case csgolog.Get5SeriesStart:
			event = pb.Get5Event_Get5SeriesStart
		case csgolog.Get5MapVeto:
			event = pb.Get5Event_Get5MapVeto
		case csgolog.Get5MapPick:
			event = pb.Get5Event_Get5MapPick
		case csgolog.Get5SidePicked:
			event = pb.Get5Event_Get5SidePicked
		case csgolog.Get5KnifeStart:
			event = pb.Get5Event_Get5KnifeStart
		case csgolog.Get5KnifeWon:
			event = pb.Get5Event_Get5KnifeWon
		case csgolog.Get5GoingLive:
			event = pb.Get5Event_Get5GoingLive
		case csgolog.Get5PlayerDeath:
			event = pb.Get5Event_Get5PlayerDeath
			log.Printf("Get5Event_Get5PlayerDeath : Params:[%v]\n", m.Params)
		case csgolog.Get5RoundEnd:
			event = pb.Get5Event_Get5RoundEnd
		case csgolog.Get5SideSwap:
			event = pb.Get5Event_Get5SideSwap
		case csgolog.Get5MapEnd:
			event = pb.Get5Event_Get5MapEnd
		case csgolog.Get5SeriesEnd:
			event = pb.Get5Event_Get5SeriesEnd
		case csgolog.Get5BackupLoaded:
			event = pb.Get5Event_Get5BackupLoaded
		case csgolog.Get5MatchConfigLoadFail:
			event = pb.Get5Event_Get5MatchConfigLoadFail
		case csgolog.Get5ClientSay:
			event = pb.Get5Event_Get5ClientSay
		case csgolog.Get5BombPlanted:
			event = pb.Get5Event_Get5BombPlanted
		case csgolog.Get5BombDefused:
			event = pb.Get5Event_Get5BombDefused
		case csgolog.Get5BombExploded:
			event = pb.Get5Event_Get5BombExploded
		case csgolog.Get5PlayerConnected:
			event = pb.Get5Event_Get5PlayerConnected
		case csgolog.Get5PlayerDisconnect:
			event = pb.Get5Event_Get5PlayerDisconnect
		case csgolog.Get5TeamReady:
			event = pb.Get5Event_Get5TeamReady
		case csgolog.Get5TeamUnready:
			event = pb.Get5Event_Get5TeamUnready
		}

		go pbservices.MatchesStream.Write(int32(matchid), &pb.MatchEventReply{
			Event: &pb.MatchEventReply_Get5Event{
				Get5Event: &pb.MatchEventGet5Event{
					Matchid: int32(matchid),
					Params:  &pb.Get5EventParams{},
					Event:   event,
				},
			},
		}, false)
	default:
		// Other log types
	}
}
