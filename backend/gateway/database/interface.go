// Package database provides interfaces for database connection and repositories for various entities.
package database

import (
	"context"
	"database/sql"
	"net"
	"time"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
)

// DBConnector is a database connection initiator.
type DBConnector interface {
	// Open opens a database connection.
	// You must call Close after using the database.
	Open() error
	// GetConnection returns a opened database connection. Potentially returns nil.
	GetConnection() *sql.DB
	// Close closes a database connection.
	Close() error
}

// DBConnectorWithTx is a database connection initiator with transaction.
type DBConnectorWithTx interface {
	// Open opens a database connection.
	// You must call Close after using the database.
	Open() error
	// GetConnection returns a opened database connection. Potentially returns nil.
	GetConnection() *sql.DB
	// Close closes a database connection.
	Close() error
	// BeginTx starts a transaction.
	BeginTx() error
	// GetTx returns a transaction. Potentially returns nil.
	GetTx() *sql.Tx
	// Commit commits a transaction.
	Commit() error
	// Rollback rollbacks a transaction.
	Rollback() error
}

// RepositoryConnector is a generic interface for opening and closing a repository connection.
type RepositoryConnector interface {
	// Open opens a repository connection. You must call Close after using the repository.
	Open() error
	// Close closes a repository connection.
	Close() error

	// GetUserRepository returns a user repository. You must open a repository connection before calling this method.
	GetUserRepository() UsersRepositry
	// GetGameServersRepository returns a game server repository. You must open a repository connection before calling this method.
	GetGameServersRepository() GameServersRepository
	// GetMatchesRepository returns a match repository. You must open a repository connection before calling this method.
	GetMatchesRepository() MatchesRepository
	// GetMapStatsRepository returns a map stats repository. You must open a repository connection before calling this method.
	GetMapStatsRepository() MapStatsRepository
	// GetPlayerStatsRepository returns a player stats repository. You must open a repository connection before calling this method.
	GetPlayerStatsRepository() PlayerStatsRepository
	// GetTeamsRepository returns a team repository. You must open a repository connection before calling this method.
	GetTeamsRepository() TeamsRepository
	// GetPlayersRepository returns a player repository. You must open a repository connection before calling this method.
	GetPlayersRepository() PlayersRepository
}

// RepositoryConnectorWithTx is a interface for opening and closing a repository connection with transaction.
type RepositoryConnectorWithTx interface {
	// Open opens a repository connection. You must call Close after using the repository.
	// Transaction is started when Open is called.
	Open() error
	// Close closes a repository connection.
	Close() error

	// GetUserRepository returns a user repository. You must open a repository connection before calling this method.
	GetUserRepository() UsersRepositry
	// GetGameServersRepository returns a game server repository. You must open a repository connection before calling this method.
	GetGameServersRepository() GameServersRepository
	// GetMatchesRepository returns a match repository. You must open a repository connection before calling this method.
	GetMatchesRepository() MatchesRepository
	// GetMapStatsRepository returns a map stats repository. You must open a repository connection before calling this method.
	GetMapStatsRepository() MapStatsRepository
	// GetPlayerStatsRepository returns a player stats repository. You must open a repository connection before calling this method.
	GetPlayerStatsRepository() PlayerStatsRepository
	// GetTeamsRepository returns a team repository. You must open a repository connection before calling this method.
	GetTeamsRepository() TeamsRepository
	// GetPlayersRepository returns a player repository. You must open a repository connection before calling this method.
	GetPlayersRepository() PlayersRepository

	// Commit commits a transaction.
	Commit() error
	// Rollback rollbacks a transaction.
	Rollback() error
}

// UsersRepositry is an interface for user repository.
type UsersRepositry interface {
	// CreateUser creates a user.
	CreateUser(ctx context.Context, steamID entity.SteamID, name string, admin bool, hash []byte) error
	// GetUser returns a user.
	GetUser(ctx context.Context, id entity.UserID) (*entity.User, error)
	GetUserBySteamID(ctx context.Context, steamID entity.SteamID) (*entity.User, error)
}

// GameServersRepository is an interface for game server repository.
type GameServersRepository interface {
	// AddGameServer adds a game server.
	AddGameServer(ctx context.Context, userID entity.UserID, ip net.IP, port uint32, rconPassword string, displayName string, isPublic bool) error
	// GetGameServer returns a game server.
	GetGameServer(ctx context.Context, id entity.GameServerID) (*entity.GameServer, error)
	// GetPublicGameServers returns public game servers.
	GetPublicGameServers(ctx context.Context) ([]*entity.GameServer, error)
	// GetGameServersByUser returns game servers owned by a user.
	GetGameServersByUser(ctx context.Context, userID entity.UserID) ([]*entity.GameServer, error)
	// DeleteGameServer deletes a game server.
	DeleteGameServer(ctx context.Context, id entity.GameServerID) error
}

