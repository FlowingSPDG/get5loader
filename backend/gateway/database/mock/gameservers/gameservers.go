package gameservers

import (
	"context"
	"net"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
)

type mockGameServersRepository struct {
	gameServerToReturnID      map[entity.GameServerID]*entity.GameServer
	gameServersToReturnUserID map[entity.UserID][]*entity.GameServer
	gameServersToReturnPublic []*entity.GameServer
}

func NewMockGameServersRepository(
	gameServerToReturnID map[entity.GameServerID]*entity.GameServer,
	gameServersToReturnUserID map[entity.UserID][]*entity.GameServer,
	gameServersToReturnPublic []*entity.GameServer,
) database.GameServersRepository {
	return &mockGameServersRepository{
		gameServerToReturnID:      gameServerToReturnID,
		gameServersToReturnUserID: gameServersToReturnUserID,
		gameServersToReturnPublic: gameServersToReturnPublic,
	}
}

// AddGameServer implements database.GameServersRepository.
func (mgsr *mockGameServersRepository) AddGameServer(ctx context.Context, userID entity.UserID, ip net.IP, port uint32, rconPassword string, displayName string, isPublic bool) error {
	return nil
}

// DeleteGameServer implements database.GameServersRepository.
func (mgsr *mockGameServersRepository) DeleteGameServer(ctx context.Context, id entity.GameServerID) error {
	return nil
}

// GetGameServer implements database.GameServersRepository.
func (mgsr *mockGameServersRepository) GetGameServer(ctx context.Context, id entity.GameServerID) (*entity.GameServer, error) {
	v, ok := mgsr.gameServerToReturnID[id]
	if !ok {
		return nil, database.ErrNotFound
	}
	return v, nil
}

// GetGameServersByUser implements database.GameServersRepository.
func (mgsr *mockGameServersRepository) GetGameServersByUser(ctx context.Context, userID entity.UserID) ([]*entity.GameServer, error) {
	v, ok := mgsr.gameServersToReturnUserID[userID]
	if !ok {
		return nil, database.ErrNotFound
	}
	return v, nil
}

// GetPublicGameServers implements database.GameServersRepository.
func (mgsr *mockGameServersRepository) GetPublicGameServers(ctx context.Context) ([]*entity.GameServer, error) {
	return mgsr.gameServersToReturnPublic, nil
}
