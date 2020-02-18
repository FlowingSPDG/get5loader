package logging

import (
	"database/sql"
	"fmt"
	"github.com/FlowingSPDG/csgo-log"
	"github.com/FlowingSPDG/get5-web-go/server/src/db"
	pb "github.com/FlowingSPDG/get5-web-go/server/src/grpc/proto"
	pbservices "github.com/FlowingSPDG/get5-web-go/server/src/grpc/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"sync"
)

// KillFeed Contains killer and victim's steamid
type KillFeed struct {
	KillerSteamID string
	VictimSteamID string
}

// KillFeeds contains killfeed map and locks
type KillFeeds struct {
	KillFeed map[int][]*KillFeed
	sync.Mutex
}

// Append adds killfeeds into memory
func (k *KillFeeds) Append(matchid int, killer string, victim string) {
	if killer == "" || victim == "" {
		return
	}
	if k.KillFeed == nil {
		k.KillFeed = make(map[int][]*KillFeed, 0)
	}
	k.Lock()
	defer k.Unlock()
	if _, ok := k.KillFeed[matchid]; !ok {
		k.KillFeed[matchid] = make([]*KillFeed, 0, 10)
	}
	k.KillFeed[matchid] = append(k.KillFeed[matchid], &KillFeed{
		KillerSteamID: killer,
		VictimSteamID: victim,
	})
	log.Printf("Adding kill feeds for matchid %d, Current stored kills:[%v]\n,", matchid, k.KillFeed[matchid])
}

// Clear Killfeeds lists
func (k *KillFeeds) Clear(matchid int) error {
	if k.KillFeed == nil {
		return fmt.Errorf("Match not found")
	}
	log.Printf("Clearing kill feeds for matchid %d\n,", matchid)
	k.Lock()
	defer k.Unlock()
	if _, ok := k.KillFeed[matchid]; !ok {
		return fmt.Errorf("Match not found")
	}
	k.KillFeed[matchid] = make([]*KillFeed, 0, 10)
	runtime.GC()
	return nil
}

