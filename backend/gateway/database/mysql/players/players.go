package players

import (
	"context"
	"database/sql"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	players_gen "github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/players/generated"
)

type playersRepository struct {
	dsn string
}

func NewPlayersRepository(dsn string) database.PlayersRepository {
	return &playersRepository{
		dsn: dsn,
	}
}

func (pr *playersRepository) open() (*sql.DB, error) {
	return sql.Open("mysql", pr.dsn)
}

// AddPlayer implements database.PlayersRepository.
func (pr *playersRepository) AddPlayer(ctx context.Context, teamID int64, steamID string, name string) (*entity.Player, error) {
	db, err := pr.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	queries := players_gen.New(db)

	res, err := queries.AddPlayer(ctx, players_gen.AddPlayerParams{
		TeamID:  teamID,
		SteamID: steamID,
		Name:    name,
	})

	if err != nil {
		return nil, err
	}

	insertedID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	player, err := queries.GetPlayer(ctx, insertedID)
	if err != nil {
		return nil, err
	}

	return &entity.Player{
		ID:      player.ID,
		TeamID:  player.TeamID,
		SteamID: player.SteamID,
		Name:    player.Name,
	}, nil
}

// GetPlayer implements database.PlayersRepository.
func (pr *playersRepository) GetPlayer(ctx context.Context, id int64) (*entity.Player, error) {
	db, err := pr.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	queries := players_gen.New(db)

	res, err := queries.GetPlayer(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.Player{
		ID:      res.ID,
		TeamID:  res.TeamID,
		SteamID: res.SteamID,
		Name:    res.Name,
	}, nil
}

// GetPlayersByTeam implements database.PlayersRepository.
func (pr *playersRepository) GetPlayersByTeam(ctx context.Context, teamID int64) ([]*entity.Player, error) {
	db, err := pr.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	queries := players_gen.New(db)

	res, err := queries.GetPlayersByTeam(ctx, teamID)
	if err != nil {
		return nil, err
	}

	players := make([]*entity.Player, 0, len(res))
	for _, p := range res {
		players = append(players, &entity.Player{
			ID:      p.ID,
			TeamID:  p.TeamID,
			SteamID: p.SteamID,
			Name:    p.Name,
		})
	}

	return players, nil
}
