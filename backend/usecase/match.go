package usecase

import (
	"context"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
)

type Match interface {
	CreateMatch(ctx context.Context, userID entity.UserID, serverID entity.GameServerID, team1ID entity.TeamID, team2ID entity.TeamID, maxMaps int, title string) (*entity.Match, error)
	GetMatch(ctx context.Context, matchID entity.MatchID) (*entity.Match, error)
	GetMatchesByUser(ctx context.Context, userID entity.UserID) ([]*entity.Match, error)
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

// CreateMatch implements Match.
func (gm *match) CreateMatch(ctx context.Context, userID entity.UserID, serverID entity.GameServerID, team1ID entity.TeamID, team2ID entity.TeamID, maxMaps int, title string) (*entity.Match, error) {
	// TODO: teamが存在しない場合のエラーハンドリング
	if err := gm.repositoryConnector.Open(); err != nil {
		return nil, err
	}
	defer gm.repositoryConnector.Close()

	matchRepository := gm.repositoryConnector.GetMatchesRepository()
	teamRepository := gm.repositoryConnector.GetTeamsRepository()
	playerRepository := gm.repositoryConnector.GetPlayersRepository()

	mID, err := matchRepository.AddMatch(ctx, userID, serverID, team1ID, team2ID, int32(maxMaps), title, false, "")
	if err != nil {
		return nil, err
	}

	m, err := matchRepository.GetMatch(ctx, mID)
	if err != nil {
		return nil, err
	}

	team1, err := teamRepository.GetTeam(ctx, team1ID)
	if err != nil {
		return nil, err
	}
	team1players, err := playerRepository.GetPlayersByTeam(ctx, team1.ID)
	if err != nil {
		return nil, err
	}
	team2, err := teamRepository.GetTeam(ctx, team2ID)
	if err != nil {
		return nil, err
	}
	team2players, err := playerRepository.GetPlayersByTeam(ctx, team2.ID)
	if err != nil {
		return nil, err
	}

	match := convertMatch(m, team1, team2, team1players, team2players, nil)

	return match, nil
}

// GetMatchesByUser implements Match.
func (gm *match) GetMatchesByUser(ctx context.Context, userID entity.UserID) ([]*entity.Match, error) {
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

	matches, err := matchRepository.GetMatchesByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	ret := make([]*entity.Match, 0, len(matches))
	for _, match := range matches {
		mapstats, err := MapStatRepository.GetMapStatsByMatch(ctx, match.ID)
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

		ret = append(ret, convertMatch(match, team1, team2, team1players, team2players, matchMapStats))
	}

	return ret, nil
}
