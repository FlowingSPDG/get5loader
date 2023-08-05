package entity

import "time"

type SERVER_STATUS int

const (
	SERVER_STATUS_UNKNOWN SERVER_STATUS = iota
	SERVER_STATUS_STANDBY
	SERVER_STATUS_INUSE
)

type MATCH_STATUS int

const (
	MATCH_STATUS_UNKNOWN MATCH_STATUS = iota
	MATCH_STATUS_PENDING
	MATCH_STATUS_LIVE
	MATCH_STATUS_FINISHED
	MATCH_STATUS_CANCELLED
)

type User struct {
	ID      int64
	SteamID string
	Name    string
	Admin   bool
}

type GameServer struct {
	UserID       int64
	ID           int64
	Ip           string
	Port         uint32
	RCONPassword string
	DisplayName  string
	IsPublic     bool
	Status       SERVER_STATUS
}

type Match struct {
	ID         int64
	UserID     int64
	ServerID   int64
	Team1ID    int64
	Team2ID    int64
	Winner     *int64
	StartTime  *time.Time
	EndTime    *time.Time
	MaxMaps    int32
	Title      string
	SkipVeto   bool
	APIKey     string
	Team1Score uint32
	Team2Score uint32
	Forfeit    *bool
	Status     MATCH_STATUS
}

type MapStats struct {
	ID         int64
	MatchID    int64
	MapNumber  uint32
	MapName    string
	StartTime  *time.Time
	EndTime    *time.Time
	Winner     *int64
	Team1Score uint32
	Team2Score uint32
}

type PlayerStats struct {
	ID      int64
	MatchID int64
	MapID   int64
	TeamID  int64
	SteamID string
	Name    string

	Kills   int32
	Assists int32
	Deaths  int32

	RoundsPlayed     uint32
	FlashbangAssists uint32
	Suicides         uint32
	HeadShotKills    uint32
	Damage           uint32
	BombPlants       uint32
	BombDefuses      uint32

	V1 uint32
	V2 uint32
	V3 uint32
	V4 uint32
	V5 uint32
	K1 uint32
	K2 uint32
	K3 uint32
	K4 uint32
	K5 uint32

	FirstDeathCT uint32
	FirstDeathT  uint32
	FirstKillCT  uint32
	FirstKillT   uint32
}

type Team struct {
	ID         int64
	UserID     int64
	Name       string
	Flag       string
	Tag        string
	Logo       string
	PublicTeam bool
}

type Player struct {
	ID      int64
	TeamID  int64
	SteamID string
	Name    string
	TeanID  int64
}
