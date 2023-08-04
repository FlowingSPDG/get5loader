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

	Kills   int64
	Assists int64
	Deaths  int64

	RoundsPlayed     int64
	FlashbangAssists int64
	TeamKills        int64
	Suicides         int64
	HeadShotKills    int64
	Damage           int64
	BombPlants       int64
	BombDefuses      int64

	V1 int64
	V2 int64
	V3 int64
	V4 int64
	V5 int64
	K1 int64
	K2 int64
	K3 int64
	K4 int64
	K5 int64

	FirstDeathCT int64
	FirstDeathT  int64
	FirstKillCT  int64
	FirstKillT   int64
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
	SteamID string
	Name    string
	TeanID  int64
}
