package database

import (
	"time"

	"github.com/FlowingSPDG/get5loader/backend/entity"
)

type User struct {
	ID        entity.UserID
	SteamID   entity.SteamID
	Name      string
	Admin     bool
	Hash      []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GameServer struct {
	UserID       entity.UserID
	ID           entity.GameServerID
	Ip           string
	Port         uint32
	RCONPassword string
	DisplayName  string
	IsPublic     bool
	Status       entity.SERVER_STATUS
}

type Match struct {
	ID         entity.MatchID
	UserID     entity.UserID
	ServerID   entity.GameServerID
	Team1ID    entity.TeamID
	Team2ID    entity.TeamID
	Winner     entity.TeamID
	StartTime  *time.Time
	EndTime    *time.Time
	MaxMaps    int32
	Title      string
	SkipVeto   bool
	APIKey     string
	Team1Score uint32
	Team2Score uint32
	Forfeit    *bool
	Status     entity.MATCH_STATUS
}

type MapStats struct {
	ID         entity.MapStatsID
	MatchID    entity.MatchID
	MapNumber  uint32
	MapName    string
	StartTime  *time.Time
	EndTime    *time.Time
	Winner     *entity.TeamID
	Team1Score uint32
	Team2Score uint32
}

type PlayerStats struct {
	ID      entity.PlayerStatsID
	MatchID entity.MatchID
	MapID   entity.MapStatsID
	TeamID  entity.TeamID
	SteamID entity.SteamID
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
	ID     entity.TeamID
	UserID entity.UserID
	Name   string
	Flag   string
	Tag    string
	Logo   string
	Public bool
}

type Player struct {
	ID      entity.PlayerID
	TeamID  entity.TeamID
	SteamID entity.SteamID
	Name    string
}
