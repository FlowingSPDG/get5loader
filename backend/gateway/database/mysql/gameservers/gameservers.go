package gameservers

import (
	"context"
	"database/sql"
	"errors"
	"net"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
	gameservers_gen "github.com/FlowingSPDG/get5loader/backend/gateway/database/mysql/gameservers/generated"
	"github.com/FlowingSPDG/get5loader/backend/service/uuid"
)

type gameServerRepository struct {
	uuidGenerator uuid.UUIDGenerator
	queries       *gameservers_gen.Queries
}

func NewGameServerRepository(uuidGenerator uuid.UUIDGenerator, db *sql.DB) database.GameServersRepository {
	queries := gameservers_gen.New(db)
	return &gameServerRepository{
		uuidGenerator: uuidGenerator,
		queries:       queries,
	}
}

func NewGameServerRepositoryWithTx(uuidGenerator uuid.UUIDGenerator, db *sql.DB, tx *sql.Tx) database.GameServersRepository {
	queries := gameservers_gen.New(db).WithTx(tx)
	return &gameServerRepository{
		uuidGenerator: uuidGenerator,
		queries:       queries,
	}
}

// AddGameServer implements database.GameServerRepository.
func (gr *gameServerRepository) AddGameServer(ctx context.Context, userID entity.UserID, ip string, port uint32, rconPassword string, displayName string, isPublic bool) (entity.GameServerID, error) {
	id := gr.uuidGenerator.Generate()
	if _, err := gr.queries.AddGameServer(ctx, gameservers_gen.AddGameServerParams{
		ID:           id,
		UserID:       string(userID),
		Ip:           ip,
		Port:         port,
		RconPassword: rconPassword,
		DisplayName:  displayName,
		IsPublic:     isPublic,
	}); err != nil {
		return "", database.NewInternalError(err)
	}

	return entity.GameServerID(id), nil
}

// DeleteGameServer implements database.GameServerRepository.
func (gr *gameServerRepository) DeleteGameServer(ctx context.Context, id entity.GameServerID) error {
	if _, err := gr.queries.DeleteGameServer(ctx, string(id)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return database.NewNotFoundError(nil)
		}
		return database.NewInternalError(err)
	}
	return nil
}

// GetGameServer implements database.GameServerRepository.
func (gr *gameServerRepository) GetGameServer(ctx context.Context, id entity.GameServerID) (*database.GameServer, error) {
	gameserver, err := gr.queries.GetGameServer(ctx, string(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.NewNotFoundError(nil)
		}
		return nil, database.NewInternalError(err)
	}

	return &database.GameServer{
		ID:           entity.GameServerID(gameserver.ID),
		UserID:       entity.UserID(gameserver.UserID),
		Ip:           net.ParseIP(string(gameserver.Ip)).To4().String(),
		Port:         gameserver.Port,
		RCONPassword: gameserver.RconPassword,
		DisplayName:  gameserver.DisplayName,
		IsPublic:     gameserver.IsPublic,
	}, nil
}

// GetGameServersByUser implements database.GameServerRepository.
func (gr *gameServerRepository) GetGameServersByUser(ctx context.Context, userID entity.UserID) ([]*database.GameServer, error) {
	gameservers, err := gr.queries.GetGameServersByUser(ctx, string(userID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.NewNotFoundError(nil)
		}
		return nil, database.NewInternalError(err)
	}

	ret := make([]*database.GameServer, 0, len(gameservers))
	for _, gameserver := range gameservers {
		ret = append(ret, &database.GameServer{
			ID:           entity.GameServerID(gameserver.ID),
			UserID:       entity.UserID(gameserver.UserID),
			Ip:           net.ParseIP(string(gameserver.Ip)).To4().String(),
			Port:         gameserver.Port,
			RCONPassword: gameserver.RconPassword,
			DisplayName:  gameserver.DisplayName,
			IsPublic:     gameserver.IsPublic,
		})
	}

	return ret, nil
}

// GetGameServersByUsers implements database.GameServersRepository.
func (gr *gameServerRepository) GetGameServersByUsers(ctx context.Context, userIDs []entity.UserID) (map[entity.UserID][]*database.GameServer, error) {
	ids := database.IDsToString(userIDs)
	gameservers, err := gr.queries.GetGameServersByUsers(ctx, ids)
	if err != nil {
		return nil, database.NewInternalError(err)
	}

	ret := make(map[entity.UserID][]*database.GameServer, len(userIDs))
	// nilが渡されるのを防ぐため、空のスライスを生成する
	for _, userID := range userIDs {
		ret[userID] = make([]*database.GameServer, 0)
	}

	for _, gameserver := range gameservers {
		ret[entity.UserID(gameserver.UserID)] = append(ret[entity.UserID(gameserver.UserID)], &database.GameServer{
			ID:           entity.GameServerID(gameserver.ID),
			UserID:       entity.UserID(gameserver.UserID),
			Ip:           net.ParseIP(string(gameserver.Ip)).To4().String(),
			Port:         gameserver.Port,
			RCONPassword: gameserver.RconPassword,
			DisplayName:  gameserver.DisplayName,
			IsPublic:     gameserver.IsPublic,
		})
	}

	return ret, nil
}

// GetPublicGameServers implements database.GameServerRepository.
func (gr *gameServerRepository) GetPublicGameServers(ctx context.Context) ([]*database.GameServer, error) {
	gameservers, err := gr.queries.GetPublicGameServers(ctx)
	if err != nil {
		return nil, database.NewInternalError(err)
	}

	ret := make([]*database.GameServer, 0, len(gameservers))
	for _, gameserver := range gameservers {
		ret = append(ret, &database.GameServer{
			ID:           entity.GameServerID(gameserver.ID),
			UserID:       entity.UserID(gameserver.UserID),
			Ip:           net.ParseIP(string(gameserver.Ip)).To4().String(),
			Port:         gameserver.Port,
			RCONPassword: gameserver.RconPassword,
			DisplayName:  gameserver.DisplayName,
			IsPublic:     gameserver.IsPublic,
		})
	}

	return ret, nil
}
