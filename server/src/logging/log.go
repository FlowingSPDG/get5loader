package logging

import (
	"fmt"
	"github.com/FlowingSPDG/csgo-log"
	"github.com/FlowingSPDG/get5-web-go/server/src/db"
	pb "github.com/FlowingSPDG/get5-web-go/server/src/grpc/proto"
	pbservices "github.com/FlowingSPDG/get5-web-go/server/src/grpc/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// MessageHandler handles message from CSGO Server and Gin middleware
func MessageHandler(msg csgolog.Message, c *gin.Context) {
	matchidstr := c.Params.ByName("matchID")
	matchid, err := strconv.Atoi(matchidstr)
	if err != nil {
		log.Printf("ERR : %v\n", err)
	}
	auth := c.Params.ByName("auth")
	match := &db.MatchData{}
	rec := db.SQLAccess.Gorm.First(&match, matchidstr)
	if rec.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if match.APIKey != auth {
		c.AbortWithError(http.StatusForbidden, fmt.Errorf("Wrong auth"))
		return
	}
	log.Printf("SRCDS Message handeler for match %d. Msg : [%v]\n", matchid, msg)
	switch m := msg.(type) {
	case csgolog.PlayerKill:
		go pbservices.MatchesStream.Write(int32(matchid), &pb.MatchEventReply{
			Event: &pb.MatchEventReply_Playerkill{
				Playerkill: &pb.MatchEventPlayerKill{
					Matchid: int32(matchid),
					Attacker: &pb.Player{
						Name:    m.Attacker.Name,
						Id:      int32(m.Attacker.ID),
						Steamid: m.Attacker.SteamID,
						Side:    m.Attacker.Side,
					},
					AttackerPosition: &pb.Position{
						X: int32(m.AttackerPosition.X),
						Y: int32(m.AttackerPosition.Y),
						Z: int32(m.AttackerPosition.Z),
					},
					Victim: &pb.Player{
						Name:    m.Victim.Name,
						Id:      int32(m.Victim.ID),
						Steamid: m.Victim.SteamID,
						Side:    m.Victim.Side},
					VictimPosition: &pb.Position{
						X: int32(m.VictimPosition.X),
						Y: int32(m.VictimPosition.Y),
						Z: int32(m.VictimPosition.Z),
					},
					Weapon:     m.Weapon,
					Headhot:    m.Headshot,
					Penetrated: m.Penetrated,
				},
			},
		}, false)
	case csgolog.Get5Event:
		log.Printf("Get5Event : [%v]\n", m)
		var event pb.Get5Event
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
					Params: &pb.Get5EventParams{
						MapNumber:        int32(m.Params.MapNumber),
						MapName:          m.Params.MapName,
						Team1Name:        m.Params.Team1Name,
						Team1Score:       int32(m.Params.Team1Score),
						Team1SeriesScore: int32(m.Params.Team1SeriesScore),
						Team2Name:        m.Params.Team2Name,
						Team2Score:       int32(m.Params.Team2Score),
						Team2SeriesScore: int32(m.Params.Team2SeriesScore),
						Headshot:         int32(m.Params.Headshot),
						Weapon:           m.Params.Weapon,
						Reason:           int32(m.Params.Reason),
						Message:          m.Params.Message,
						File:             m.Params.File,
						Site:             int32(m.Params.Site),
						Stage:            m.Params.Stage,
						Attacker:         m.Params.Attacker, // FlowingSPDG<5><STEAM_1:1:55894410><>
						Victim:           m.Params.Victim,
						Winner:           m.Params.Winner,
						WinnerSide:       m.Params.WinnerSide,
					},
					Event: event,
				},
			},
		}, false)
	default:
		// Other log types
	}
}
