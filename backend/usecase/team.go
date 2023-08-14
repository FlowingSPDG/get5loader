package usecase

import (
	"context"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
)

type InputPlayers struct {
	SteamID entity.SteamID
	Name    string
}

type RegisterTeamInput struct {
	UserID     entity.UserID
	Name       string
	Flag       string
	Tag        string
	Logo       string
	PublicTeam bool
	Players    []InputPlayers
}

type Team interface {
	RegisterTeam(ctx context.Context, input RegisterTeamInput) (*entity.Team, error)
	GetTeam(ctx context.Context, id entity.TeamID) (*entity.Team, error)
	GetTeamsByMatch(ctx context.Context, matchID entity.MatchID) (*entity.Team, *entity.Team, error)
	GetTeamsByUser(ctx context.Context, userID entity.UserID) ([]*entity.Team, error)
}

type team struct {
}

func NewTeam() Team {
	return &team{}
}

func (t *team) RegisterTeam(ctx context.Context, input RegisterTeamInput) (*entity.Team, error) {
	repositoryConnector := database.GetConnection(ctx)

	teamRepository := repositoryConnector.GetTeamsRepository()
	playerRepository := repositoryConnector.GetPlayersRepository()

	teamID, err := teamRepository.AddTeam(ctx, input.UserID, input.Name, input.Tag, input.Flag, input.Logo, input.PublicTeam)
	if err != nil {
		return nil, err
	}
	// TODO: Batch addする
	for _, player := range input.Players {
		playerRepository.AddPlayer(ctx, teamID, player.SteamID, player.Name)
	}

	team, err := teamRepository.GetTeam(ctx, teamID)
	if err != nil {
		return nil, err
	}
	return convertTeam(team), nil
}

// GetTeam implements Team.
func (t *team) GetTeam(ctx context.Context, id entity.TeamID) (*entity.Team, error) {
	repositoryConnector := database.GetConnection(ctx)

	teamsRepository := repositoryConnector.GetTeamsRepository()

	team, err := teamsRepository.GetTeam(ctx, id)
	if err != nil {
		return nil, err
	}

	return convertTeam(team), nil
}

// GetTeamsByMatch implements Team.
func (t *team) GetTeamsByMatch(ctx context.Context, matchID entity.MatchID) (*entity.Team, *entity.Team, error) {
	repositoryConnector := database.GetConnection(ctx)

	matchRepository := repositoryConnector.GetMatchesRepository()
	teamRepository := repositoryConnector.GetTeamsRepository()

	match, err := matchRepository.GetMatch(ctx, matchID)
	if err != nil {
		return nil, nil, err
	}
	team1, err := teamRepository.GetTeam(ctx, match.Team1ID)
	if err != nil {
		return nil, nil, err
	}
	team2, err := teamRepository.GetTeam(ctx, match.Team2ID)
	if err != nil {
		return nil, nil, err
	}

	return convertTeam(team1), convertTeam(team2), nil
}

// GetTeamsByUser implements Team.
func (t *team) GetTeamsByUser(ctx context.Context, userID entity.UserID) ([]*entity.Team, error) {
	repositoryConnector := database.GetConnection(ctx)

	teamRepository := repositoryConnector.GetTeamsRepository()

	teams, err := teamRepository.GetTeamsByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	ret := make([]*entity.Team, 0, len(teams))
	for _, team := range teams {
		ret = append(ret, convertTeam(team))
	}
	return ret, nil
}
