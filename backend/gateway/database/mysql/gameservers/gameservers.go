package gameservers

import (
	"context"
	"database/sql"
	"errors"
	"net"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	gameservers_gen "github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/gameservers/generated"
)

type gameServerRepository struct {
	queries *gameservers_gen.Queries
}

func NewGameServerRepository(db *sql.DB) database.GameServersRepository {
	queries := gameservers_gen.New(db)
	return &gameServerRepository{
		queries: queries,
	}
}

func NewGameServerRepositoryWithTx(db *sql.DB, tx *sql.Tx) database.GameServersRepository {
	queries := gameservers_gen.New(db).WithTx(tx)
	return &gameServerRepository{
		queries: queries,
	}
}

// AddGameServer implements database.GameServerRepository.
func (gr *gameServerRepository) AddGameServer(ctx context.Context, userID entity.UserID, ip net.IP, port uint32, rconPassword string, displayName string, isPublic bool) error {
	if _, err := gr.queries.AddGameServer(ctx, gameservers_gen.AddGameServerParams{
		UserID:       string(userID),
		Ip:           ip,
		Port:         port,
		RconPassword: rconPassword,
		DisplayName:  displayName,
		IsPublic:     isPublic,
	}); err != nil {
		return database.NewInternalError(err)
	}

	return nil
}

// DeleteGameServer implements database.GameServerRepository.
func (gr *gameServerRepository) DeleteGameServer(ctx context.Context, id entity.GameServerID) error {
	if _, err := gr.queries.DeleteGameServer(ctx, string(id)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return database.NewNotFoundError(err)
		}
		return database.NewInternalError(err)
	}
	return nil
}

// GetGameServer implements database.GameServerRepository.
func (gr *gameServerRepository) GetGameServer(ctx context.Context, id entity.GameServerID) (*entity.GameServer, error) {
	gameserver, err := gr.queries.GetGameServers(ctx, string(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.NewNotFoundError(err)
		}
		return nil, database.NewInternalError(err)
	}

	return &entity.GameServer{
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
func (gr *gameServerRepository) GetGameServersByUser(ctx context.Context, userID entity.UserID) ([]*entity.GameServer, error) {
	gameservers, err := gr.queries.GetGameServersByUser(ctx, string(userID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.NewNotFoundError(err)
		}
		return nil, database.NewInternalError(err)
	}

	ret := make([]*entity.GameServer, 0, len(gameservers))
	for _, gameserver := range gameservers {
		ret = append(ret, &entity.GameServer{
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
func (gr *gameServerRepository) GetPublicGameServers(ctx context.Context) ([]*entity.GameServer, error) {
	gameservers, err := gr.queries.GetPublicGameServers(ctx)
	if err != nil {
		return nil, database.NewInternalError(err)
	}

	ret := make([]*entity.GameServer, 0, len(gameservers))
	for _, gameserver := range gameservers {
		ret = append(ret, &entity.GameServer{
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
