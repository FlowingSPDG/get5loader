package db

import (
	"database/sql"
	"fmt"
	// "strings"
	//"github.com/Philipp15b/go-steam"
	_ "github.com/FlowingSPDG/go-steamapi"
	"github.com/go-ini/ini"
	_ "github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	steam "github.com/kidoman/go-steam"
	//"github.com/solovev/steam_go"
	//_ "html/template"
	"math"
	_ "net/http"
	"strconv"
)

func init() {
	c, _ := ini.Load("config.ini")
	SteamAPIKey = c.Section("Steam").Key("APIKey").MustString("")
}

type UserData struct {
	ID      int    `gorm:"primary_key"column:id`
	SteamID string `gorm:"column:steam_id"`
	Name    string `gorm:"column:name"`
	Admin   bool   `gorm:"column:admin"`
	Servers []GameServerData
	Teams   []TeamData
	Matches []MatchData
}

func (u *UserData) TableName() string {
	return "user"
}

func (u *UserData) GetOrCreate(g *gorm.DB, steamid string) (*UserData, error) {
	var User *UserData
	SQLUserData := UserData{}
	SQLUserData.SteamID = steamid

	if g.Find(&SQLUserData).RecordNotFound() {
		fmt.Println("USER NOT EXIST")
	}
	fmt.Println("USER EXIST?")
	User.Name = SQLUserData.Name
	User.ID = SQLUserData.ID
	User.Admin = SQLUserData.Admin
	User.SteamID = SQLUserData.SteamID
	return User, nil
	/*
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
	*/
}

func (u *UserData) GetURL() string {
	return fmt.Sprintf("http://%s/user/%d", Cnf.HOST, u.ID)
}

func (u *UserData) GetSteamURL() string {
	return "http://steamcommunity.com/profiles/" + u.SteamID
}

func (u *UserData) get_recent_matches(limit int) string {
	return "http://steamcommunity.com/profiles/" + u.SteamID
}

