package models

import (
	"database/sql"
	"fmt"
	//"github.com/Philipp15b/go-steam"
	"github.com/FlowingSPDG/go-steamapi"
	"github.com/go-ini/ini"
	_ "github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
	//"github.com/solovev/steam_go"
	//_ "html/template"
	_ "net/http"
	"strconv"
)

var (
	SteamAPIKey = ""
)

func init() {
	c, _ := ini.Load("config.ini")
	SteamAPIKey = c.Section("Steam").Key("APIKey").MustString("")
}

type UserData struct {
	ID      int
	SteamID string
	Name    string
	Admin   bool
	Servers []GameServerData
	Teams   []TeamData
	Matches []MatchData
}

func (u *UserData) GetOrCreate(s *sql.DB, steamid string) (*UserData, error) {
	var EmptyUser *UserData
	var SQLUserData SQLUserData
	err := s.Ping()
	if err != nil {
		return EmptyUser, err
	}

	//データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
	q := fmt.Sprintf("SELECT * FROM `user` WHERE steam_id=%s LIMIT 1 ", steamid)
	fmt.Println(q)
	rows, err := s.Query(q)
	if err != nil {
		fmt.Println("USER NOT EXIST")
		return EmptyUser, err
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&SQLUserData.Id, &SQLUserData.Steam_id, &SQLUserData.Name, &SQLUserData.Admin)
		if err != nil {
			return EmptyUser, err
		}
		fmt.Printf("UserData : %v", SQLUserData)
		u.ID = SQLUserData.Id
		u.Name = SQLUserData.Name
		u.SteamID = SQLUserData.Steam_id
		u.Admin = SQLUserData.Admin
		return u, nil
	} else {
		fmt.Println("USER NOT EXIST. REGISTER!")
		//playerinfo, err := steam_go.GetPlayerSummaries(steamid, SteamAPIKey)
		steamid64, _ := strconv.Atoi(steamid)
		steamid64arr := []uint64{uint64(steamid64)}
		playersumamry, err := steamapi.GetPlayerSummaries(steamid64arr, SteamAPIKey)
		fmt.Printf("\nsteamid : %s / SteamAPIKey : %s\n", steamid, SteamAPIKey)
		fmt.Println(playersumamry)

		if err != nil { // fix here
			fmt.Println(err)
			return EmptyUser, err
		}

		q := fmt.Sprintf("INSERT INTO `user` (steam_id,name,admin) values (%s,'%s',0);", steamid, playersumamry[0].PersonaName)
		fmt.Println(q)
		res, err := s.Exec(q)
		if err != nil {
			return EmptyUser, err
		}
		fmt.Println(res)
		rows, err := s.Query(q)
		if err != nil {
			fmt.Println("UNKNOWN FAIL")
			return EmptyUser, err
		}
		defer rows.Close()
		if rows.Next() {
			err := rows.Scan(&SQLUserData.Id, &SQLUserData.Steam_id, &SQLUserData.Name, &SQLUserData.Admin)
			if err != nil {
				return EmptyUser, err
			}
			fmt.Printf("UserData : %v", SQLUserData)
			u.ID = SQLUserData.Id
			u.Name = SQLUserData.Name
			u.SteamID = SQLUserData.Steam_id
			u.Admin = SQLUserData.Admin
			return u, nil
		}
		return EmptyUser, err
	}
}

/*
func (u *UserData) GetURL() {

}
*/

func (u *UserData) GetSteamURL() string {
	return "http://steamcommunity.com/profiles/" + u.SteamID
}

func (u *UserData) get_recent_matches(limit int) string {
	return "http://steamcommunity.com/profiles/" + u.SteamID
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
	Auth        []byte // not []byte...?
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
	StartTime     sql.NullTime
	EndTime       sql.NullTime
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
	Server_id      sql.NullInt64
	Team1_id       int
	Team2_id       int
	Winner         sql.NullInt64
	Cancelled      bool
	Start_time     sql.NullTime
	End_time       sql.NullTime
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
	StartTime  sql.NullTime
	EndTIme    sql.NullTime
	Winner     int
	Team1Score int
	Team2Score int
}

type SQLMapStatsData struct {
	Id          int
	Match_id    int
	Map_number  int
	Map_name    string
	Start_time  sql.NullTime
	End_time    sql.NullTime
	Winner      sql.NullInt64
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

type TeamsPageData struct {
	LoggedIn   bool
	User       SQLUserData
	IsYourTeam bool
	Teams      []SQLTeamData
}

type TeamPageData struct {
	LoggedIn   bool
	IsYourTeam bool
	Team       SQLTeamData
}

type UserPageData struct {
	LoggedIn bool
	User     SQLUserData
	Teams    []SQLTeamData
}

type MyserversPageData struct {
	Servers  []GameServerData
	LoggedIn bool
}

func get_flag_html(country string, scale int) string {
	width := 32 * scale
	height := 21 * scale
	html := fmt.Sprintf(`<img src="%s%s"  width="%d" height="%d">`, "/static/img/valve_flags/", country, width, height)
	return html
}
