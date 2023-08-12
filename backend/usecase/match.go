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
	MapStatRepository := gm.repositoryConnector.GetMapStatRepository()
	PlayerStatRepository := gm.repositoryConnector.GetPlayerStatRepository()
	teamRepository := gm.repositoryConnector.GetTeamsRepository()
	playerRepository := gm.repositoryConnector.GetPlayersRepository()

	match, err := matchRepository.GetMatch(ctx, matchID)
	if err != nil {
		return nil, err
	}
	mapstats, err := MapStatRepository.GetMapStatsByMatch(ctx, matchID)
	if err != nil {
		return nil, err
	}

	matchMapStats := make([]*entity.MapStat, 0, len(mapstats))
	for _, mapstat := range mapstats {
		playerStats, err := PlayerStatRepository.GetPlayerStatsByMapstats(ctx, mapstat.ID)
		if err != nil {
			return nil, err
		}
		matchMapStats = append(matchMapStats, convertMapstat(mapstat, playerStats))
	}

	team1, err := teamRepository.GetTeam(ctx, match.Team1ID)
	if err != nil {
		return nil, err
	}
	team1players, err := playerRepository.GetPlayersByTeam(ctx, team1.ID)
	if err != nil {
		return nil, err
	}

	team2, err := teamRepository.GetTeam(ctx, match.Team2ID)
	if err != nil {
		return nil, err
	}
	team2players, err := playerRepository.GetPlayersByTeam(ctx, team2.ID)
	if err != nil {
		return nil, err
	}

	return convertMatch(match, team1, team2, team1players, team2players, matchMapStats), nil
}
