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
	"github.com/hydrogen18/stalecucumber"
	"github.com/jinzhu/gorm"
	gosteam "github.com/kidoman/go-steam"

	//"github.com/solovev/steam_go"
	//_ "html/template"
	"math"
	"strconv"
)

func init() {
	c, _ := ini.Load("config.ini")
	Cnf = Config{
		HOST:      c.Section("GET5").Key("HOST").MustString("localhost:8080"),
		SQLHost:   c.Section("sql").Key("host").MustString(""),
		SQLUser:   c.Section("sql").Key("user").MustString(""),
		SQLPass:   c.Section("sql").Key("pass").MustString(""),
		SQLPort:   c.Section("sql").Key("port").MustInt(3306),
		SQLDBName: c.Section("sql").Key("database").MustString(""),
	}
	SteamAPIKey = c.Section("Steam").Key("APIKey").MustString("")
}

// UserData Struct for "user" table.
type UserData struct {
	ID      int    `gorm:"primary_key;column:id;AUTO_INCREMENT"`
	SteamID string `gorm:"column:steam_id;unique"`
	Name    string `gorm:"column:name"`
	Admin   bool   `gorm:"column:admin"`

	Servers []GameServerData `gorm:"foreignkey:user_id"`
	Teams   []TeamData       `gorm:"foreignkey:user_id"`
	Matches []MatchData      `gorm:"foreignkey:user_id"`
}

// TableName declairation for GORM
func (u *UserData) TableName() string {
	return "user"
}

// GetOrCreate Get or Register Userdata into DB.
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

// GetURL Get user page URL
func (u *UserData) GetURL() string {
	fmt.Println(Cnf)
	return fmt.Sprintf("http://%s/user/%d", Cnf.HOST, u.ID)
}

// GetSteamURL Get user's steam page URL by their steamid64
func (u *UserData) GetSteamURL() string {
	return "http://steamcommunity.com/profiles/" + u.SteamID
}

// GetRecentMatches Gets match history
func (u *UserData) GetRecentMatches(limit int) []MatchData {
	SQLAccess.Gorm.Where("user_id = ?", u.ID).Limit(limit).Find(&u.Matches)
	//SQLAccess.Gorm.Model(&u).Related(&u.Matches)
	return u.Matches
	//u.Matches
	//return self.matches.filter_by(cancelled=False).limit(limit)
}

// GetTeams Get teams which is owened by user
func (u *UserData) GetTeams(limit int) []TeamData {
	SQLAccess.Gorm.Where("user_id = ?", u.ID).Limit((limit)).Find(&u.Teams)
	return u.Teams
}

// GameServerData Struct for game_server table.
type GameServerData struct {
	ID           int    `gorm:"primary_key;column:id;AUTO_INCREMENT;NOT NULL"`
	UserID       int    `gorm:"column:user_id;DEFAULT NULL"`
	InUse        bool   `gorm:"column:in_use;DEFAULT NULL"`
	IPString     string `gorm:"column:ip_string;DEFAULT NULL"`
	Port         int    `gorm:"column:port;DEFAULT NULL"`
	RconPassword string `gorm:"column:rcon_password;DEFAULT NULL"`
	DisplayName  string `gorm:"column:display_name;DEFAULT NULL"`
	PublicServer bool   `gorm:"column:public_server;DEFAULT NULL"`

	User UserData `gorm:"ASSOCIATION_FOREIGNKEY:user_id"`
}

// TableName declairation for GORM
func (g *GameServerData) TableName() string {
	return "game_server"
}

// Create Register GameServer into DB. not implemented yet.
func (g *GameServerData) Create(userid int, displayname string, ipstring string, port int, rconpassword string, publicserver bool) *GameServerData {
	g.UserID = userid
	g.DisplayName = displayname
	g.IPString = ipstring
	g.Port = port
	g.RconPassword = rconpassword
	g.PublicServer = publicserver
	// ADD TO DB TODO
	return g
}