type GameServerData struct {
	Id            int    `gorm:"primary_key"column:id`
	User_id       int    `gorm:"column:user_id"`
	In_use        bool   `gorm:"column:in_use"`
	Ip_string     string `gorm:"column:ip_string"`
	Port          int    `gorm:"column:port"`
	Rcon_password string `gorm:"column:rcon_password"`
	Display_name  string `gorm:"column:display_name"`
	Public_server bool   `gorm:"column:public_server"`
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
	ID     int    `gorm:"primary_key"column:id`
	UserID int    `gorm:"column:user_id"`
	Name   string `gorm:"column:name"`
	Tag    string `gorm:"column:tag"`
	Flag   string `gorm:"column:flag"`
	Logo   string `gorm:"column:logo"`
	//Auths      []string
	Auths      []byte `gorm:"column:auth"`
	PublicTeam bool   `gorm:"column:public_team"`
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

/*
func (t *TeamData) CanDelete(userid int) bool {
	if t.CanEdit(userid) == false {
		return false
	}
	return len(t.GetRecentMatches()) == 0
}
*/

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

func (t *TeamData) GetLogoHtml(scale float32) string {
	if t.Logo == "" {
		return ""
	}
	width := 32 * scale
	height := 32 * scale
	html := fmt.Sprintf(`<img src="%s"  width="%f" height="%f">`, t.Logo, width, height)
	return html
}

func (t *TeamData) GetURL() string {
	return fmt.Sprintf("http://%s/team/%d", Cnf.HOST, t.ID)
}

func (t *TeamData) GetNameURLHtml() string {
	return fmt.Sprintf(`<a href="%s">%s</a>`, t.GetURL(), t.Name)
}

func (t *TeamData) GetLogoOrFlagHtml(scale float32, otherteam TeamData) string {
	if t.Logo == "" && otherteam.Logo != "" { // or otherteam is empty...
		return t.GetLogoHtml(scale)
	}
	return t.GetFlagHTML(scale)
}

func (m *TeamData) TableName() string {
	return "team"
}

type MatchData struct {
	ID            int64         `gorm:"primary_key"column:id`
	UserID        int64         `gorm:"column:user_id"`
	ServerID      int64         `gorm:"column:server_id"`
	Team1ID       int64         `gorm:"column:team1_id"`
	Team2ID       int64         `gorm:"column:team2_id"`
	Winner        sql.NullInt64 `gorm:"column:winner"`
	Cancelled     bool          `gorm:"column:cancelled"`
	StartTime     sql.NullTime  `gorm:"column:start_time"`
	EndTime       sql.NullTime  `gorm:"column:end_time"`
	MaxMaps       int           `gorm:"column:max_maps"`
	Title         string        `gorm:"column:title"`
	SkipVeto      bool          `gorm:"column:skip_veto"`
	APIKey        string        `gorm:"column:api_key"`
	VetoMapPool   string        `gorm:"column:veto_mappool"`
	Team1Score    int           `gorm:"column:team1_score"`
	Team2Score    int           `gorm:"column:team2_score"`
	Team1String   string        `gorm:"column:team1_string"`
	Team2String   string        `gorm:"column:team2_string"`
	Forfeit       bool          `gorm:"column:forfeit"`
	PluginVersion string        `gorm:"column:plugin_version"`

	MapStats []MapStatsData
	Server   GameServerData
}

func (u *MatchData) TableName() string {
	return "match"
}

func (m *MatchData) Create(userid int64, team1_id int64, team2_id int64, team1_string string, team2_string string, max_maps int, skip_veto bool, title string, veto_mappool string, server_id int64) *MatchData {
	m.UserID = userid
	m.Team1ID = team1_id
	m.Team2ID = team2_id
	m.SkipVeto = skip_veto
	m.Title = title
	m.VetoMapPool = veto_mappool // array...?
	m.ServerID = server_id
	m.MaxMaps = max_maps
	m.APIKey = "" //random?
	return m
}

func (m *MatchData) GetStatusString(ShowWinner bool) (string, error) {
	if m.Pending() {
		return "Pending", nil
	} else if m.Live() {
		teams1core, team2score := m.GetCurrentScore()
		return fmt.Sprintf("Live, %d:%d", teams1core, team2score), nil
	} else if m.Finished() {
		teams1core, team2score := m.GetCurrentScore()
		minscore := math.Min(float64(teams1core), float64(team2score))
		maxscore := math.Max(float64(teams1core), float64(team2score))
		ScoreString := fmt.Sprintf("%f:%f", maxscore, minscore)
		winner, _ := m.Winner.Value()
		if !ShowWinner {
			return "Finished", nil
		} else if winner == m.Team1ID {
			team1, err := m.GetTeam1()
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("Won %s by %f", ScoreString, team1.Name), nil
		} else if winner == m.Team2ID {
			team2, err := m.GetTeam2()
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("Won %s by %f", ScoreString, team2.Name), nil
		} else {
			return fmt.Sprintf("Tied %s", ScoreString), nil
		}
	} else {
		return "Cancelled", nil
	}
}

func (m *MatchData) GetVSString() (string, error) {
	team1, err := m.GetTeam1()
	team2, err := m.GetTeam2()
	if err != nil {
		return "", err
	}
	team1score, team2score := m.GetCurrentScore()
	str := fmt.Sprintf("%s VS %s (%d:%d)", team1.GetNameURLHtml(), team2.GetNameURLHtml(), team1score, team2score)
	return str, nil
}

func (m *MatchData) Finalized() bool {
	return m.Cancelled || m.Finished()
}

func (m *MatchData) Pending() bool {
	return !m.StartTime.Valid && !m.Cancelled
}

func (m *MatchData) Finished() bool {
	return !m.EndTime.Valid && !m.Cancelled
}

func (m *MatchData) Live() bool {
	return m.StartTime.Valid && !m.EndTime.Valid && !m.Cancelled
}

func (m *MatchData) GetServer() int64 {
	return m.ServerID // TODO
}

func (m *MatchData) GetCurrentScore() (int, int) {
	if m.MaxMaps == 1 {
		if len(m.MapStats) == 0 { // check ok?
			return 0, 0
		}
		return m.MapStats[0].Team1Score, m.MapStats[0].Team2Score
	} else {
		return m.Team1Score, m.Team2Score
	}
}

/*func (m *MatchData) SendToServer() {
	// get5_loadmatch_url things
}*/

func (m *MatchData) GetTeam1() (TeamData, error) {
	var Team = TeamData{}
	STeam, err := SQLAccess.MySQLGetTeamData(1, "id", strconv.Itoa(int(m.Team1ID)))
	Team.ID = STeam[0].ID
	Team.Name = STeam[0].Name
	Team.Tag = STeam[0].Tag
	Team.Flag = STeam[0].Flag
	Team.Logo = STeam[0].Logo
	Team.Auths = STeam[0].Auths
	Team.PublicTeam = STeam[0].PublicTeam
	if err != nil {
		return Team, err
	}
	return Team, nil
}

func (m *MatchData) GetTeam2() (TeamData, error) {
	var Team = TeamData{}
	STeam, err := SQLAccess.MySQLGetTeamData(1, "id", strconv.Itoa(int(m.Team2ID)))
	Team.ID = STeam[0].ID
	Team.Name = STeam[0].Name
	Team.Tag = STeam[0].Tag
	Team.Flag = STeam[0].Flag
	Team.Logo = STeam[0].Logo
	Team.Auths = STeam[0].Auths
	Team.PublicTeam = STeam[0].PublicTeam
	if err != nil {
		return Team, err
	}
	return Team, nil
}

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

type MapStatsData struct {
	ID         int          `gorm:"primary_key"column:id`
	MatchID    int          `gorm:"column:match_id"`
	MapNumber  int          `gorm:"column:map_number"`
	MapName    string       `gorm:"column:map_name"`
	StartTime  sql.NullTime `gorm:"column:start_time"`
	EndTIme    sql.NullTime `gorm:"column:end_time"`
	Winner     int          `gorm:"column:winner"`
	Team1Score int          `gorm:"column:team1_score"`
	Team2Score int          `gorm:"column:team2_score"`
}

/*func (m *MapStatsData) GetOrCreate(matchID int,MapNumber int,mapname string){

}*/

type SQLMapStatsData struct {
	Id          int `gorm:"primary_key"`
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
	Id                int `gorm:"primary_key"`
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

func (p *SQLPlayerStatsData) GetSteamURL() string {
	return fmt.Sprintf("http://steamcommunity.com/profiles/%s", p.Steam_id)
}

func (p *SQLPlayerStatsData) GetRating() float32 { // Rating value can be more accurate??
	var AverageKPR float32 = 0.679
	var AverageSPR float32 = 0.317
	var AverageRMK float32 = 1.277
	var KillRating float32 = float32(p.Kills) / float32(p.Roundsplayed) / AverageKPR
	var SurvivalRating float32 = float32(p.Roundsplayed-p.Deaths) / float32(p.Roundsplayed) / float32(AverageSPR)
	var killcount float32 = float32(p.K1 + 4*p.K2 + 9*p.K3 + 16*p.K4 + 25*p.K5)
	var RoundsWithMultipleKillsRating float32 = killcount / float32(p.Roundsplayed) / float32(AverageRMK)
	var rating float32 = (KillRating + 0.7*SurvivalRating + RoundsWithMultipleKillsRating) / 2.7
	return rating
}

func (p *SQLPlayerStatsData) GetKDR() int {
	if p.Deaths == 0 {
		return p.Kills
	}
	return p.Kills / p.Deaths
}

func (p *SQLPlayerStatsData) GetHSP() float32 {
	if p.Deaths == 0 {
		return float32(p.Kills)
	}
	return float32(p.Headshot_kills / p.Kills)
}

type MatchesPageData struct {
	LoggedIn bool
	UserName string
	UserID   int
	Matches  []MatchData
}

type TeamsPageData struct {
	LoggedIn   bool
	User       UserData
	IsYourTeam bool
	Teams      []TeamData
}

type TeamPageData struct {
	LoggedIn   bool
	IsYourTeam bool
	Team       TeamData
}

type UserPageData struct {
	LoggedIn bool
	User     UserData
	Teams    []TeamData
}

type MyserversPageData struct {
	Servers  []GameServerData
	LoggedIn bool
}

/*
func GetSteamName(){

}
*/
