package db

import (
	"bytes"
	"database/sql"
	"fmt"

	// "strings"
	//"github.com/Philipp15b/go-steam"
	//"github.com/FlowingSPDG/go-steamapi"
	"github.com/Acidic9/steam"
	"github.com/go-ini/ini"
	_ "github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
	"github.com/hydrogen18/stalecucumber"
	"github.com/jinzhu/gorm"

	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	gosteam "github.com/kidoman/go-steam"

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
	SQLUserData := UserData{}
	SQLUserData.SteamID = steamid

	record := g.Where("steam_id = ?", steamid).First(&SQLUserData)
	if record.RecordNotFound() {
		fmt.Println("USER NOT EXIST!")
		fmt.Println("CREATING USER")
		steamid64, err := strconv.Atoi(steamid)
		if err != nil {
			return u, err
		}
		SQLUserData.Name, err = GetSteamName(uint64(steamid64))
		if err != nil {
			return u, err
		}
		g.Create(&SQLUserData)
	} else {
		fmt.Println("USER EXIST")
		fmt.Println(SQLUserData)
		u.Name = SQLUserData.Name
		u.ID = SQLUserData.ID
		u.Admin = SQLUserData.Admin
		u.SteamID = SQLUserData.SteamID
	}
	return u, nil
}

func (u *UserData) GetURL() string {
	return fmt.Sprintf("http://%s/user/%d", Cnf.HOST, u.ID)
}

func (u *UserData) GetSteamURL() string {
	return "http://steamcommunity.com/profiles/" + u.SteamID
}

/*
func (u *UserData) get_recent_matches(limit int) string {
	return "http://steamcommunity.com/profiles/" + u.SteamID
}
*/

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

func (u *GameServerData) TableName() string {
	return "game_server"
}

func (g *GameServerData) Create(userid int, display_name string, ip_string string, port int, rcon_password string, public_server bool) *GameServerData {
	g.User_id = userid
	g.Display_name = display_name
	g.Ip_string = ip_string
	g.Port = port
	g.Rcon_password = rcon_password
	g.Public_server = public_server
	// ADD TO DB TODO
	return g
}

