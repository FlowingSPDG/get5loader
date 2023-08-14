package players

import (
	"context"
	"database/sql"
	"errors"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
	players_gen "github.com/FlowingSPDG/get5loader/backend/gateway/database/mysql/players/generated"
	"github.com/FlowingSPDG/get5loader/backend/service/uuid"
)

type playersRepository struct {
	uuidGenerator uuid.UUIDGenerator
	queries       *players_gen.Queries
}

func NewPlayersRepository(uuidGenerator uuid.UUIDGenerator, db *sql.DB) database.PlayersRepository {
	queries := players_gen.New(db)
	return &playersRepository{
		uuidGenerator: uuidGenerator,
		queries:       queries,
	}
}

func NewPlayersRepositoryWithTx(uuidGenerator uuid.UUIDGenerator, db *sql.DB, tx *sql.Tx) database.PlayersRepository {
	queries := players_gen.New(db).WithTx(tx)
	return &playersRepository{
		uuidGenerator: uuidGenerator,
		queries:       queries,
	}
}

// AddPlayer implements database.PlayersRepository.
func (pr *playersRepository) AddPlayer(ctx context.Context, teamID entity.TeamID, steamID entity.SteamID, name string) (entity.PlayerID, error) {
	id := pr.uuidGenerator.Generate()
	if _, err := pr.queries.AddPlayer(ctx, players_gen.AddPlayerParams{
		ID:      id,
		TeamID:  string(teamID),
		SteamID: uint64(steamID),
		Name:    name,
	}); err != nil {
		return "", database.NewInternalError(err)
	}

	return entity.PlayerID(id), nil
}

// GetPlayer implements database.PlayersRepository.
func (pr *playersRepository) GetPlayer(ctx context.Context, id entity.PlayerID) (*database.Player, error) {
	res, err := pr.queries.GetPlayer(ctx, string(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.NewNotFoundError(err)
		}
		return nil, database.NewInternalError(err)
	}

	return &database.Player{
		ID:      entity.PlayerID(res.ID),
		TeamID:  entity.TeamID(res.TeamID),
		SteamID: entity.SteamID(res.SteamID),
		Name:    res.Name,
	}, nil
}

// GetPlayersByTeam implements database.PlayersRepository.
func (pr *playersRepository) GetPlayersByTeam(ctx context.Context, teamID entity.TeamID) ([]*database.Player, error) {
	res, err := pr.queries.GetPlayersByTeam(ctx, string(teamID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []*database.Player{}, nil
		}
		return nil, database.NewInternalError(err)
	}

	players := make([]*database.Player, 0, len(res))
	for _, p := range res {
		players = append(players, &database.Player{
			ID:      entity.PlayerID(p.ID),
			TeamID:  entity.TeamID(p.TeamID),
			SteamID: entity.SteamID(p.SteamID),
			Name:    p.Name,
		})
	}

	return players, nil
}
