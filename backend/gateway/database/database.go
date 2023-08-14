// Package database provides interfaces for database connection and repositories for various entities.
package database

import (
	"context"
	"database/sql"

	"github.com/FlowingSPDG/get5loader/backend/entity"
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
	// GetMapStatRepository returns a map stats repository. You must open a repository connection before calling this method.
	GetMapStatRepository() MapStatRepository
	// GetPlayerStatRepository returns a player stats repository. You must open a repository connection before calling this method.
	GetPlayerStatRepository() PlayerStatRepository
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
	// GetMapStatRepository returns a map stats repository. You must open a repository connection before calling this method.
	GetMapStatRepository() MapStatRepository
	// GetPlayerStatRepository returns a player stats repository. You must open a repository connection before calling this method.
	GetPlayerStatRepository() PlayerStatRepository
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
	CreateUser(ctx context.Context, steamID entity.SteamID, name string, admin bool, hash []byte) (entity.UserID, error)
	// GetUser returns a user.
	GetUser(ctx context.Context, id entity.UserID) (*User, error)
	// GetUserBySteamID returns a user.
	GetUserBySteamID(ctx context.Context, steamID entity.SteamID) (*User, error)
}

// GameServersRepository is an interface for game server repository.
type GameServersRepository interface {
	// AddGameServer adds a game server.
	AddGameServer(ctx context.Context, userID entity.UserID, ip string, port uint32, rconPassword string, displayName string, isPublic bool) (entity.GameServerID, error)
	// GetGameServer returns a game server.
	GetGameServer(ctx context.Context, id entity.GameServerID) (*GameServer, error)
	// GetPublicGameServers returns public game servers.
	GetPublicGameServers(ctx context.Context) ([]*GameServer, error)
	// GetGameServersByUser returns game servers owned by a user.
	GetGameServersByUser(ctx context.Context, userID entity.UserID) ([]*GameServer, error)
	// GetGameServersByUsers returns game servers owned by users.
	GetGameServersByUsers(ctx context.Context, userIDs []entity.UserID) (map[entity.UserID][]*GameServer, error)
	// DeleteGameServer deletes a game server.
	DeleteGameServer(ctx context.Context, id entity.GameServerID) error
}

// MatchesRepository is an interface for match repository.
type MatchesRepository interface {
	// AddMatch adds a match.
	AddMatch(ctx context.Context, userID entity.UserID, serverID entity.GameServerID, team1ID entity.TeamID, team2ID entity.TeamID, maxMaps int32, title string, skipVeto bool, apiKey string) (entity.MatchID, error)
	// GetMatch returns a match.
	GetMatch(ctx context.Context, id entity.MatchID) (*Match, error)
	// GetMatchesByUser returns matches owned by a user.
	GetMatchesByUser(ctx context.Context, userID entity.UserID) ([]*Match, error)
	// GetMatchesByUsers returns matches owned by users.
	GetMatchesByUsers(ctx context.Context, userIDs []entity.UserID) (map[entity.UserID][]*Match, error)
	// GetMatchesByTeam returns matches owned by a team.
	GetMatchesByTeam(ctx context.Context, teamID entity.TeamID) ([]*Match, error)
	// GetMatchesByWinner returns matches won by a team.
	GetMatchesByWinner(ctx context.Context, teamID entity.TeamID) ([]*Match, error)
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

// MapStatRepository is an interface for map stats repository.
type MapStatRepository interface {
	// TODO.
	// AddMapStats(ctx context.Context, matchID int64, mapNumber uint32, mapName string, winnerID int64, team1Score uint32, team2Score uint32) (*entity.MapStats, error)
	// GetMapStats returns map stats.
	GetMapStat(ctx context.Context, id entity.MapStatsID) (*MapStat, error)
	// GetMapStatsByMatch returns map stats owned by a match.
	GetMapStatsByMatch(ctx context.Context, matchID entity.MatchID) ([]*MapStat, error)
	// BatchGetMapStatsByMatches returns map stats owned by matches.
	GetMapStatsByMatches(ctx context.Context, matchIDs []entity.MatchID) (map[entity.MatchID][]*MapStat, error)
	// GetMapStatsByMatchAndMap returns map stats owned by a match and map number.
	GetMapStatsByMatchAndMap(ctx context.Context, matchID entity.MatchID, mapNumber uint32) (*MapStat, error)
}

// PlayerStatRepository is an interface for player stats repository.
type PlayerStatRepository interface {
	// TODO.
	// AddPlayerStats(ctx context.Context, mapStatsID int64, steamID string, name string, teamID int64, kills uint32, assists uint32, deaths uint32, hs uint32, flashAssists uint32, kast float32, rating float32) (*entity.PlayerStats, error)
	// GetPlayerStatsBySteamID returns player stats owned by a steam ID.
	GetPlayerStatsBySteamID(ctx context.Context, steamID entity.SteamID) ([]*PlayerStat, error)
	// GetPlayerStatsByMatch returns player stats owned by a match.
	GetPlayerStatsByMatch(ctx context.Context, matchID entity.MatchID) ([]*PlayerStat, error)
	// GetPlayerStatsByMapstats returns player stats owned by a map stats.
	GetPlayerStatsByMapstats(ctx context.Context, mapStatsID []entity.MapStatsID) (map[entity.MapStatsID][]*PlayerStat, error)
}

// TeamsRepository is an interface for team repository.
type TeamsRepository interface {
	// AddTeam adds a team.
	AddTeam(ctx context.Context, userID entity.UserID, name string, tag string, flag string, logo string, public bool) (entity.TeamID, error)
	// GetTeam returns a team.
	GetTeam(ctx context.Context, id entity.TeamID) (*Team, error)
	// GetTeams returns teams.
	GetTeams(ctx context.Context, ids []entity.TeamID) ([]*Team, error)
	// GetTeamsByUser returns teams owned by a user.
	GetTeamsByUser(ctx context.Context, userID entity.UserID) ([]*Team, error)
	// GetTeamsByUsers returns teams owned by users.
	GetTeamsByUsers(ctx context.Context, userIDs []entity.UserID) (map[entity.UserID][]*Team, error)
	// GetPublicTeams returns public teams.
	GetPublicTeams(ctx context.Context) ([]*Team, error)
}

// PlayersRepository is an interface for player repository.
type PlayersRepository interface {
	// AddPlayer adds a player.
	AddPlayer(ctx context.Context, teamID entity.TeamID, steamID entity.SteamID, name string) (entity.PlayerID, error)
	// GetPlayer returns a player.
	GetPlayer(ctx context.Context, id entity.PlayerID) (*Player, error)
	// GetPlayersByTeam returns players owned by a team.
	GetPlayersByTeam(ctx context.Context, teamID entity.TeamID) ([]*Player, error)
	// GetPlayersByTeams returns players owned by teams.
	GetPlayersByTeams(ctx context.Context, teamIDs []entity.TeamID) (map[entity.TeamID][]*Player, error)
	// DeletePlayer deletes a player.
	// DeletePlayer(ctx context.Context, id int64) error
}