// SendRcon Sends Remote-Commands to server.
func (g *GameServerData) SendRcon(cmd string) (string, error) {
	o := &gosteam.ConnectOptions{RCONPassword: g.RconPassword}
	rcon, err := gosteam.Connect(g.IPString, o)
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

// GetHostPort gets gameserver addr and port. e.g. localhost:27015
func (g *GameServerData) GetHostPort() string {
	return fmt.Sprintf("%s:%d", g.IPString, g.Port)
}

// GetDisplay Returns "DisplayName" if its not empty. otherwise it returns address and port.
func (g *GameServerData) GetDisplay() string {
	if g.DisplayName == "" {
		return g.DisplayName
	}
	return g.GetHostPort()
}

// TeamData Struct for team table.
type TeamData struct {
	ID          int               `gorm:"primary_key;column:id"`
	UserID      int               `gorm:"column:user_id"`
	Name        string            `gorm:"column:name"`
	Tag         string            `gorm:"column:tag"`
	Flag        string            `gorm:"column:flag"`
	Logo        string            `gorm:"column:logo"`
	AuthsPickle []byte            `gorm:"column:auths"`
	Auths       []string          // converts pickle []byte to []string
	Players     []PlayerStatsData `gorm:"-"`
	PublicTeam  bool              `gorm:"column:public_team"`

	User UserData `gorm:"ASSOCIATION_FOREIGNKEY:user_id"`
}

// TableName declairation for GORM
func (t *TeamData) TableName() string {
	return "team"
}

// Create Register Team information into DB. not implemented.
func (t *TeamData) Create(userid int, name string, tag string, flag string, logo string, auths []byte, publicteam bool) *TeamData {
	t.UserID = userid
	t.Name = name
	t.Tag = tag
	t.Flag = flag
	t.Logo = logo
	t.AuthsPickle = auths // should convert into pickle. TODO
	t.PublicTeam = publicteam
	// should register into DB. TODO
	return t
}

// SetData Modify team data.
func (t *TeamData) SetData(name string, tag string, flag string, logo string, auths []byte, publicteam bool) *TeamData {
	t.Name = name
	t.Tag = tag
	t.Flag = flag
	t.Logo = logo
	t.AuthsPickle = auths // should convert into pickle. TODO
	t.PublicTeam = publicteam
	return t
}

// CanEdit Check if server is editable for user or not.
func (t *TeamData) CanEdit(userid int) bool {
	if userid == 0 {
		return false
	} else if t.UserID == userid {
		return true
	}
	return false
}

// CanDelete Check if server is deletable for user or not.
func (t *TeamData) CanDelete(userid int) bool {
	if t.CanEdit(userid) == false {
		return false
	}
	return len(t.GetRecentMatches(10)) == 0
}

// GetPlayerData Struct for GetPlayers() function.
type GetPlayerData struct {
	Auth string
	//Name string
}

// GetPlayers Gets registered player's steamid64.
func (t *TeamData) GetPlayers() ([]GetPlayerData, error) {
	reader := bytes.NewReader(t.AuthsPickle)
	var auths []string
	var results []GetPlayerData
	err := stalecucumber.UnpackInto(&auths).From(stalecucumber.Unpickle(reader))
	if err != nil {
		return results, err
	}
	for i := 0; i < len(auths); i++ {
		results = append(results, GetPlayerData{Auth: auths[i]})
	}
	/*for i := 0; i < len(auths); i++ {
		steamid64, err := strconv.Atoi(auths[i])
		if err != nil {
			return results, err
		}
		playername, err := GetSteamName(uint64(steamid64))
		if err != nil {
			return results, err
		}
		//results = append(results, GetPlayerData{Auth: auths[i], Name: playername})
		results = append(results, GetPlayerData{Auth: auths[i]})
	}*/

	return results, nil
}

/*
func (t *TeamData) GetPlayers() ([]PlayerStatsData, error) {
	reader := bytes.NewReader(t.AuthsPickle)
	var results []string
	err := stalecucumber.UnpackInto(&results).From(stalecucumber.Unpickle(reader))
	if err != nil {
		return t.Players, err
	}
	// it won't return datas in case team never played game...
		t.Players = []PlayerStatsData{}
		for i := 0; i < len(results); i++ {
			if results[i] != "" {
				p := PlayerStatsData{}
				SQLAccess.Gorm.Where("steam_id = ?", results[i]).First(&p)
				fmt.Println(p)
				if err != nil {
					return t.Players, err
				}
				if p.Steam_id != "" {
					t.Players = append(t.Players, p)
				}
			}
		}
	//SQLAccess.Gorm.Where("team_id = ?", t.ID).Find(&t.Players) // N+1 issue
	return t.Players, nil
}
*/

// GetRecentMatches Gets team match history.
func (t *TeamData) GetRecentMatches(limit int) []MatchData {
	var matches []MatchData
	if t.PublicTeam == true {
		SQLAccess.Gorm.Where("team1_id = ?", t.ID).Or("team2_id = ?", t.ID).Not("start_time = null AND cancelled = true").Limit(limit).Find(&matches)
	} else {
		var owner UserData
		SQLAccess.Gorm.Where("id = ?", t.UserID).First(&owner)
		SQLAccess.Gorm.Where("user_id = ?", t.UserID).Find(&owner.Matches).Limit(limit)
		matches = owner.Matches
	}
	return matches
}

// GetVSMatchResult Returns Match result as string for gorazor template.
func (t *TeamData) GetVSMatchResult(matchid int) (string, error) {
	var otherteam TeamData
	myscore := 0
	otherteamscore := 0

	matches, err := SQLAccess.MySQLGetMatchData(1, "id", strconv.Itoa(matchid))
	if err != nil {
		return "", err
	}
	if int(matches[0].Team1ID) == t.ID {
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
		if len(mapstats) <= 0 {
			return fmt.Sprintf("Pending, vs %s", otherteam.Name), nil // maybe add <a> tag for otherteam.Name ?
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

// GetFlagHTML Get team's flag as a HTML string. for gorazor template
func (t *TeamData) GetFlagHTML(scale float64) string {
	if t.Flag == "" {
		return ""
	}
	width := 32.0 * scale
	height := 21.0 * scale
	html := fmt.Sprintf(`<img src="%s%s%s"  width="%f" height="%f">`, "../static/img/valve_flags/", t.Flag, ".png", width, height)
	return html
}

// GetLogoHTML Get team's Logo as a HTML string. for gorazor template
func (t *TeamData) GetLogoHTML(scale float64) string {
	if t.Logo == "" {
		return ""
	}
	width := 32 * scale
	height := 32 * scale
	html := fmt.Sprintf(`<img src="%s"  width="%f" height="%f">`, t.Logo, width, height)
	return html
}

// GetURL Get URL of team page.
func (t *TeamData) GetURL() string {
	return fmt.Sprintf("http://%s/team/%d", Cnf.HOST, t.ID)
}

// GetNameURLHtml Get team page and name as a-tag. for gorazor template
func (t *TeamData) GetNameURLHtml() string {
	return fmt.Sprintf(`<a href="%s">%s</a>`, t.GetURL(), t.Name)
}

// GetLogoOrFlagHTML Get team logo or flag as a HTML.
func (t *TeamData) GetLogoOrFlagHTML(scale float64, otherteam TeamData) string {
	if t.Logo == "" && otherteam.Logo != "" { // or otherteam is empty...
		return t.GetLogoHTML(scale)
	}
	return t.GetFlagHTML(scale)
}

// MatchData Struct for match table.
type MatchData struct {
	ID            int64         `gorm:"primary_key;column:id"`
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

	User UserData `gorm:"ASSOCIATION_FOREIGNKEY:user_id"`
}

// TableName declairation for GORM
func (m *MatchData) TableName() string {
	return "match"
}

// Create Register Match information into DB. not implemented yet
func (m *MatchData) Create(userid int64, team1id int64, team2id int64, team1string string, team2string string, maxmaps int, skipveto bool, title string, vetomappool string, serverid int64) *MatchData {
	m.UserID = userid
	m.Team1ID = team1id
	m.Team2ID = team2id
	m.SkipVeto = skipveto
	m.Title = title
	m.VetoMapPool = vetomappool // array...?
	m.ServerID = serverid
	m.MaxMaps = maxmaps
	m.APIKey = "" //random?
	return m      // TODO
}

// GetStatusString Get match status as string. for gorazor template
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

// GetVSString Get Match VS information as string. for gorazor template
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

// Finalized Returns true if match is finished or cancelled
func (m *MatchData) Finalized() bool {
	return m.Cancelled || m.Finished()
}

// Pending Returns true if match is not started and not cancelled
func (m *MatchData) Pending() bool {
	return !m.StartTime.Valid && !m.Cancelled
}

// Finished Returns true if match is ended and not cancelled
func (m *MatchData) Finished() bool {
	return m.EndTime.Valid && !m.Cancelled
}

// Live Retursn true if match is in-progress
func (m *MatchData) Live() bool {
	return m.StartTime.Valid && !m.EndTime.Valid && !m.Cancelled
}

// GetServer Get match server ID as GameServerData
func (m *MatchData) GetServer() GameServerData {
	SQLAccess.Gorm.Where("user_id = ?", m.UserID).First(&m.Server)
	return m.Server
}

// GetCurrentScore Returns current match score. returns map-score if match is BO1.
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

// SendToServer Let gameserver load match configration via RCON. not implemented yet
/*func (m *MatchData) SendToServer() {
	// get5_loadmatch_url things
}*/

// GetTeam1 Get Team1 as "TeamData" struct.
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
		return Team, err
	}
	return Team, nil
}

// GetTeam2 Get Team2 as "TeamData" struct.
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
		return Team, err
	}
	return Team, nil
}

