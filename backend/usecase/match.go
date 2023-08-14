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
}

func NewMatch() Match {
	return &match{}
}

func (gm *match) GetMatch(ctx context.Context, matchID entity.MatchID) (*entity.Match, error) {
	// TODO: publicでない場合の認証処理の追加
	repositoryConnector := database.GetConnection(ctx)

	matchRepository := repositoryConnector.GetMatchesRepository()

	match, err := matchRepository.GetMatch(ctx, matchID)
	if err != nil {
		return nil, err
	}

	return convertMatch(match), nil
}

// CreateMatch implements Match.
func (gm *match) CreateMatch(ctx context.Context, userID entity.UserID, serverID entity.GameServerID, team1ID entity.TeamID, team2ID entity.TeamID, maxMaps int, title string) (*entity.Match, error) {
	repositoryConnector := database.GetConnection(ctx)

	matchRepository := repositoryConnector.GetMatchesRepository()

	mID, err := matchRepository.AddMatch(ctx, userID, serverID, team1ID, team2ID, int32(maxMaps), title, false, "")
	if err != nil {
		return nil, err
	}

	m, err := matchRepository.GetMatch(ctx, mID)
	if err != nil {
		return nil, err
	}

	match := convertMatch(m)

	return match, nil
}

// GetMatchesByUser implements Match.
func (gm *match) GetMatchesByUser(ctx context.Context, userID entity.UserID) ([]*entity.Match, error) {
	// TODO: publicでない場合の認証処理の追加
	repositoryConnector := database.GetConnection(ctx)

	matchRepository := repositoryConnector.GetMatchesRepository()

	matches, err := matchRepository.GetMatchesByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return convertMatches(matches), nil
}
