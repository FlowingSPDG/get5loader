package usecase

import (
	"context"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
)

type GameServer interface {
	// GetPublicServers returns public game servers.
	GetPublicServers(ctx context.Context) ([]*entity.GameServer, error)
	// GetGameServer returns a game server.
	GetGameServer(ctx context.Context, id entity.GameServerID) (*entity.GameServer, error)
	// GetGameServersByUser returns game servers owned by a user.
	GetGameServersByUser(ctx context.Context, userID entity.UserID) ([]*entity.GameServer, error)
	// AddGameServer adds a game server.
	AddGameServer(ctx context.Context, userID entity.UserID, ip string, port uint32, rconPassword string, name string, isPublic bool) (*entity.GameServer, error)
	// DeleteGameServer deletes a game server.
	DeleteGameServer(ctx context.Context, id entity.GameServerID) error
}

type gameServer struct {
	repositoryConnector database.RepositoryConnector
}

func NewGameServer(
	repositoryConnector database.RepositoryConnector,
) GameServer {
	return &gameServer{
		repositoryConnector: repositoryConnector,
	}
}

// GetPublicServers implements GameServer.
func (gs *gameServer) GetPublicServers(ctx context.Context) ([]*entity.GameServer, error) {
	if err := gs.repositoryConnector.Open(); err != nil {
		return nil, err
	}
	defer gs.repositoryConnector.Close()

	repository := gs.repositoryConnector.GetGameServersRepository()

	gameServers, err := repository.GetPublicGameServers(ctx)
	if err != nil {
		return nil, err
	}

	return convertGameServers(gameServers), nil
}

// AddGameServer implements GameServer.
func (gs *gameServer) AddGameServer(ctx context.Context, userID entity.UserID, ip string, port uint32, rconPassword string, name string, isPublic bool) (*entity.GameServer, error) {
	// TODO: Query SRCDS
	// TODO: Check RCON Password
	if err := gs.repositoryConnector.Open(); err != nil {
		return nil, err
	}
	defer gs.repositoryConnector.Close()

	repository := gs.repositoryConnector.GetGameServersRepository()

	gameServerID, err := repository.AddGameServer(ctx, userID, ip, port, rconPassword, name, isPublic)
	if err != nil {
		return nil, err
	}

	gameServer, err := repository.GetGameServer(ctx, gameServerID)
	if err != nil {
		return nil, err
	}

	return convertGameServer(gameServer), nil
}

// DeleteGameServer implements GameServer.
func (gs *gameServer) DeleteGameServer(ctx context.Context, id entity.GameServerID) error {
	panic("unimplemented")
}

// GetGameServer implements GameServer.
func (gs *gameServer) GetGameServer(ctx context.Context, id entity.GameServerID) (*entity.GameServer, error) {
	panic("unimplemented")
}

// GetGameServersByUser implements GameServer.
func (gs *gameServer) GetGameServersByUser(ctx context.Context, userID entity.UserID) ([]*entity.GameServer, error) {
	if err := gs.repositoryConnector.Open(); err != nil {
		return nil, err
	}
	defer gs.repositoryConnector.Close()

	repository := gs.repositoryConnector.GetGameServersRepository()

	gss, err := repository.GetGameServersByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return convertGameServers(gss), nil
}
