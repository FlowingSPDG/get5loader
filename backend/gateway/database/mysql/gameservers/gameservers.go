package gameservers

import (
	"context"
	"database/sql"
	"errors"

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
func (gr *gameServerRepository) AddGameServer(ctx context.Context, userID int64, ip string, port uint32, rconPassword string, displayName string, isPublic bool) (*entity.GameServer, error) {
	result, err := gr.queries.AddGameServer(ctx, gameservers_gen.AddGameServerParams{
		UserID:       userID,
		Ip:           ip,
		Port:         port,
		RconPassword: rconPassword,
		DisplayName:  displayName,
		IsPublic:     isPublic,
	})
	if err != nil {
		return nil, database.NewInternalError(err)
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.NewNotFoundError(err)
		}
		return nil, database.NewInternalError(err)
	}

	gameserver, err := gr.queries.GetGameServers(ctx, insertedID)
	if err != nil {
		return nil, database.NewInternalError(err)
	}

	return &entity.GameServer{
		ID:           gameserver.ID,
		UserID:       gameserver.UserID,
		Ip:           gameserver.Ip,
		Port:         gameserver.Port,
		RCONPassword: gameserver.RconPassword,
		DisplayName:  gameserver.DisplayName,
		IsPublic:     gameserver.IsPublic,
	}, nil
}

// DeleteGameServer implements database.GameServerRepository.
func (gr *gameServerRepository) DeleteGameServer(ctx context.Context, id int64) error {
	if _, err := gr.queries.DeleteGameServer(ctx, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return database.NewNotFoundError(err)
		}
		return database.NewInternalError(err)
	}
	return nil
}

// GetGameServer implements database.GameServerRepository.
func (gr *gameServerRepository) GetGameServer(ctx context.Context, id int64) (*entity.GameServer, error) {
	gameserver, err := gr.queries.GetGameServers(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.NewNotFoundError(err)
		}
		return nil, database.NewInternalError(err)
	}

	return &entity.GameServer{
		ID:           gameserver.ID,
		UserID:       gameserver.UserID,
		Ip:           gameserver.Ip,
		Port:         gameserver.Port,
		RCONPassword: gameserver.RconPassword,
		DisplayName:  gameserver.DisplayName,
		IsPublic:     gameserver.IsPublic,
	}, nil
}

// GetGameServersByUser implements database.GameServerRepository.
func (gr *gameServerRepository) GetGameServersByUser(ctx context.Context, userID int64) ([]*entity.GameServer, error) {
	gameservers, err := gr.queries.GetGameServersByUser(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.NewNotFoundError(err)
		}
		return nil, database.NewInternalError(err)
	}

	ret := make([]*entity.GameServer, 0, len(gameservers))
	for _, gameserver := range gameservers {
		ret = append(ret, &entity.GameServer{
			ID:           gameserver.ID,
			UserID:       gameserver.UserID,
			Ip:           gameserver.Ip,
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
			ID:           gameserver.ID,
			UserID:       gameserver.UserID,
			Ip:           gameserver.Ip,
			Port:         gameserver.Port,
			RCONPassword: gameserver.RconPassword,
			DisplayName:  gameserver.DisplayName,
			IsPublic:     gameserver.IsPublic,
		})
	}

	return ret, nil
}