// GetUser Get Match owner as "UserData" struct.
func (m *MatchData) GetUser() UserData {
	SQLAccess.Gorm.Where("id = ?", m.UserID).First(&m.User)
	return m.User
}

// GetWinner Get Winner team as "TeamData" struct.
func (m *MatchData) GetWinner() (TeamData, error) {
	var Empty TeamData
	if m.Team1Score > m.Team2Score {
		winner, err := m.GetTeam1()
		if err != nil {
			return Empty, err
		}
		return winner, nil
	} else if m.Team2Score > m.Team1Score {
		winner, err := m.GetTeam2()
		if err != nil {
			return Empty, err
		}
		return winner, nil
	}
	return Empty, nil
}

// GetLoser Get Loser team as "TeamData" struct.
func (m *MatchData) GetLoser() (TeamData, error) {
	var Empty TeamData
	if m.Team1Score > m.Team2Score {
		loser, err := m.GetTeam2()
		if err != nil {
			return Empty, err
		}
		return loser, nil
	} else if m.Team2Score > m.Team1Score {
		loser, err := m.GetTeam1()
		if err != nil {
			return Empty, err
		}
		return loser, nil
	}
	return Empty, nil
}

// BuildMatchDict No idea.
/*func (m *MatchData) BuildMatchDict() TeamData {
	//return m.UserID //get5 thing??
}*/

