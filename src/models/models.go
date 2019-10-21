package models

import (
	"database/sql"
	"fmt"
	"strings"
	//"github.com/Philipp15b/go-steam"
	"github.com/FlowingSPDG/go-steamapi"
	"github.com/go-ini/ini"
	_ "github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
	steam "github.com/kidoman/go-steam"
	//"github.com/solovev/steam_go"
	//_ "html/template"
	"math"
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

func (g *GameServerData) Create(userid int, display_name string, ip_string string, port int, rcon_password string, public_server bool) *GameServerData {
	g.User_id = userid
	g.Display_name = display_name
	g.Ip_string = ip_string
	g.Port = port
	g.Rcon_password = rcon_password
	g.Public_server = public_server
	// ADD TO DB
	return g
}

func (g *GameServerData) SendRcon(cmd string) (string, error) {
	o := &steam.ConnectOptions{RCONPassword: g.Rcon_password}
	rcon, err := steam.Connect(g.Ip_string, o)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer rcon.Close()

	resp, err := rcon.Send(cmd)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return resp, nil
}

func (g *GameServerData) GetHostPort() string {
	return fmt.Sprintf("%s:%d", g.Ip_string, g.Port)
}

/*
func (g *GameServerData) GetDisplay() string {

}*/

type TeamData struct {
	ID     int
	UserID int
	Name   string
	Tag    string
	Flag   string
	Logo   string
	//Auths      []string
	Auths      []byte
	PublicTeam bool
}

func (t *TeamData) Create(userid int, name string, tag string, flag string, logo string, auths []byte, public_team bool) *TeamData {
	t.UserID = userid
	t.Name = name
	t.Tag = tag
	t.Flag = flag
	t.Logo = logo
	t.Auths = auths
	t.PublicTeam = public_team
	return t
}

func (t *TeamData) SetData(name string, tag string, flag string, logo string, auths []byte, public_team bool) *TeamData {
	t.Name = name
	t.Tag = tag
	t.Flag = flag
	t.Logo = logo
	t.Auths = auths
	t.PublicTeam = public_team
	return t
}

func (t *TeamData) CanEdit(userid int) bool {
	if userid == 0 {
		return false
	} else if t.UserID == userid {
		return true
	}
	return false
}

/*
func (t *TeamData) GetPlayers(userid int) bool {
	var results []string
	//Py
	for steam64 in self.auths:
            if steam64:
                name = get_steam_name(steam64)
                if not name:
                    name = ''

                results.append((steam64, name))
		return results
	//?? TODO
}
*/

func (t *TeamData) CanDelete(userid int) bool {
	if t.CanEdit(userid) == false {
		return false
	}
	return len(t.GetRecentMatches()) == 0
}

/*
func (t *TeamData) GetRecentMatches(limit int) []SQLMatchData {

}
*/

/*
func (t *TeamData) GetVSMatchResult(matchid int) []SQLMatchData {

}
*/

func (t *TeamData) GetFlagHTML(scale float32) string {
	width := 32.0 * scale
	height := 21.0 * scale
	html := fmt.Sprintf(`<img src="%s%s"  width="%f" height="%d">`, "../static/img/valve_flags", t.Flag, width, height)
	return html
}

/*
func (t *TeamData) GetLogoHtml(scale float32) string {

}
*/

/*
func (t *TeamData) GetURL(scale float32) string {

}
*/

/*
func (t *TeamData) GetNameURLHtml(scale float32) string {

}
*/

/*
func (t *TeamData) GetLogoOrFlagHtml(scale float32) string {

}
*/

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
	Winner        int64
	PluginVersion string
	Forfeit       bool
	Cancelled     bool
	StartTime     sql.NullTime
	EndTime       sql.NullTime
	MaxMaps       int
	Title         string
	SkipVeto      bool
	APIKey        string

	VetoMapPool []string
	MapStats    []MapStatsData
}

func (m *MatchData) Create(userid int64, team1_id int64, team2_id int64, team1_string string, team2_string string, max_maps int, skip_veto bool, title string, veto_mappool string, server_id int64) *MatchData {
	m.UserID = userid
	m.Team1ID = team1_id
	m.Team2ID = team2_id
	m.SkipVeto = skip_veto
	m.Title = title
	m.VetoMapPool = strings.Split(veto_mappool, ",") // should work
	m.ServerID = server_id
	m.MaxMaps = max_maps
	m.APIKey = "" //random?
	return m
}

