package models

import (
	"database/sql"
	"fmt"

	_ "github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
	_ "github.com/solovev/steam_go"

	//_ "html/template"
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
	Id       int
	Steam_id string
	Name     string
	Admin    bool
}

type GameServerData struct {
	Id            int
	User_id       int
	In_use        bool
	Ip_string     string
	Port          int
	Rcon_password string
	Display_name  string
	Public_server bool
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
	Id          int
	User_id     int
	Name        string
	Flag        string
	Logo        string
	Auth        []byte
	Tag         string
	Public_team bool
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
	Id             int
	User_id        int
	Server_id      interface{}
	Team1_id       int
	Team2_id       int
	Winner         sql.NullInt64
	Cancelled      bool
	Start_time     interface{}
	End_time       interface{}
	Max_maps       int
	Title          string
	Skip_veto      bool
	Api_key        string
	Veto_mappool   string
	Team1_score    int
	Team2_score    int
	Team1_string   string
	Team2_string   string
	Forfeit        bool
	Plugin_version string

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
	Id          int
	Match_id    int
	Map_number  int
	Map_name    string
	Start_time  interface{}
	End_time    interface{}
	Winner      interface{}
	Team1_score int
	Team2_score int
}

type SQLPlayerStatsData struct {
	Id                int
	Match_id          int
	Map_id            int
	Team_id           int
	Steam_id          string
	Name              string
	Kills             int
	Deaths            int
	Roundsplayed      int
	Assists           int
	Flashbang_assists int
	Teamkills         int
	Suicides          int
	Headshot_kills    int
	Damage            int64
	Bomb_plants       int
	Bomb_defuses      int
	V1                int
	V2                int
	V3                int
	V4                int
	V5                int
	K1                int
	K2                int
	K3                int
	K4                int
	K5                int
	Firstdeath_Ct     int
	Firstdeath_t      int
	Firstkill_ct      int
	Firstkill_t       int
}

type MatchesPageData struct {
	LoggedIn bool
	UserName string
	UserID   int
	Matches  []SQLMatchData
}

type TeamPageData struct {
	LoggedIn   bool
	IsYourTeam bool
	Teams      []SQLTeamData
	tp         string
	test       string
}

func get_flag_html(country string, scale int) string {
	width := 32 * scale
	height := 21 * scale
	html := fmt.Sprintf(`<img src="%s%s"  width="%d" height="%d">`, "/static/img/valve_flags/", country, width, height)
	return html
}