// GetMapStat Gets each map stat data as "MapStatsData" struct array.
func (m *MatchData) GetMapStat() ([]MapStatsData, error) {
	var err error
	m.MapStats, err = SQLAccess.MySQLGetMapStatsData(7, "match_id", strconv.Itoa(int(m.ID)))
	if err != nil {
		return m.MapStats, err
	}
	return m.MapStats, nil
}

// MapStatsData MapStatsData struct for map_stats table.
type MapStatsData struct {
	ID         int          `gorm:"primary_key" gorm:"column:id"`
	MatchID    int          `gorm:"column:match_id" gorm:"ForeignKey:match_id"`
	MapNumber  int          `gorm:"column:map_number"`
	MapName    string       `gorm:"column:map_name"`
	StartTime  sql.NullTime `gorm:"column:start_time"`
	EndTime    sql.NullTime `gorm:"column:end_time"`
	Winner     int          `gorm:"column:winner"`
	Team1Score int          `gorm:"column:team1_score"`
	Team2Score int          `gorm:"column:team2_score"`

	User UserData `gorm:"ASSOCIATION_FOREIGNKEY:user_id"`
}

// TableName declairation for GORM
func (m *MapStatsData) TableName() string {
	return "map_stats"
}

// GetOrCreate Get or register mapstats data. not implemented yet
/*func (m *MapStatsData) GetOrCreate(matchID int,MapNumber int,mapname string){

}*/

// PlayerStatsData Player stats data struct for player_stats table.
type PlayerStatsData struct {
	ID               int    `gorm:"primary_key;column:id"`
	MatchID          int    `gorm:"column:match_id"`
	MapID            int    `gorm:"column:map_id"`
	TeamID           int    `gorm:"column:team_id"`
	SteamID          string `gorm:"column:steam_id;unique"`
	Name             string `gorm:"column:name"`
	Kills            int    `gorm:"column:kills"`
	Deaths           int    `gorm:"column:deaths"`
	Roundsplayed     int    `gorm:"column:roundsplayed"`
	Assists          int    `gorm:"column:assists"`
	FlashbangAssists int    `gorm:"column:flashbang_assists"`
	Teamkills        int    `gorm:"column:teamkills"`
	Suicides         int    `gorm:"column:suicides"`
	HeadshotKills    int    `gorm:"column:headshot_kills"`
	Damage           int64  `gorm:"column:damage"`
	BombPlants       int    `gorm:"column:bomb_plants"`
	BombDefuses      int    `gorm:"column:bomb_defuses"`
	V1               int    `gorm:"column:v1"`
	V2               int    `gorm:"column:v2"`
	V3               int    `gorm:"column:v3"`
	V4               int    `gorm:"column:v4"`
	V5               int    `gorm:"column:v5"`
	K1               int    `gorm:"column:k1"`
	K2               int    `gorm:"column:k2"`
	K3               int    `gorm:"column:k3"`
	K4               int    `gorm:"column:k4"`
	K5               int    `gorm:"column:k5"`
	FirstdeathCT     int    `gorm:"column:firstdeath_Ct"`
	FirstdeathT      int    `gorm:"column:firstdeath_t"`
	FirstkillCT      int    `gorm:"column:firstkill_ct"`
	FirstkillT       int    `gorm:"column:firstkill_t"`

	User UserData `gorm:"ASSOCIATION_FOREIGNKEY:user_id"`
}

