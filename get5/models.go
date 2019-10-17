package get5

import (
	_ "fmt"
	_ "github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
	_ "github.com/solovev/steam_go"
	_ "html/template"
	_ "net/http"
	_ "strconv"
	"time"
)

type UserData struct {
	ID      int
	SteamID string
	Name    string
	Admin   bool
	Servers []GameServerData
	Teams   []TeamData
	Matches []MatchData
}

type GameServerData struct {
	ID           int
	UserID       int
	DisplayName  string
	IPstring     string
	port         int
	RconPassword string
	InUse        bool
	PublicServer bool
}

type TeamData struct {
	ID         int
	UserID     int
	Name       string
	Tag        string
	Flag       string
	Logo       string
	Auths      []string
	PublicTeam bool
}

type MatchData struct {
	ID            int64
	UserID        int64
	ServerID      int64
	Team1ID       int64
	Team1Score    int
	Team1String   string
	Team2ID       int64
	Team2Score    int
	Team2String   string
	winner        int64
	PluginVersion string
	forfeit       bool
	cancelled     bool
	StartTime     time.Time
	EndTime       time.Time
	MaxMaps       int
	title         string
	SkipVeto      bool
	APIKey        string

	VetoMapPool []string
	MapStats    []MapStatsData
}

type MapStatsData struct {
	ID         int
	MatchID    int
	MapNumber  int
	MapName    string
	StartTime  time.Time
	EndTIme    time.Time
	Winner     int
	Team1Score int
	Team2Score int
}
