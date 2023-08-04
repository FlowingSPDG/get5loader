package gameservers

import (
	"context"
	"database/sql"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	gameservers_gen "github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/gameservers/generated"
)

type gameServerRepository struct {
	dsn string
}

func NewGameServerRepository(dsn string) database.GameServerRepository {
	return &gameServerRepository{
		dsn: dsn,
	}
}

func (gsr *gameServerRepository) Open() (*sql.DB, error) {
	return sql.Open("mysql", gsr.dsn)
}

// AddGameServer implements database.GameServerRepository.
func (gr *gameServerRepository) AddGameServer(ctx context.Context, userID int64, ip string, port int32, rconPassword string, displayName string, isPublic bool) (*entity.GameServer, error) {
	db, err := gr.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	queries := gameservers_gen.New(db)

	result, err := queries.AddGameServer(ctx, gameservers_gen.AddGameServerParams{
		UserID:       userID,
		Ip:           ip,
		Port:         port,
		RconPassword: rconPassword,
		DisplayName:  displayName,
		IsPublic:     isPublic,
	})
	if err != nil {
		return nil, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	gameserver, err := queries.GetGameServers(ctx, insertedID)
	if err != nil {
		return nil, err
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
	db, err := gr.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	queries := gameservers_gen.New(db)

	if _, err := queries.DeleteGameServer(ctx, id); err != nil {
		return err
	}
	return nil
}

// GetGameServer implements database.GameServerRepository.
func (gr *gameServerRepository) GetGameServer(ctx context.Context, id int64) (*entity.GameServer, error) {
	db, err := gr.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	queries := gameservers_gen.New(db)

	gameserver, err := queries.GetGameServers(ctx, id)
	if err != nil {
		return nil, err
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
	db, err := gr.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	queries := gameservers_gen.New(db)

	gameservers, err := queries.GetGameServersByUser(ctx, userID)
	if err != nil {
		return nil, err
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
	db, err := gr.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	queries := gameservers_gen.New(db)

	gameservers, err := queries.GetPublicGameServers(ctx)
	if err != nil {
		return nil, err
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
