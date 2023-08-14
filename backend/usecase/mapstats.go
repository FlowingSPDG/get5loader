package usecase

import (
	"context"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
)

type Mapstat interface {
	GetMapStat(ctx context.Context, id entity.MapStatsID) (*entity.MapStat, error)
	GetMapStatsByMatch(ctx context.Context, matchID entity.MatchID) ([]*entity.MapStat, error)
	// GetMapStatsByTeam(ctx context.Context, teamID entity.TeamID) ([]*entity.MapStat, error)
}

type mapstat struct {
}

func NewMapStats() Mapstat {
	return &mapstat{}
}

// GetMapStats implements Mapstats.
func (m *mapstat) GetMapStat(ctx context.Context, id entity.MapStatsID) (*entity.MapStat, error) {
	repositoryConnector := database.GetConnection(ctx)

	MapStatRepository := repositoryConnector.GetMapStatRepository()
	PlayerStatRepository := repositoryConnector.GetPlayerStatRepository()

	mapstats, err := MapStatRepository.GetMapStat(ctx, id)
	if err != nil {
		return nil, err
	}

	playerStats, err := PlayerStatRepository.GetPlayerStatsByMapstats(ctx, id)
	if err != nil {
		return nil, err
	}

	return convertMapstat(mapstats, playerStats), nil
}

// GetMapStatsByMatch implements Mapstats.
func (m *mapstat) GetMapStatsByMatch(ctx context.Context, matchID entity.MatchID) ([]*entity.MapStat, error) {
	repositoryConnector := database.GetConnection(ctx)

	MapStatRepository := repositoryConnector.GetMapStatRepository()
	PlayerStatRepository := repositoryConnector.GetPlayerStatRepository()

	mapstats, err := MapStatRepository.GetMapStatsByMatch(ctx, matchID)
	if err != nil {
		return nil, err
	}

	ret := make([]*entity.MapStat, 0, len(mapstats))
	for _, mapstat := range mapstats {
		playerStats, err := PlayerStatRepository.GetPlayerStatsByMapstats(ctx, mapstat.ID)
		if err != nil {
			return nil, err
		}
		ret = append(ret, convertMapstat(mapstat, playerStats))
	}

	return ret, nil
}