func (m *MatchData) GetStatusString(ShowWinner bool) string {
	if m.Pending() {
		return "Pending"
	} else if m.Live() {
		teams1core, team2score := m.GetCurrentScore()
		return fmt.Sprintf("Live, %d:%d", teams1core, team2score)
	} else if m.Finished() {
		teams1core, team2score := m.GetCurrentScore()
		minscore := math.Min(float64(teams1core), float64(team2score))
		maxscore := math.Max(float64(teams1core), float64(team2score))
		ScoreString := fmt.Sprintf("%f:%f", maxscore, minscore)
		if !ShowWinner {
			return "Finished"
		} else if m.Winner == m.Team1ID {
			return fmt.Sprintf("Won %s by %f", ScoreString, m.GetTeam1().Name)
		} else if m.Winner == m.Team2ID {
			return fmt.Sprintf("Won %s by %f", ScoreString, m.GetTeam2().Name)
		} else {
			return fmt.Sprintf("Tied %s", ScoreString)
		}
	} else {
		return "Cancelled"
	}
}

func (m *MatchData) GetVSString() string {
	team1 := m.GetTeam1()
	team2 := m.GetTeam2()
	scores := m.GetCurrentScore()
	str := fmt.Sprintf("%s VS %s (%d:%d)", team1.GetNameURLHtml(), team2.GetNameURLHtml(), scores[0], scores[1])
	return str
}

func (m *MatchData) Finalized() bool {
	return m.Cancelled || m.Finished()
}

func (m *MatchData) Pending() bool {
	return m.StartTime == nil && !m.Cancelled
}

func (m *MatchData) Finished() bool {
	return m.EndTime == nil && !m.Cancelled
}

func (m *MatchData) Live() bool {
	return m.StartTime != nil && m.EndTime == nil && !m.Cancelled()
}

func (m *MatchData) GetServer() int {
	return m.ServerID
}

func (m *MatchData) GetCurrentScore() (int, int) {
	if m.MaxMaps == 1 {
		if m.MapStats[0] == nil { // check ok?
			return 0, 0
		} else {
			return m.MapStats[0].Team1_score, m.MapStats[0].Team1_score
		}
	} else {
		return m.Team1Score, m.Team2Score
	}
}

/*func (m *MatchData) SendToServer() {
	// get5_loadmatch_url things
}*/

/*func (m *MatchData) GetTeam1() TeamData {
	return Teamdata by using mysql
}*/

/*func (m *MatchData) GetTeam2() TeamData {
	return Teamdata by using mysql
}*/

/*func (m *MatchData) GetUser() UserData {
	//return m.UserID
}*/

/*func (m *MatchData) GetWinner() TeamData {
	//return m.UserID
}*/

/*func (m *MatchData) GetLoser() TeamData {
	//return m.UserID
}*/

/*func (m *MatchData) BuildMatchDict() TeamData {
	//return m.UserID //get5 thing??
}*/

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

/*func (m *MapStatsData) GetOrCreate(matchID int,MapNumber int,mapname string){

}*/

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

func (p *SQLPlayerStatsData) GetOrCreate() {

}

func (p *SQLPlayerStatsData) GetSteamURL() {
	return fmt.Sprintf("http://steamcommunity.com/profiles/%s", p.Steam_id)
}

/*func (p *SQLPlayerStatsData) GetRating() {
	AverageKPR = 0.679
	AverageSPR = 0.317
	AverageRMK = 1.277
	KillRating = float(self.kills) / float(self.roundsplayed) / AverageKPR
	SurvivalRating = float(self.roundsplayed -
						   self.deaths) / self.roundsplayed / AverageSPR
	killcount = float(self.k1 + 4 * self.k2 + 9 *
					  self.k3 + 16 * self.k4 + 25 * self.k5)
	RoundsWithMultipleKillsRating = killcount / \
		self.roundsplayed / AverageRMK
	rating = (KillRating + 0.7 * SurvivalRating +
			  RoundsWithMultipleKillsRating) / 2.7
	return rating
}*/
func (p *SQLPlayerStatsData) GetKDR() int {
	if p.Deaths == 0 {
		return p.Kills
	}
	return p.Kills / p.Deaths
}

func (p *SQLPlayerStatsData) GetHSP() float32 {
	if p.Deaths == 0 {
		return p.Kills
	}
	return float32(p.Headshot_kills / p.Kills)
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

/*
func GetSteamName(){

}
*/
