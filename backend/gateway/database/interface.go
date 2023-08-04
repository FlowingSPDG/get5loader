package database

import (
	"context"
	"time"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
)

type UsersRepositry interface {
	CreateUser(ctx context.Context, steamID string, name string, admin bool) (*entity.User, error)
	GetUser(ctx context.Context, id int64) (*entity.User, error)
}

type GameServersRepository interface {
	AddGameServer(ctx context.Context, userID int64, ip string, port int32, rconPassword string, displayName string, isPublic bool) (*entity.GameServer, error)
	GetGameServer(ctx context.Context, id int64) (*entity.GameServer, error)
	GetPublicGameServers(ctx context.Context) ([]*entity.GameServer, error)
	GetGameServersByUser(ctx context.Context, userID int64) ([]*entity.GameServer, error)
	DeleteGameServer(ctx context.Context, id int64) error
}

type MatchesRepository interface {
	AddMatch(ctx context.Context, userID int64, serverID int64, team1ID int64, team2ID int64, startTime time.Time, endTime time.Time, maxMaps int32, title string, skipVeto bool, apiKey string) (*entity.Match, error)
	GetMatch(ctx context.Context, id int64) (*entity.Match, error)
	GetMatchesByUser(ctx context.Context, userID int64) ([]*entity.Match, error)
	GetMatchesByTeam(ctx context.Context, teamID int64) ([]*entity.Match, error)
	GetMatchesByWinner(ctx context.Context, teamID int64) ([]*entity.Match, error)
	UpdateMatchWinner(ctx context.Context, matchID int64, winnerID int64) error
	UpdateTeam1Score(ctx context.Context, matchID int64, score int32) error
	UpdateTeam2Score(ctx context.Context, matchID int64, score int32) error
	CancelMatch(ctx context.Context, matchID int64) error
	StartMatch(ctx context.Context, matchID int64) error
}

type MapStatsRepository interface {
	GetMapStats(ctx context.Context, id int64) (*entity.MapStats, error)
	GetMapStatsByMatch(ctx context.Context, mapStatsID int64) ([]*entity.MapStats, error)
}

type PlayerStatsRepository interface {
	GetPlayerStatsBySteamID(ctx context.Context, steamID string) ([]*entity.PlayerStats, error)
	GetPlayerStatsByMatch(ctx context.Context, matchID int64) ([]*entity.PlayerStats, error)
	GetPlayerStatsByMapstats(ctx context.Context, mapStatsID int64) (*entity.PlayerStats, error)
}

type TeamsRepository interface {
	AddTeam(ctx context.Context, userID int64, name string, tag string, flag string, logo string) (*entity.Team, error)
	GetTeam(ctx context.Context, id int64) (*entity.Team, error)
	GetTeamsByUser(ctx context.Context, userID int64) ([]*entity.Team, error)
	GetPublicTeams(ctx context.Context) ([]*entity.Team, error)
}

type PlayersRepository interface {
	AddPlayer(ctx context.Context, teamID int64, steamID string, name string) (*entity.Player, error)
	GetPlayer(ctx context.Context, id int64) (*entity.Player, error)
	GetPlayersByTeam(ctx context.Context, teamID int64) ([]*entity.Player, error)
}
