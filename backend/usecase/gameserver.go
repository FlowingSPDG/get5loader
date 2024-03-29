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
	// AddGameServer adds a game server.
	AddGameServer(ctx context.Context, userID entity.UserID, ip string, port uint32, rconPassword string, name string, isPublic bool) (*entity.GameServer, error)
	// DeleteGameServer deletes a game server.
	DeleteGameServer(ctx context.Context, id entity.GameServerID) error

	BatchGetGameServersByUser(ctx context.Context, userIDs []entity.UserID) (map[entity.UserID][]*entity.GameServer, error)
}

type gameServer struct {
}

func NewGameServer() GameServer {
	return &gameServer{}
}

// GetPublicServers implements GameServer.
func (gs *gameServer) GetPublicServers(ctx context.Context) ([]*entity.GameServer, error) {
	repositoryConnector := database.GetConnection(ctx)
	repository := repositoryConnector.GetGameServersRepository()

	gameServers, err := repository.GetPublicGameServers(ctx)
	if err != nil {
		return nil, err
	}

	return convertGameServers(gameServers), nil
}

// AddGameServer implements GameServer.
func (gs *gameServer) AddGameServer(ctx context.Context, userID entity.UserID, ip string, port uint32, rconPassword string, name string, isPublic bool) (*entity.GameServer, error) {
	repositoryConnector := database.GetConnection(ctx)
	repository := repositoryConnector.GetGameServersRepository()

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
	repositoryConnector := database.GetConnection(ctx)

	repository := repositoryConnector.GetGameServersRepository()

	gss, err := repository.GetGameServersByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return convertGameServers(gss), nil
}

// BatchGetGameServersByUser implements GameServer.
func (gs *gameServer) BatchGetGameServersByUser(ctx context.Context, userIDs []entity.UserID) (map[entity.UserID][]*entity.GameServer, error) {
	repositoryConnector := database.GetConnection(ctx)

	repository := repositoryConnector.GetGameServersRepository()

	gss, err := repository.GetGameServersByUsers(ctx, userIDs)
	if err != nil {
		return nil, err
	}

	ret := make(map[entity.UserID][]*entity.GameServer, len(userIDs))
	// nilが渡されるのを防ぐため、空のスライスを生成する
	for userID, gss := range gss {
		ret[userID] = convertGameServers(gss)
	}

	return ret, nil
}