// MatchesRepository is an interface for match repository.
type MatchesRepository interface {
	// AddMatch adds a match.
	AddMatch(ctx context.Context, userID entity.UserID, serverID entity.GameServerID, team1ID entity.TeamID, team2ID entity.TeamID, startTime time.Time, endTime time.Time, maxMaps int32, title string, skipVeto bool, apiKey string) error
	// GetMatch returns a match.
	GetMatch(ctx context.Context, id entity.MatchID) (*entity.Match, error)
	// GetMatchesByUser returns matches owned by a user.
	GetMatchesByUser(ctx context.Context, userID entity.UserID) ([]*entity.Match, error)
	// GetMatchesByTeam returns matches owned by a team.
	GetMatchesByTeam(ctx context.Context, teamID entity.TeamID) ([]*entity.Match, error)
	// GetMatchesByWinner returns matches won by a team.
	GetMatchesByWinner(ctx context.Context, teamID entity.TeamID) ([]*entity.Match, error)
	// UpdateMatchWinner updates a match winner.
	UpdateMatchWinner(ctx context.Context, matchID entity.MatchID, winnerID entity.TeamID) error
	// UpdateTeam1Score updates a match team1 score.
	UpdateTeam1Score(ctx context.Context, matchID entity.MatchID, score uint32) error
	// UpdateTeam2Score updates a match team2 score.
	UpdateTeam2Score(ctx context.Context, matchID entity.MatchID, score uint32) error
	// CancelMatch cancels a match.
	CancelMatch(ctx context.Context, matchID entity.MatchID) error
	// StartMatch starts a match.
	StartMatch(ctx context.Context, matchID entity.MatchID) error
}

// MapStatsRepository is an interface for map stats repository.
type MapStatsRepository interface {
	// TODO.
	// AddMapStats(ctx context.Context, matchID int64, mapNumber uint32, mapName string, winnerID int64, team1Score uint32, team2Score uint32) (*entity.MapStats, error)
	// GetMapStats returns map stats.
	GetMapStats(ctx context.Context, id entity.MapStatsID) (*entity.MapStats, error)
	// GetMapStatsByMatch returns map stats owned by a match.
	GetMapStatsByMatch(ctx context.Context, matchID entity.MatchID) ([]*entity.MapStats, error)
	// GetMapStatsByMatchAndMap returns map stats owned by a match and map number.
	GetMapStatsByMatchAndMap(ctx context.Context, matchID entity.MatchID, mapNumber uint32) (*entity.MapStats, error)
}

// PlayerStatsRepository is an interface for player stats repository.
type PlayerStatsRepository interface {
	// TODO.
	// AddPlayerStats(ctx context.Context, mapStatsID int64, steamID string, name string, teamID int64, kills uint32, assists uint32, deaths uint32, hs uint32, flashAssists uint32, kast float32, rating float32) (*entity.PlayerStats, error)
	// GetPlayerStatsBySteamID returns player stats owned by a steam ID.
	GetPlayerStatsBySteamID(ctx context.Context, steamID entity.SteamID) ([]*entity.PlayerStats, error)
	// GetPlayerStatsByMatch returns player stats owned by a match.
	GetPlayerStatsByMatch(ctx context.Context, matchID entity.MatchID) ([]*entity.PlayerStats, error)
	// GetPlayerStatsByMapstats returns player stats owned by a map stats.
	GetPlayerStatsByMapstats(ctx context.Context, mapStatsID entity.MapStatsID) (*entity.PlayerStats, error)
}

// TeamsRepository is an interface for team repository.
type TeamsRepository interface {
	// AddTeam adds a team.
	AddTeam(ctx context.Context, userID entity.UserID, name string, tag string, flag string, logo string) error
	// GetTeam returns a team.
	GetTeam(ctx context.Context, id entity.TeamID) (*entity.Team, error)
	// GetTeamsByUser returns teams owned by a user.
	GetTeamsByUser(ctx context.Context, userID entity.UserID) ([]*entity.Team, error)
	// GetPublicTeams returns public teams.
	GetPublicTeams(ctx context.Context) ([]*entity.Team, error)
}

// PlayersRepository is an interface for player repository.
type PlayersRepository interface {
	// AddPlayer adds a player.
	AddPlayer(ctx context.Context, teamID entity.TeamID, steamID entity.SteamID, name string) error
	// GetPlayer returns a player.
	GetPlayer(ctx context.Context, id entity.PlayerID) (*entity.Player, error)
	// GetPlayersByTeam returns players owned by a team.
	GetPlayersByTeam(ctx context.Context, teamID entity.TeamID) ([]*entity.Player, error)
	// DeletePlayer deletes a player.
	// DeletePlayer(ctx context.Context, id int64) error
}
