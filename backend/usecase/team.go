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

	players, err := playerRepository.GetPlayersByTeam(ctx, teamID)
	if err != nil {
		return nil, err
	}
	return convertTeam(team, players), nil
}

// GetTeam implements Team.
func (t *team) GetTeam(ctx context.Context, id entity.TeamID) (*entity.Team, error) {
	repositoryConnector := database.GetConnection(ctx)

	teamsRepository := repositoryConnector.GetTeamsRepository()
	playersRepository := repositoryConnector.GetPlayersRepository()

	team, err := teamsRepository.GetTeam(ctx, id)
	if err != nil {
		return nil, err
	}

	players, err := playersRepository.GetPlayersByTeam(ctx, id)
	if err != nil {
		return nil, err
	}

	return convertTeam(team, players), nil
}

// GetTeamsByUser implements Team.
func (t *team) GetTeamsByUser(ctx context.Context, userID entity.UserID) ([]*entity.Team, error) {
	repositoryConnector := database.GetConnection(ctx)

	teamRepository := repositoryConnector.GetTeamsRepository()
	playersRepository := repositoryConnector.GetPlayersRepository()

	teams, err := teamRepository.GetTeamsByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	ret := make([]*entity.Team, 0, len(teams))
	for _, team := range teams {
		players, err := playersRepository.GetPlayersByTeam(ctx, team.ID)
		if err != nil {
			return nil, err
		}

		ret = append(ret, convertTeam(team, players))
	}
	return ret, nil
}
