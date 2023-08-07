package entity

import "time"

type (
	UserID        string
	GameServerID  string
	MatchID       string
	TeamID        string
	MapStatsID    string
	PlayerID      string
	PlayerStatsID string
	SteamID       uint64 // SteamID64. Note: SteamID3 is not supported. Note: some database drivers may not support uint64.
)

type User struct {
	ID        UserID
	SteamID   SteamID
	Name      string
	Admin     bool
	Hash      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GameServer struct {
	UserID       UserID
	ID           GameServerID
	Ip           string
	Port         uint32
	RCONPassword string
	DisplayName  string
	IsPublic     bool
	Status       SERVER_STATUS
}

type Match struct {
	ID         MatchID
	UserID     UserID
	ServerID   GameServerID
	Team1ID    TeamID
	Team2ID    TeamID
	Winner     *TeamID
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
	ID         MapStatsID
	MatchID    MatchID
	MapNumber  uint32
	MapName    string
	StartTime  *time.Time
	EndTime    *time.Time
	Winner     *TeamID
	Team1Score uint32
	Team2Score uint32
}

type PlayerStats struct {
	ID      PlayerStatsID
	MatchID MatchID
	MapID   MapStatsID
	TeamID  TeamID
	SteamID SteamID
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
	ID         TeamID
	UserID     UserID
	Name       string
	Flag       string
	Tag        string
	Logo       string
	PublicTeam bool
}

type Player struct {
	ID      PlayerID
	TeamID  TeamID
	SteamID SteamID
	Name    string
}
