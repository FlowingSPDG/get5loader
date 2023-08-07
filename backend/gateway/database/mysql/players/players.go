package players

import (
	"context"
	"database/sql"
	"errors"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	players_gen "github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/players/generated"
)

type playersRepository struct {
	queries *players_gen.Queries
}

func NewPlayersRepository(db *sql.DB) database.PlayersRepository {
	queries := players_gen.New(db)
	return &playersRepository{
		queries: queries,
	}
}

func NewPlayersRepositoryWithTx(db *sql.DB, tx *sql.Tx) database.PlayersRepository {
	queries := players_gen.New(db).WithTx(tx)
	return &playersRepository{
		queries: queries,
	}
}

// AddPlayer implements database.PlayersRepository.
func (pr *playersRepository) AddPlayer(ctx context.Context, teamID entity.TeamID, steamID entity.SteamID, name string) error {
	if _, err := pr.queries.AddPlayer(ctx, players_gen.AddPlayerParams{
		TeamID:  string(teamID),
		SteamID: uint64(steamID),
		Name:    name,
	}); err != nil {
		return database.NewInternalError(err)
	}

	return nil
}

// GetPlayer implements database.PlayersRepository.
func (pr *playersRepository) GetPlayer(ctx context.Context, id entity.PlayerStatsID) (*entity.Player, error) {
	res, err := pr.queries.GetPlayer(ctx, string(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.NewNotFoundError(err)
		}
		return nil, database.NewInternalError(err)
	}

	return &entity.Player{
		ID:      entity.PlayerStatsID(res.ID),
		TeamID:  entity.TeamID(res.TeamID),
		SteamID: entity.SteamID(res.SteamID),
		Name:    res.Name,
	}, nil
}

// GetPlayersByTeam implements database.PlayersRepository.
func (pr *playersRepository) GetPlayersByTeam(ctx context.Context, teamID entity.TeamID) ([]*entity.Player, error) {
	res, err := pr.queries.GetPlayersByTeam(ctx, string(teamID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.NewNotFoundError(err)
		}
		return nil, database.NewInternalError(err)
	}

	players := make([]*entity.Player, 0, len(res))
	for _, p := range res {
		players = append(players, &entity.Player{
			ID:      entity.PlayerStatsID(p.ID),
			TeamID:  entity.TeamID(p.TeamID),
			SteamID: entity.SteamID(p.SteamID),
			Name:    p.Name,
		})
	}

	return players, nil
}