// TableName declairation for GORM
func (p *PlayerStatsData) TableName() string {
	return "player_stats"
}

// GetOrCreate Get or register player stats data into DB.
/*
func (p *PlayerStatsData) GetOrCreate() string {
	return "player_stats"
}
*/

// GetSteamURL get player's Steam community URL by their steamid64.
func (p *PlayerStatsData) GetSteamURL() string {
	return fmt.Sprintf("http://steamcommunity.com/profiles/%s", p.SteamID)
}

// GetRating Get player's rating. Average datas are static tho.
func (p *PlayerStatsData) GetRating() float64 { // Rating value can be more accurate??
	var AverageKPR float64 = 0.679
	var AverageSPR float64 = 0.317
	var AverageRMK float64 = 1.277
	var KillRating float64 = float64(p.Kills) / float64(p.Roundsplayed) / AverageKPR
	var SurvivalRating float64 = float64(p.Roundsplayed-p.Deaths) / float64(p.Roundsplayed) / float64(AverageSPR)
	var killcount float64 = float64(p.K1 + 4*p.K2 + 9*p.K3 + 16*p.K4 + 25*p.K5)
	var RoundsWithMultipleKillsRating float64 = killcount / float64(p.Roundsplayed) / float64(AverageRMK)
	var rating float64 = (KillRating + 0.7*SurvivalRating + RoundsWithMultipleKillsRating) / 2.7
	return rating
}

// GetKDR Returns player's KDR(Kill/Deaths Ratio).
func (p *PlayerStatsData) GetKDR() float64 {
	if p.Deaths == 0 {
		return float64(p.Kills)
	}
	return float64(p.Kills) / float64(p.Deaths)
}

// GetHSP Returns player's HSP(HeadShot Percentage).
func (p *PlayerStatsData) GetHSP() float64 {
	if p.Deaths == 0 {
		return float64(p.Kills)
	}
	return float64(p.HeadshotKills) / float64(p.Kills) * 100
}

// GetADR Returns player's ADR(Average Damage per Round).
func (p *PlayerStatsData) GetADR() float64 {
	if p.Roundsplayed == 0 {
		return 0.0
	}
	return float64(p.Damage) / float64(p.Roundsplayed)
}

// GetFPR Returns player's FPR(Frags Per Round).
func (p *PlayerStatsData) GetFPR() float64 {
	if p.Roundsplayed == 0 {
		return 0.0
	}
	return float64(p.Kills) / float64(p.Roundsplayed)
}

// GetSteamName Get steam profile name by steamid64 via Steam web API
func GetSteamName(steamid uint64) (string, error) {
	summary, err := steam.GetPlayerSummaries(SteamAPIKey, steam.SteamID64(steamid))
	if err != nil {
		return "", err
	}
	return summary.DisplayName, nil
}

// MetricsData Struct metrics analysys.
type MetricsData struct {
	RegisteredUsers    int
	SavedTeams         int
	MatchesCreated     int
	CompletedMatches   int
	ServersAdded       int
	MapsWithStatsSaved int
	UniquePlayers      int
}

// GetMetrics Get Each table's count.
func GetMetrics() MetricsData {
	var result MetricsData

	SQLAccess.Gorm.Table("user").Count(&result.RegisteredUsers)
	SQLAccess.Gorm.Table("team").Count(&result.SavedTeams)
	SQLAccess.Gorm.Table("match").Count(&result.MatchesCreated)
	SQLAccess.Gorm.Table("match").Where("end_time IS NOT NULL").Count(&result.CompletedMatches)
	SQLAccess.Gorm.Table("game_server").Count(&result.ServersAdded)
	SQLAccess.Gorm.Table("map_stats").Count(&result.MapsWithStatsSaved)
	SQLAccess.Gorm.Table("player_stats").Count(&result.UniquePlayers)

	return result
}
