package usecase

import (
	"context"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
)

type Mapstats interface {
	GetMapStats(ctx context.Context, id entity.MapStatsID) (*entity.MapStat, error)
	GetMapStatsByMatch(ctx context.Context, matchID entity.MatchID) ([]*entity.MapStat, error)
	// GetMapStatsByTeam(ctx context.Context, teamID entity.TeamID) ([]*entity.MapStat, error)
}

type mapstats struct {
	repositoryConnector database.RepositoryConnector
}

func NewMapStats(repositoryConnector database.RepositoryConnector) Mapstats {
	return &mapstats{
		repositoryConnector: repositoryConnector,
	}
}

// GetMapStats implements Mapstats.
func (m *mapstats) GetMapStats(ctx context.Context, id entity.MapStatsID) (*entity.MapStat, error) {
	if err := m.repositoryConnector.Open(); err != nil {
		return nil, err
	}
	defer m.repositoryConnector.Close()

	MapStatRepository := m.repositoryConnector.GetMapStatRepository()
	PlayerStatRepository := m.repositoryConnector.GetPlayerStatRepository()

	mapstats, err := MapStatRepository.GetMapStats(ctx, id)
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
func (m *mapstats) GetMapStatsByMatch(ctx context.Context, matchID entity.MatchID) ([]*entity.MapStat, error) {
	if err := m.repositoryConnector.Open(); err != nil {
		return nil, err
	}
	defer m.repositoryConnector.Close()

	MapStatRepository := m.repositoryConnector.GetMapStatRepository()
	PlayerStatRepository := m.repositoryConnector.GetPlayerStatRepository()

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