// Register Killfeeds into DB
func (k *KillFeeds) Register(matchid int, mapnumber int, winner string, winnerside string) error {
	if k.KillFeed == nil {
		return fmt.Errorf("Match not found")
	}
	k.Lock()
	defer k.Unlock()
	if _, ok := k.KillFeed[matchid]; !ok {
		return fmt.Errorf("Match not found")
	}
	sqlwinner := sql.NullString{}
	sqlwinner.Scan(winner)
	sqlwinnerside := sql.NullString{}
	sqlwinnerside.Scan(winnerside)
	round := &db.RoundStatsData{
		MatchID:    matchid,
		MapNumber:  mapnumber,
		Winner:     sqlwinner,
		WinnerSide: sqlwinnerside,
	}
	logs := k.KillFeed[matchid]
	if len(logs) >= 1 {
		sqlfirstkiller := sql.NullString{}
		sqlfirstvictim := sql.NullString{}
		sqlfirstkiller.Scan(k.KillFeed[matchid][0].KillerSteamID)
		sqlfirstvictim.Scan(k.KillFeed[matchid][0].VictimSteamID)
		round.FirstKillerSteamID = sqlfirstkiller
		round.FirstVictimSteamID = sqlfirstvictim
	}
	if len(logs) >= 2 {
		sqlsecondkiller := sql.NullString{}
		sqlsecondvictim := sql.NullString{}
		sqlsecondkiller.Scan(k.KillFeed[matchid][1].KillerSteamID)
		sqlsecondvictim.Scan(k.KillFeed[matchid][1].VictimSteamID)
		round.SecondKillerSteamID = sqlsecondkiller
		round.SecondVictimSteamID = sqlsecondvictim
	}
	if len(logs) >= 3 {
		sqlthirdkiller := sql.NullString{}
		sqlthirdvictim := sql.NullString{}
		sqlthirdkiller.Scan(k.KillFeed[matchid][2].KillerSteamID)
		sqlthirdvictim.Scan(k.KillFeed[matchid][2].VictimSteamID)
		round.ThirdKillerSteamID = sqlthirdkiller
		round.ThirdVictimSteamID = sqlthirdvictim
	}
	if len(logs) >= 4 {
		sqlfourthkiller := sql.NullString{}
		sqlfourthvictim := sql.NullString{}
		sqlfourthkiller.Scan(k.KillFeed[matchid][3].KillerSteamID)
		sqlfourthvictim.Scan(k.KillFeed[matchid][3].VictimSteamID)
		round.FourthKillerSteamID = sqlfourthkiller
		round.FourthVictimSteamID = sqlfourthvictim
	}
	if len(logs) >= 5 {
		sqlfifthkiller := sql.NullString{}
		sqlfifthvictim := sql.NullString{}
		sqlfifthkiller.Scan(k.KillFeed[matchid][4].KillerSteamID)
		sqlfifthvictim.Scan(k.KillFeed[matchid][4].VictimSteamID)
		round.FifthKillerSteamID = sqlfifthkiller
		round.FifthVictimSteamID = sqlfifthvictim
	}

	if len(logs) >= 6 {
		sqlsixthkiller := sql.NullString{}
		sqlsixthvictim := sql.NullString{}
		sqlsixthkiller.Scan(k.KillFeed[matchid][5].KillerSteamID)
		sqlsixthvictim.Scan(k.KillFeed[matchid][5].VictimSteamID)
		round.SixthKillerSteamID = sqlsixthkiller
		round.SixthVictimSteamID = sqlsixthvictim
	}

	if len(logs) >= 7 {
		sqlseventhkiller := sql.NullString{}
		sqlseventhvictim := sql.NullString{}
		sqlseventhkiller.Scan(k.KillFeed[matchid][6].KillerSteamID)
		sqlseventhvictim.Scan(k.KillFeed[matchid][6].VictimSteamID)
		round.SeventhKillerSteamID = sqlseventhkiller
		round.SeventhVictimSteamID = sqlseventhvictim
	}

	if len(logs) >= 8 {
		sqleighthkiller := sql.NullString{}
		sqleighthvictim := sql.NullString{}
		sqleighthkiller.Scan(k.KillFeed[matchid][7].KillerSteamID)
		sqleighthvictim.Scan(k.KillFeed[matchid][7].VictimSteamID)
		round.EighthKillerSteamID = sqleighthkiller
		round.EighthVictimSteamID = sqleighthvictim
	}

	if len(logs) >= 9 {
		sqlninthkiller := sql.NullString{}
		sqlninthvictim := sql.NullString{}
		sqlninthkiller.Scan(k.KillFeed[matchid][8].KillerSteamID)
		sqlninthvictim.Scan(k.KillFeed[matchid][8].VictimSteamID)
		round.NinthKillerSteamID = sqlninthkiller
		round.NinthVictimSteamID = sqlninthvictim
	}

	if len(logs) >= 10 {

		sqltenthkiller := sql.NullString{}
		sqltenthkictim := sql.NullString{}
		sqltenthkiller.Scan(k.KillFeed[matchid][9].KillerSteamID)
		sqltenthkictim.Scan(k.KillFeed[matchid][9].VictimSteamID)
		round.TenthKillerSteamID = sqltenthkiller
		round.TenthVictimSteamID = sqltenthkictim
	}

	stats, err := round.Register(matchid, mapnumber)
	if err != nil {
		log.Printf("Failed to register round info : %v\n", err)
		return err
	}
	log.Printf("STATS : %v\n", stats)
	return nil
}

var (
	// KillLogs contains Killers and victims
	KillLogs KillFeeds
)

func init() {
	KillLogs = KillFeeds{
		KillFeed: make(map[int][]*KillFeed, 0),
	}
}

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
	// log.Printf("SRCDS Message handeler for match %d. Msg : [%v]\n", matchid, msg)
	switch m := msg.(type) {
	case csgolog.PlayerKill:
		KillLogs.Append(matchid, m.Attacker.SteamID, m.Victim.SteamID)
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
		// log.Printf("Get5Event : [%v]\n", m)
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
			// this misses after-round kill(such as saving), but registering on FreezeTime can't get get5's matchid/matchparams/mapnumber...
			err := KillLogs.Register(matchid, m.Params.MapNumber, m.Params.Winner, m.Params.WinnerSide)
			if err != nil {
				log.Printf("Failed to register match info : %v\n", err)
			}
			KillLogs.Clear(matchid)
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
