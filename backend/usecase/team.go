package usecase

import (
	"context"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
)

type Team interface {
	RegisterTeam(ctx context.Context, userID entity.UserID, name string, flag string, tag string, logo string, publicTeam bool) (*entity.Team, error)
	GetTeam(ctx context.Context, id entity.TeamID) (*entity.Team, error)
	GetTeamsByUser(ctx context.Context, userID entity.UserID) ([]*entity.Team, error)
}

type team struct {
	repositoryConnector database.RepositoryConnector
}

func NewTeam(repositoryConnector database.RepositoryConnector) Team {
	return &team{
		repositoryConnector: repositoryConnector,
	}
}

func (t *team) RegisterTeam(ctx context.Context, userID entity.UserID, name string, flag string, tag string, logo string, publicTeam bool) (*entity.Team, error) {
	if err := t.repositoryConnector.Open(); err != nil {
		return nil, err
	}
	defer t.repositoryConnector.Close()

	repository := t.repositoryConnector.GetTeamsRepository()

	teamID, err := repository.AddTeam(ctx, userID, name, tag, flag, logo, publicTeam)
	if err != nil {
		return nil, err
	}
	team, err := repository.GetTeam(ctx, teamID)
	if err != nil {
		return nil, err
	}

	return &entity.Team{
		ID:      entity.TeamID(team.ID),
		UserID:  entity.UserID(team.UserID),
		Name:    team.Name,
		Flag:    team.Flag,
		Tag:     team.Tag,
		Logo:    team.Logo,
		Public:  team.Public,
		Players: []*entity.Player{},
	}, nil
}

// GetTeam implements Team.
func (t *team) GetTeam(ctx context.Context, id entity.TeamID) (*entity.Team, error) {
	if err := t.repositoryConnector.Open(); err != nil {
		return nil, err
	}
	defer t.repositoryConnector.Close()

	teamsRepository := t.repositoryConnector.GetTeamsRepository()
	playersRepository := t.repositoryConnector.GetPlayersRepository()

	team, err := teamsRepository.GetTeam(ctx, id)
	if err != nil {
		return nil, err
	}

	players, err := playersRepository.GetPlayersByTeam(ctx, id)
	if err != nil {
		return nil, err
	}

	teamPlayers := make([]*entity.Player, 0, len(players))
	for _, player := range players {
		teamPlayers = append(teamPlayers, &entity.Player{
			ID:      entity.PlayerID(player.ID),
			TeamID:  entity.TeamID(player.TeamID),
			SteamID: entity.SteamID(player.SteamID),
			Name:    player.Name,
		})
	}

	return &entity.Team{
		ID:      entity.TeamID(team.ID),
		UserID:  entity.UserID(team.UserID),
		Name:    team.Name,
		Flag:    team.Flag,
		Tag:     team.Tag,
		Logo:    team.Logo,
		Public:  team.Public,
		Players: teamPlayers,
	}, nil
}

// GetTeamsByUser implements Team.
func (t *team) GetTeamsByUser(ctx context.Context, userID entity.UserID) ([]*entity.Team, error) {
	if err := t.repositoryConnector.Open(); err != nil {
		return nil, err
	}
	defer t.repositoryConnector.Close()

	teamRepository := t.repositoryConnector.GetTeamsRepository()
	playersRepository := t.repositoryConnector.GetPlayersRepository()

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

		teamPlayers := make([]*entity.Player, 0, len(players))
		for _, player := range players {
			teamPlayers = append(teamPlayers, &entity.Player{
				ID:      entity.PlayerID(player.ID),
				TeamID:  entity.TeamID(player.TeamID),
				SteamID: entity.SteamID(player.SteamID),
				Name:    player.Name,
			})
		}

		ret = append(ret, &entity.Team{
			ID:      entity.TeamID(team.ID),
			UserID:  entity.UserID(team.UserID),
			Name:    team.Name,
			Flag:    team.Flag,
			Tag:     team.Tag,
			Logo:    team.Logo,
			Public:  team.Public,
			Players: teamPlayers,
		})
	}
	return ret, nil
}
