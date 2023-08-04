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

type GameServerRepository interface {
	AddGameServer(ctx context.Context, userID int64, ip string, port int32, rconPassword string, displayName string, isPublic bool) (*entity.GameServer, error)
	GetGameServer(ctx context.Context, id int64) (*entity.GameServer, error)
	GetPublicGameServers(ctx context.Context) ([]*entity.GameServer, error)
	GetGameServersByUser(ctx context.Context, userID int64) ([]*entity.GameServer, error)
	DeleteGameServer(ctx context.Context, id int64) error
}

type MatchRepository interface {
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
}

type PlayerStatsRepository interface {
}

type TeamRepository interface {
}

type PlayerRepository interface {
}
