package usecase

import (
	"context"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
)

type Player interface {
	GetPlayersByTeam(ctx context.Context, teamID entity.TeamID) ([]*entity.Player, error)
	// BATCH
	BatchGetPlayersByTeam(ctx context.Context, teamIDs []entity.TeamID) (map[entity.TeamID][]*entity.Player, error)
}

type player struct {
}

func NewPlayer() Player {
	return &player{}
}

// GetPlayersByTeam implements Player.
func (p *player) GetPlayersByTeam(ctx context.Context, teamID entity.TeamID) ([]*entity.Player, error) {
	repositoryConnector := database.GetConnection(ctx)

	playerRepository := repositoryConnector.GetPlayersRepository()

	players, err := playerRepository.GetPlayersByTeam(ctx, teamID)
	if err != nil {
		return nil, err
	}

	return convertPlayers(players), nil
}

// BatchGetPlayersByTeam implements Player.
func (p *player) BatchGetPlayersByTeam(ctx context.Context, teamIDs []entity.TeamID) (map[entity.TeamID][]*entity.Player, error) {
	repositoryConnector := database.GetConnection(ctx)

	playerRepository := repositoryConnector.GetPlayersRepository()

	teamPlayers, err := playerRepository.GetPlayersByTeams(ctx, teamIDs)
	if err != nil {
		return nil, err
	}

	ret := make(map[entity.TeamID][]*entity.Player, len(teamIDs))
	for teamID, players := range teamPlayers {
		ret[teamID] = convertPlayers(players)
	}

	return ret, nil
}
