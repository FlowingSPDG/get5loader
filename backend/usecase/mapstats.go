package usecase

import (
	"context"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
)

type Mapstats interface {
	GetMapStats(ctx context.Context, id entity.MapStatsID) (*entity.MapStats, error)
	GetMapStatsByMatch(ctx context.Context, matchID entity.MatchID) ([]*entity.MapStats, error)
	// GetMapStatsByTeam(ctx context.Context, teamID entity.TeamID) ([]*entity.MapStats, error)
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
func (m *mapstats) GetMapStats(ctx context.Context, id entity.MapStatsID) (*entity.MapStats, error) {
	if err := m.repositoryConnector.Open(); err != nil {
		return nil, err
	}
	defer m.repositoryConnector.Close()

	mapStatsRepository := m.repositoryConnector.GetMapStatsRepository()
	playerStatsRepository := m.repositoryConnector.GetPlayerStatsRepository()

	mapstats, err := mapStatsRepository.GetMapStats(ctx, id)
	if err != nil {
		return nil, err
	}

	fetchedPlayerStats, err := playerStatsRepository.GetPlayerStatsByMapstats(ctx, id)
	if err != nil {
		return nil, err
	}

	playerStats := make([]*entity.PlayerStats, 0, len(fetchedPlayerStats))

	for _, playerStat := range fetchedPlayerStats {
		playerStats = append(playerStats, &entity.PlayerStats{
			ID:               playerStat.ID,
			MatchID:          playerStat.MatchID,
			MapID:            playerStat.MapID,
			TeamID:           playerStat.TeamID,
			SteamID:          playerStat.SteamID,
			Name:             playerStat.Name,
			Kills:            playerStat.Kills,
			Assists:          playerStat.Assists,
			Deaths:           playerStat.Deaths,
			RoundsPlayed:     playerStat.RoundsPlayed,
			FlashbangAssists: playerStat.FlashbangAssists,
			Suicides:         playerStat.Suicides,
			HeadShotKills:    playerStat.HeadShotKills,
			Damage:           playerStat.Damage,
			BombPlants:       playerStat.BombPlants,
			BombDefuses:      playerStat.BombDefuses,
			V1:               playerStat.V1,
			V2:               playerStat.V2,
			V3:               playerStat.V3,
			V4:               playerStat.V4,
			V5:               playerStat.V5,
			K1:               playerStat.K1,
			K2:               playerStat.K2,
			K3:               playerStat.K3,
			K4:               playerStat.K4,
			K5:               playerStat.K5,
			FirstDeathCT:     playerStat.FirstDeathCT,
			FirstDeathT:      playerStat.FirstDeathT,
			FirstKillCT:      playerStat.FirstKillCT,
			FirstKillT:       playerStat.FirstKillT,
		})
	}

	return &entity.MapStats{
		ID:          entity.MapStatsID(mapstats.ID),
		MatchID:     entity.MatchID(mapstats.MatchID),
		MapNumber:   mapstats.MapNumber,
		MapName:     mapstats.MapName,
		StartTime:   mapstats.StartTime,
		EndTime:     mapstats.EndTime,
		Winner:      mapstats.Winner,
		Team1Score:  mapstats.Team1Score,
		Team2Score:  mapstats.Team2Score,
		PlayerStats: []*entity.PlayerStats{},
	}, nil
}

// GetMapStatsByMatch implements Mapstats.
func (m *mapstats) GetMapStatsByMatch(ctx context.Context, matchID entity.MatchID) ([]*entity.MapStats, error) {
	if err := m.repositoryConnector.Open(); err != nil {
		return nil, err
	}
	defer m.repositoryConnector.Close()

	mapStatsRepository := m.repositoryConnector.GetMapStatsRepository()
	playerStatsRepository := m.repositoryConnector.GetPlayerStatsRepository()

	mapstats, err := mapStatsRepository.GetMapStatsByMatch(ctx, matchID)
	if err != nil {
		return nil, err
	}

	ret := make([]*entity.MapStats, 0, len(mapstats))
	for _, mapstat := range mapstats {
		fetchedPlayerStats, err := playerStatsRepository.GetPlayerStatsByMapstats(ctx, mapstat.ID)
		if err != nil {
			return nil, err
		}

		playerStats := make([]*entity.PlayerStats, 0, len(fetchedPlayerStats))

		for _, playerStat := range fetchedPlayerStats {
			playerStats = append(playerStats, &entity.PlayerStats{
				ID:               playerStat.ID,
				MatchID:          playerStat.MatchID,
				MapID:            playerStat.MapID,
				TeamID:           playerStat.TeamID,
				SteamID:          playerStat.SteamID,
				Name:             playerStat.Name,
				Kills:            playerStat.Kills,
				Assists:          playerStat.Assists,
				Deaths:           playerStat.Deaths,
				RoundsPlayed:     playerStat.RoundsPlayed,
				FlashbangAssists: playerStat.FlashbangAssists,
				Suicides:         playerStat.Suicides,
				HeadShotKills:    playerStat.HeadShotKills,
				Damage:           playerStat.Damage,
				BombPlants:       playerStat.BombPlants,
				BombDefuses:      playerStat.BombDefuses,
				V1:               playerStat.V1,
				V2:               playerStat.V2,
				V3:               playerStat.V3,
				V4:               playerStat.V4,
				V5:               playerStat.V5,
				K1:               playerStat.K1,
				K2:               playerStat.K2,
				K3:               playerStat.K3,
				K4:               playerStat.K4,
				K5:               playerStat.K5,
				FirstDeathCT:     playerStat.FirstDeathCT,
				FirstDeathT:      playerStat.FirstDeathT,
				FirstKillCT:      playerStat.FirstKillCT,
				FirstKillT:       playerStat.FirstKillT,
			})
		}

		ret = append(ret, &entity.MapStats{
			ID:          entity.MapStatsID(mapstat.ID),
			MatchID:     mapstat.MatchID,
			MapNumber:   mapstat.MapNumber,
			MapName:     mapstat.MapName,
			StartTime:   mapstat.StartTime,
			EndTime:     mapstat.EndTime,
			Winner:      mapstat.Winner,
			Team1Score:  mapstat.Team1Score,
			Team2Score:  mapstat.Team2Score,
			PlayerStats: playerStats,
		})
	}

	return ret, nil
}
