package api

import (
	"database/sql"
	"time"
)

// UserData Struct for "user" table.
type APIUserData struct {
	ID      int    `gorm:"primary_key;column:id;AUTO_INCREMENT" json:"id"`
	SteamID string `gorm:"column:steam_id;unique" json:"steam_id"`
	Name    string `gorm:"column:name" json:"name"`
	Admin   bool   `gorm:"column:admin" json:"admin"`

	Servers []APIGameServerData `gorm:"foreignkey:user_id" json:"servers"`
	Teams   []APITeamData       `gorm:"foreignkey:user_id" json:"teams"`
	Matches []APIMatchData      `gorm:"foreignkey:user_id" json:"matches"`
}

// TableName declairation for GORM
func (u *APIUserData) TableName() string {
	return "user"
}

// GameServerData Struct for game_server table.
type APIGameServerData struct {
	ID           int    `gorm:"primary_key;column:id;AUTO_INCREMENT;NOT NULL" json:"id"`
	UserID       int    `gorm:"column:user_id;DEFAULT NULL" json:"user_id"`
	InUse        bool   `gorm:"column:in_use;DEFAULT NULL" json:"in_use"`
	IPString     string `gorm:"column:ip_string;DEFAULT NULL" json:"ip_string"`
	Port         int    `gorm:"column:port;DEFAULT NULL" json:"port"`
	Display      string `json:"display"`
	PublicServer bool   `gorm:"column:public_server;DEFAULT NULL" json:"public_server"`
}

// TableName declairation for GORM
func (u *APIGameServerData) TableName() string {
	return "game_server"
}

// TeamData Struct for team table.
type APITeamData struct {
	ID          int                  `gorm:"primary_key;column:id" json:"id"`
	UserID      int                  `gorm:"column:user_id" json:"user_id"`
	Name        string               `gorm:"column:name" json:"name"`
	Tag         string               `gorm:"column:tag" json:"tag"`
	Flag        string               `gorm:"column:flag" json:"flag"`
	Logo        string               `gorm:"column:logo" json:"logo"`
	AuthsPickle []byte               `gorm:"column:auths" json:"-"`
	Auths       []string             `json:"auths"` // converts pickle []byte to []string
	Players     []APIPlayerStatsData `gorm:"-" json:"-"`
	PublicTeam  bool                 `gorm:"column:public_team" json:"public_team"`

	User APIUserData `gorm:"ASSOCIATION_FOREIGNKEY:user_id" json:"-"`
}

// TableName declairation for GORM
func (u *APITeamData) TableName() string {
	return "team"
}

// MatchData Struct for match table.
type APIMatchData struct {
	ID          int64       `json:"id"`
	UserID      int64       `json:"user_id"`
	Team1       APITeamData `json:"team1"`
	Team2       APITeamData `json:"team2"`
	Winner      int64       `json:"winner"`
	Cancelled   bool        `json:"cancelled"`
	StartTime   time.Time   `json:"start_time"`
	EndTime     time.Time   `json:"end_time"`
	MaxMaps     int         `json:"max_maps"`
	Title       string      `json:"title"`
	SkipVeto    bool        `json:"skip_veto"`
	VetoMapPool []string    `json:"veto_mappool"`
	Team1Score  int         `json:"team1_score"`
	Team2Score  int         `json:"team2_score"`
	Team1String string      `json:"team1_string"`
	Team2String string      `json:"team2_string"`
	Forfeit     bool        `json:"forfeit"`

	MapStats []APIMapStatsData `json:"map_stats"`
	Server   APIGameServerData `json:"server"`
	User     APIUserData       `json:"user"`

	Pending bool   `json:"pending"`
	Live    bool   `json:"live"`
	Status  string `json:"status"`
}

// TableName declairation for GORM
func (u *APIMatchData) TableName() string {
	return "match"
}

// MapStatsData MapStatsData struct for map_stats table.
type APIMapStatsData struct {
	ID         int          `gorm:"primary_key" gorm:"column:id" json:"id"`
	MatchID    int          `gorm:"column:match_id" gorm:"ForeignKey:match_id" json:"match_id"`
	MapNumber  int          `gorm:"column:map_number" json:"map_number"`
	MapName    string       `gorm:"column:map_name" json:"map_name"`
	StartTime  sql.NullTime `gorm:"column:start_time" json:"start_time"`
	EndTime    sql.NullTime `gorm:"column:end_time" json:"end_time"`
	Winner     int          `gorm:"column:winner" json:"winner"`
	Team1Score int          `gorm:"column:team1_score" json:"team1_score"`
	Team2Score int          `gorm:"column:team2_score" json:"team2_score"`
}

// TableName declairation for GORM
func (u *APIMapStatsData) TableName() string {
	return "map_stats"
}

// PlayerStatsData Player stats data struct for player_stats table.
type APIPlayerStatsData struct {
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

	User APIUserData `gorm:"ASSOCIATION_FOREIGNKEY:user_id"`
}

// TableName declairation for GORM
func (u *APIPlayerStatsData) TableName() string {
	return "game_server"
}
