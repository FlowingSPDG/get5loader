package entity

import "time"

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
	Port         int32
	RCONPassword string
	InUse        bool
	DisplayName  string
	IsPublic     bool
}

type Match struct {
	ID         int64
	UserID     int64
	ServerID   int64
	Team1ID    int64
	Team2ID    int64
	Winner     *int64
	Cancelled  bool
	StartTime  *time.Time
	EndTime    *time.Time
	MaxMaps    int32
	Title      string
	SkipVeto   bool
	APIKey     string
	Team1Score int32
	Team2Score int32
	Forfeit    *bool
}

type MapStats struct {
	ID         int64
	MatchID    int64
	MapNumber  int32
	MapName    string
	StartTime  *time.Time
	EndTime    *time.Time
	Winner     *int64
	Team1Score int32
	Team2Score int32
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

	RoundsPlayed     int32
	FlashbangAssists int32
	Suicides         int32
	HeadShotKills    int32
	Damage           int32
	BombPlants       int32
	BombDefuses      int32

	V1 int32
	V2 int32
	V3 int32
	V4 int32
	V5 int32
	K1 int32
	K2 int32
	K3 int32
	K4 int32
	K5 int32

	FirstDeathCT int32
	FirstDeathT  int32
	FirstKillCT  int32
	FirstKillT   int32
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
