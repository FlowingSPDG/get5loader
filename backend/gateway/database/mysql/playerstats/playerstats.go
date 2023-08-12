package playerstats

import (
	"context"
	"database/sql"
	"errors"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
	playerstats_gen "github.com/FlowingSPDG/get5loader/backend/gateway/database/mysql/playerstats/generated"
	"github.com/FlowingSPDG/get5loader/backend/service/uuid"
)

type playerStatsRepository struct {
	uuidGenerator uuid.UUIDGenerator
	queries       *playerstats_gen.Queries
}

func NewPlayerStatsRepository(uuidGenerator uuid.UUIDGenerator, db *sql.DB) database.PlayerStatsRepository {
	queries := playerstats_gen.New(db)
	return &playerStatsRepository{
		uuidGenerator: uuidGenerator,
		queries:       queries,
	}
}

func NewPlayerStatsRepositoryWithTx(uuidGenerator uuid.UUIDGenerator, db *sql.DB, tx *sql.Tx) database.PlayerStatsRepository {
	queries := playerstats_gen.New(db).WithTx(tx)
	return &playerStatsRepository{
		uuidGenerator: uuidGenerator,
		queries:       queries,
	}
}

// GetPlayerStatsByMapstats implements database.PlayerStatsRepository.
func (psr *playerStatsRepository) GetPlayerStatsByMapstats(ctx context.Context, mapstatsID entity.MapStatsID) ([]*database.PlayerStats, error) {
	stats, err := psr.queries.GetPlayerStatsByMap(ctx, string(mapstatsID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.NewNotFoundError(err)
		}
		return nil, database.NewInternalError(err)
	}

	ret := make([]*database.PlayerStats, 0, len(stats))
	for _, stat := range stats {
		ret = append(ret, &database.PlayerStats{
			ID:               entity.PlayerStatsID(stat.ID),
			MatchID:          entity.MatchID(stat.MatchID),
			MapID:            entity.MapStatsID(stat.MapID),
			TeamID:           entity.TeamID(stat.TeamID),
			SteamID:          entity.SteamID(stat.SteamID),
			Name:             stat.Name,
			Kills:            stat.Kills,
			Assists:          stat.Assists,
			Deaths:           stat.Deaths,
			FlashbangAssists: stat.FlashbangAssists,
			Suicides:         stat.Suicides,
			HeadShotKills:    stat.HeadshotKills,
			Damage:           stat.Damage,
			BombPlants:       stat.BombPlants,
			BombDefuses:      stat.BombDefuses,
			V1:               stat.V1,
			V2:               stat.V2,
			V3:               stat.V3,
			V4:               stat.V4,
			V5:               stat.V5,
			K1:               stat.K1,
			K2:               stat.K2,
			K3:               stat.K3,
			K4:               stat.K4,
			K5:               stat.K5,
			FirstDeathCT:     stat.FirstdeathCt,
			FirstDeathT:      stat.FirstdeathCt,
			FirstKillCT:      stat.FirstkillCt,
			FirstKillT:       stat.FirstkillT,
		})
	}

	return ret, nil

}

// GetPlayerStatsByMatch implements database.PlayerStatsRepository.
func (psr *playerStatsRepository) GetPlayerStatsByMatch(ctx context.Context, matchID entity.MatchID) ([]*database.PlayerStats, error) {
	stats, err := psr.queries.GetPlayerStatsByMatch(ctx, string(matchID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.NewNotFoundError(err)
		}
		return nil, database.NewInternalError(err)
	}

	playerStats := make([]*database.PlayerStats, 0, len(stats))
	for _, stat := range stats {
		playerStats = append(playerStats, &database.PlayerStats{
			ID:               entity.PlayerStatsID(stat.ID),
			MatchID:          entity.MatchID(stat.MatchID),
			MapID:            entity.MapStatsID(stat.MapID),
			TeamID:           entity.TeamID(stat.TeamID),
			SteamID:          entity.SteamID(stat.SteamID),
			Name:             stat.Name,
			Kills:            stat.Kills,
			Assists:          stat.Assists,
			Deaths:           stat.Deaths,
			FlashbangAssists: stat.FlashbangAssists,
			Suicides:         stat.Suicides,
			HeadShotKills:    stat.HeadshotKills,
			Damage:           stat.Damage,
			BombPlants:       stat.BombPlants,
			BombDefuses:      stat.BombDefuses,
			V1:               stat.V1,
			V2:               stat.V2,
			V3:               stat.V3,
			V4:               stat.V4,
			V5:               stat.V5,
			K1:               stat.K1,
			K2:               stat.K2,
			K3:               stat.K3,
			K4:               stat.K4,
			K5:               stat.K5,
			FirstDeathCT:     stat.FirstdeathCt,
			FirstDeathT:      stat.FirstdeathCt,
			FirstKillCT:      stat.FirstkillCt,
			FirstKillT:       stat.FirstkillT,
		})
	}
	return playerStats, nil
}

// GetPlayerStatsBySteamID implements database.PlayerStatsRepository.
func (psr *playerStatsRepository) GetPlayerStatsBySteamID(ctx context.Context, steamID entity.SteamID) ([]*database.PlayerStats, error) {
	stats, err := psr.queries.GetPlayerStatsBySteamID(ctx, uint64(steamID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.NewNotFoundError(err)
		}
		return nil, database.NewInternalError(err)
	}

	playerStats := make([]*database.PlayerStats, 0, len(stats))
	for _, stat := range stats {
		playerStats = append(playerStats, &database.PlayerStats{
			ID:               entity.PlayerStatsID(stat.ID),
			MatchID:          entity.MatchID(stat.MatchID),
			MapID:            entity.MapStatsID(stat.MapID),
			TeamID:           entity.TeamID(stat.TeamID),
			SteamID:          entity.SteamID(stat.SteamID),
			Name:             stat.Name,
			Kills:            stat.Kills,
			Assists:          stat.Assists,
			Deaths:           stat.Deaths,
			FlashbangAssists: stat.FlashbangAssists,
			Suicides:         stat.Suicides,
			HeadShotKills:    stat.HeadshotKills,
			Damage:           stat.Damage,
			BombPlants:       stat.BombPlants,
			BombDefuses:      stat.BombDefuses,
			V1:               stat.V1,
			V2:               stat.V2,
			V3:               stat.V3,
			V4:               stat.V4,
			V5:               stat.V5,
			K1:               stat.K1,
			K2:               stat.K2,
			K3:               stat.K3,
			K4:               stat.K4,
			K5:               stat.K5,
			FirstDeathCT:     stat.FirstdeathCt,
			FirstDeathT:      stat.FirstdeathCt,
			FirstKillCT:      stat.FirstkillCt,
			FirstKillT:       stat.FirstkillT,
		})
	}
	return playerStats, nil

}
