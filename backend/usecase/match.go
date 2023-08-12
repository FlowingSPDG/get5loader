package usecase

import (
	"context"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
)

type Match interface {
	GetMatch(ctx context.Context, matchID entity.MatchID) (*entity.Match, error)
}

type match struct {
	repositoryConnector database.RepositoryConnector
}

func NewMatch(
	repositoryConnector database.RepositoryConnector,
) Match {
	return &match{
		repositoryConnector: repositoryConnector,
	}
}

func (gm *match) GetMatch(ctx context.Context, matchID entity.MatchID) (*entity.Match, error) {
	// TODO: publicでない場合の認証処理の追加
	if err := gm.repositoryConnector.Open(); err != nil {
		return nil, err
	}
	defer gm.repositoryConnector.Close()

	matchRepository := gm.repositoryConnector.GetMatchesRepository()
	mapStatsRepository := gm.repositoryConnector.GetMapStatsRepository()
	playerStatsRepository := gm.repositoryConnector.GetPlayerStatsRepository()
	gameServerRepository := gm.repositoryConnector.GetGameServersRepository()
	teamRepository := gm.repositoryConnector.GetTeamsRepository()
	playerRepository := gm.repositoryConnector.GetPlayersRepository()

	match, err := matchRepository.GetMatch(ctx, matchID)
	if err != nil {
		return nil, err
	}
	mapstats, err := mapStatsRepository.GetMapStatsByMatch(ctx, matchID)
	if err != nil {
		return nil, err
	}

	matchMapStats := make([]*entity.MapStats, 0, len(mapstats))
	for _, mapstat := range mapstats {
		playerStats, err := playerStatsRepository.GetPlayerStatsByMapstats(ctx, mapstat.ID)
		if err != nil {
			return nil, err
		}
		matchPlayerStats := make([]*entity.PlayerStats, 0, len(playerStats))
		for _, playerStat := range playerStats {
			matchPlayerStats = append(matchPlayerStats, &entity.PlayerStats{
				ID:               entity.PlayerStatsID(playerStat.ID),
				MatchID:          entity.MatchID(playerStat.MatchID),
				MapID:            entity.MapStatsID(playerStat.MapID),
				TeamID:           entity.TeamID(playerStat.TeamID),
				SteamID:          entity.SteamID(playerStat.SteamID),
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
			matchMapStats = append(matchMapStats, &entity.MapStats{
				ID:          entity.MapStatsID(mapstat.ID),
				MatchID:     entity.MatchID(mapstat.MatchID),
				MapNumber:   mapstat.MapNumber,
				MapName:     mapstat.MapName,
				StartTime:   mapstat.StartTime,
				EndTime:     mapstat.EndTime,
				Winner:      mapstat.Winner,
				Team1Score:  mapstat.Team1Score,
				Team2Score:  mapstat.Team2Score,
				PlayerStats: matchPlayerStats,
			})
		}
	}

	gameServer, err := gameServerRepository.GetGameServer(ctx, match.ServerID)
	if err != nil {
		return nil, err
	}

	team1, err := teamRepository.GetTeam(ctx, match.Team1ID)
	if err != nil {
		return nil, err
	}
	team1players, err := playerRepository.GetPlayersByTeam(ctx, team1.ID)
	if err != nil {
		return nil, err
	}
	t1p := make([]*entity.Player, 0, len(team1players))
	for _, player := range team1players {
		t1p = append(t1p, &entity.Player{
			ID:      entity.PlayerID(player.ID),
			TeamID:  entity.TeamID(player.TeamID),
			SteamID: entity.SteamID(player.SteamID),
			Name:    player.Name,
		})
	}

	team2, err := teamRepository.GetTeam(ctx, match.Team2ID)
	if err != nil {
		return nil, err
	}
	team2players, err := playerRepository.GetPlayersByTeam(ctx, team2.ID)
	t2p := make([]*entity.Player, 0, len(team2players))
	for _, player := range team2players {
		t2p = append(t2p, &entity.Player{
			ID:      entity.PlayerID(player.ID),
			TeamID:  entity.TeamID(player.TeamID),
			SteamID: entity.SteamID(player.SteamID),
			Name:    player.Name,
		})
	}

	return &entity.Match{
		ID:     entity.MatchID(match.ID),
		UserID: entity.UserID(match.UserID),
		GameServer: entity.GameServer{
			ID:          entity.GameServerID(gameServer.ID),
			Ip:          gameServer.Ip,
			Port:        gameServer.Port,
			DisplayName: gameServer.DisplayName,
			IsPublic:    gameServer.IsPublic,
		},
		Team1: entity.Team{
			ID:      entity.TeamID(team1.ID),
			UserID:  entity.UserID(team1.UserID),
			Name:    team1.Name,
			Flag:    team1.Flag,
			Tag:     team1.Tag,
			Logo:    team1.Logo,
			Public:  team1.Public,
			Players: t1p,
		},
		Team2: entity.Team{
			ID:      entity.TeamID(team2.ID),
			UserID:  entity.UserID(team2.UserID),
			Name:    team2.Name,
			Flag:    team2.Flag,
			Tag:     team2.Tag,
			Logo:    team2.Logo,
			Public:  team2.Public,
			Players: t2p,
		},
		Winner:     match.Winner,
		StartTime:  match.StartTime,
		EndTime:    match.EndTime,
		MaxMaps:    match.MaxMaps,
		Title:      match.Title,
		SkipVeto:   match.SkipVeto,
		APIKey:     match.APIKey,
		Team1Score: match.Team1Score,
		Team2Score: match.Team2Score,
		Forfeit:    match.Forfeit,
		Status:     match.Status,
		Mapstats:   matchMapStats,
	}, nil
}