func (g *GameServerData) SendRcon(cmd string) (string, error) {
	o := &gosteam.ConnectOptions{RCONPassword: g.Rcon_password}
	rcon, err := gosteam.Connect(g.Ip_string, o)
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

func (g *GameServerData) GetDisplay() string {
	if g.Display_name == "" {
		return g.Display_name
	}
	return g.GetHostPort()
}

type TeamData struct {
	ID          int      `gorm:"primary_key"column:id`
	UserID      int      `gorm:"column:user_id"`
	Name        string   `gorm:"column:name"`
	Tag         string   `gorm:"column:tag"`
	Flag        string   `gorm:"column:flag"`
	Logo        string   `gorm:"column:logo"`
	AuthsPickle []byte   `gorm:"column:auths"`
	Auths       []string // converts pickle []byte to []string
	Players     []PlayerStatsData
	PublicTeam  bool `gorm:"column:public_team"`
}

func (u *TeamData) TableName() string {
	return "team"
}

func (t *TeamData) Create(userid int, name string, tag string, flag string, logo string, auths []byte, public_team bool) *TeamData {
	t.UserID = userid
	t.Name = name
	t.Tag = tag
	t.Flag = flag
	t.Logo = logo
	t.AuthsPickle = auths // should convert into pickle. TODO
	t.PublicTeam = public_team
	return t
}

func (t *TeamData) SetData(name string, tag string, flag string, logo string, auths []byte, public_team bool) *TeamData {
	t.Name = name
	t.Tag = tag
	t.Flag = flag
	t.Logo = logo
	t.AuthsPickle = auths // should convert into pickle. TODO
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

type PlayerIDName struct {
	Name string
	ID   string
}

func (t *TeamData) GetPlayers() ([]PlayerIDName, error) {
	reader := bytes.NewReader(t.AuthsPickle)
	var results []string
	var players = []PlayerIDName{}
	err := stalecucumber.UnpackInto(&results).From(stalecucumber.Unpickle(reader))
	if err != nil {
		return players, err
	}
	for i := 0; i < len(results); i++ {
		p := PlayerStatsData{}
		SQLAccess.Gorm.Where("steam_id = ?", results[i]).First(&p)
		fmt.Println(p)
		if err != nil {
			return players, err
		}
		player := PlayerIDName{
			ID:   results[i],
			Name: p.Name,
		}
		players = append(players, player)
	}
	return players, nil
}

/*
func (t *TeamData) CanDelete(userid int) bool {
	if t.CanEdit(userid) == false {
		return false
	}
	return len(t.GetRecentMatches()) == 0
}
*/

func (t *TeamData) GetRecentMatches(limit int) []MatchData {
	var matches []MatchData
	if t.PublicTeam == true {
		SQLAccess.Gorm.Where("team1_id = ? AND cancelled = false", t.ID).Or("team2_id = ? AND cancelled = false", t.ID).Not("start_time = null").Limit(limit).Find(&matches)
	} else {
		var owner UserData
		SQLAccess.Gorm.Where("id = ?", t.UserID).First(&owner)
		SQLAccess.Gorm.Where("user_id = ?", t.UserID).Find(&owner.Matches).Limit(limit)
		matches = owner.Matches
	}
	return matches
}

func (t *TeamData) GetVSMatchResult(matchid int) (string, error) {
	var otherteam TeamData
	myscore := 0
	otherteamscore := 0

	matches, err := SQLAccess.MySQLGetMatchData(1, "id", strconv.Itoa(matchid))
	if err != nil {
		return "", err
	}
	if int(matches[0].ID) == t.ID {
		myscore = matches[0].Team1Score
		otherteamscore = matches[0].Team2Score
		otherteams, err := SQLAccess.MySQLGetTeamData(1, "id", strconv.Itoa(int(matches[0].Team2ID)))
		if err != nil {
			return "", err
		}
		otherteam = otherteams[0]
	} else {
		myscore = matches[0].Team2Score
		otherteamscore = matches[0].Team1Score
		otherteams, err := SQLAccess.MySQLGetTeamData(1, "id", strconv.Itoa(int(matches[0].Team1ID)))
		if err != nil {
			return "", err
		}
		otherteam = otherteams[0]
	}

	// for a bo1 replace series score with the map score...
	if matches[0].MaxMaps == 1 {
		mapstats, err := matches[0].GetMapStat()
		if err != nil {
			return "", err
		}
		mapstat := mapstats[0]
		if int(matches[0].Team1ID) == t.ID {
			myscore = mapstat.Team1Score
			otherteamscore = mapstat.Team2Score
		} else {
			myscore = mapstat.Team2Score
			otherteamscore = mapstat.Team1Score
		}
	}
	if matches[0].Live() == true {
		return fmt.Sprintf("Live, %d:%d vs %s", myscore, otherteamscore, otherteam.Name), nil // maybe add <a> tag for otherteam.Name ?
	}
	if myscore < otherteamscore {
		return fmt.Sprintf("Lost %d:%d vs %s", myscore, otherteamscore, otherteam.Name), nil
	} else if myscore > otherteamscore {
		return fmt.Sprintf("Won %d:%d vs %s", otherteamscore, myscore, otherteam.Name), nil
	} else {
		return fmt.Sprintf("Tied %d:%d vs %s", otherteamscore, myscore, otherteam.Name), nil
	}

}

func (t *TeamData) GetFlagHTML(scale float32) string {
	if t.Flag == "" {
		return ""
	}
	width := 32.0 * scale
	height := 21.0 * scale
	html := fmt.Sprintf(`<img src="%s%s%s"  width="%f" height="%f">`, "../static/img/valve_flags/", t.Flag, ".png", width, height)
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

type MatchData struct {
	ID            int64 `gorm:"primary_key"column:id`
	UserID        int64 `gorm:"column:user_id"`
	User          UserData
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
		teams1core, team2score := m.GetCurrentScore(SQLAccess.Gorm)
		return fmt.Sprintf("Live, %d:%d", teams1core, team2score), nil
	} else if m.Finished() {
		teams1core, team2score := m.GetCurrentScore(SQLAccess.Gorm)
		minscore := math.Min(float64(teams1core), float64(team2score))
		maxscore := math.Max(float64(teams1core), float64(team2score))
		ScoreString := fmt.Sprintf("%d:%d", int(maxscore), int(minscore))
		winner, _ := m.Winner.Value()
		if !ShowWinner {
			return "Finished", nil
		} else if winner == m.Team1ID {
			team1, err := m.GetTeam1()
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("Won %s by %s", ScoreString, team1.Name), nil
		} else if winner == m.Team2ID {
			team2, err := m.GetTeam2()
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("Won %s by %s", ScoreString, team2.Name), nil
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
	team1score, team2score := m.GetCurrentScore(SQLAccess.Gorm)
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
	return m.EndTime.Valid && !m.Cancelled
}

func (m *MatchData) Live() bool {
	return m.StartTime.Valid && !m.EndTime.Valid && !m.Cancelled
}

func (m *MatchData) GetServer() int64 { // TODO : return server instance
	return m.ServerID // TODO
}

func (m *MatchData) GetCurrentScore(g *gorm.DB) (int, int) {
	//g.First(&m).Association("MapStats").Find(&m)
	m.MapStats = []MapStatsData{}
	g.First(&m.MapStats, "match_id = ?", m.ID)
	fmt.Println(m.MapStats)
	if m.MaxMaps == 1 {
		if len(m.MapStats) == 0 { // check ok?
			return 0, 0 // TODO
		}
		return m.MapStats[0].Team1Score, m.MapStats[0].Team2Score
	}
	return m.Team1Score, m.Team2Score
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
	Team.AuthsPickle = STeam[0].AuthsPickle
	Team.PublicTeam = STeam[0].PublicTeam
	reader := bytes.NewReader(STeam[0].AuthsPickle)
	Team.Auths = make([]string, 0)
	err = stalecucumber.UnpackInto(&Team.Auths).From(stalecucumber.Unpickle(reader))
	if err != nil {
		panic(err)
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
	Team.AuthsPickle = STeam[0].AuthsPickle
	Team.PublicTeam = STeam[0].PublicTeam
	reader := bytes.NewReader(STeam[0].AuthsPickle)
	Team.Auths = make([]string, 0)
	err = stalecucumber.UnpackInto(&Team.Auths).From(stalecucumber.Unpickle(reader))
	if err != nil {
		panic(err)
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

func (m *MatchData) GetMapStat() ([]MapStatsData, error) {
	var err error
	m.MapStats, err = SQLAccess.MySQLGetMapStatsData(7, "match_id", strconv.Itoa(int(m.ID)))
	if err != nil {
		return m.MapStats, err
	}
	return m.MapStats, nil
}

type MapStatsData struct {
	ID         int          `gorm:"primary_key" gorm:"column:id"`
	MatchID    int          `gorm:"column:match_id" gorm:"ForeignKey:match_id"`
	MapNumber  int          `gorm:"column:map_number"`
	MapName    string       `gorm:"column:map_name"`
	StartTime  sql.NullTime `gorm:"column:start_time"`
	EndTIme    sql.NullTime `gorm:"column:end_time"`
	Winner     int          `gorm:"column:winner"`
	Team1Score int          `gorm:"column:team1_score"`
	Team2Score int          `gorm:"column:team2_score"`
}

func (m *MapStatsData) TableName() string {
	return "map_stats"
}

/*func (m *MapStatsData) GetOrCreate(matchID int,MapNumber int,mapname string){

}*/

type PlayerStatsData struct {
	Id                int    `gorm:"primary_key" gorm:"column:id"`
	Match_id          int    `gorm:"column:match_id"`
	Map_id            int    `gorm:"column:map_id"`
	Team_id           int    `gorm:"column:team_id"`
	Steam_id          string `gorm:"column:steam_id"`
	Name              string `gorm:"column:name"`
	Kills             int    `gorm:"column:kills"`
	Deaths            int    `gorm:"column:deaths"`
	Roundsplayed      int    `gorm:"column:roundsplayed"`
	Assists           int    `gorm:"column:assists"`
	Flashbang_assists int    `gorm:"column:map_numbeflashbang_assists"`
	Teamkills         int    `gorm:"column:teamkills"`
	Suicides          int    `gorm:"column:suicides"`
	Headshot_kills    int    `gorm:"column:headshot_kills"`
	Damage            int64  `gorm:"column:damage"`
	Bomb_plants       int    `gorm:"column:bomb_plants"`
	Bomb_defuses      int    `gorm:"column:bomb_defuses"`
	V1                int    `gorm:"column:v1"`
	V2                int    `gorm:"column:v2"`
	V3                int    `gorm:"column:v3"`
	V4                int    `gorm:"column:v4"`
	V5                int    `gorm:"column:v5"`
	K1                int    `gorm:"column:k1"`
	K2                int    `gorm:"column:k2"`
	K3                int    `gorm:"column:k3"`
	K4                int    `gorm:"column:k4"`
	K5                int    `gorm:"column:k5"`
	Firstdeath_Ct     int    `gorm:"column:firstdeath_Ct"`
	Firstdeath_t      int    `gorm:"column:firstdeath_t"`
	Firstkill_ct      int    `gorm:"column:firstkill_ct"`
	Firstkill_t       int    `gorm:"column:firstkill_t"`
}

func (p *PlayerStatsData) TableName() string {
	return "player_stats"
}

/*
func (p *PlayerStatsData) GetOrCreate() string {
	return "player_stats"
}
*/

func (p *PlayerStatsData) GetSteamURL() string {
	return fmt.Sprintf("http://steamcommunity.com/profiles/%s", p.Steam_id)
}

func (p *PlayerStatsData) GetRating() float32 { // Rating value can be more accurate??
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

func (p *PlayerStatsData) GetKDR() int {
	if p.Deaths == 0 {
		return p.Kills
	}
	return p.Kills / p.Deaths
}

func (p *PlayerStatsData) GetHSP() float32 {
	if p.Deaths == 0 {
		return float32(p.Kills)
	}
	return float32(p.Headshot_kills / p.Kills)
}

type MatchesPageData struct {
	LoggedIn   bool
	UserName   string
	UserID     int
	Matches    []MatchData
	AllMatches bool
	MyMatches  bool
	Owner      UserData
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
	User       UserData
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

func GetSteamName(steamid uint64) (string, error) {
	summary, err := steam.GetPlayerSummaries(SteamAPIKey, steam.SteamID64(steamid))
	if err != nil {
		return "", err
	}
	return summary.DisplayName, nil
}
