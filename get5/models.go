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

type SQLUserData struct {
	id       int
	steam_id string
	name     string
	admin    bool
}

type GameServerData struct {
	id            int
	user_id       int
	in_use        bool
	ip_string     string
	port          int
	rcon_password string
	display_name  string
	public_server bool
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

type SQLTeamData struct {
	id          int
	user_id     int
	name        string
	flag        string
	logo        string
	auth        []byte
	tag         string
	public_team bool
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

type SQLMatchData struct {
	id             int
	user_id        int
	server_id      interface{}
	team1_id       int
	team2_id       int
	winner         interface{}
	cancelled      bool
	start_time     interface{}
	end_time       interface{}
	max_maps       int
	title          string
	skip_veto      bool
	api_key        string
	veto_mappool   string
	team1_score    int
	team2_score    int
	team1_string   string
	team2_string   string
	forfeit        bool
	plugin_version string

	//VetoMapPool []string
	//MapStats    []MapStatsData
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

type SQLMapStatsData struct {
	id          int
	match_id    int
	map_number  int
	map_name    string
	start_time  interface{}
	end_time    interface{}
	winner      interface{}
	team1_score int
	team2_score int
}

type SQLPlayerStatsData struct {
	id                int
	match_id          int
	map_id            int
	team_id           int
	steam_id          string
	name              string
	kills             int
	deaths            int
	roundsplayed      int
	assists           int
	flashbang_assists int
	teamkills         int
	suicides          int
	headshot_kills    int
	damage            int64
	bomb_plants       int
	bomb_defuses      int
	v1                int
	v2                int
	v3                int
	v4                int
	v5                int
	k1                int
	k2                int
	k3                int
	k4                int
	k5                int
	firstdeath_Ct     int
	firstdeath_t      int
	firstkill_ct      int
	firstkill_t       int
}
