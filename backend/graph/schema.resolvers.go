package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"strconv"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/g5ctx"
	"github.com/FlowingSPDG/get5loader/backend/graph/model"
)

// RegisterTeam is the resolver for the registerTeam field.
func (r *mutationResolver) RegisterTeam(ctx context.Context, input model.NewTeam) (*model.Team, error) {
	token, err := g5ctx.GetUserToken(ctx)
	if err != nil {
		return nil, err
	}
	team, err := r.TeamUsecase.RegisterTeam(ctx, token.UserID, input.Name, input.Flag, input.Tag, input.Logo, input.Public)
	if err != nil {
		return nil, err
	}
	return &model.Team{
		ID:     string(team.ID),
		Name:   team.Name,
		Tag:    team.Tag,
		Flag:   team.Flag,
		Logo:   team.Logo,
		Public: team.Public,
	}, nil
}

// AddServer is the resolver for the addServer field.
func (r *mutationResolver) AddServer(ctx context.Context, input model.NewGameServer) (*model.GameServer, error) {
	token, err := g5ctx.GetUserToken(ctx)
	if err != nil {
		return nil, err
	}
	gs, err := r.GameServerUsecase.AddGameServer(ctx, token.UserID, input.IP, uint32(input.Port), input.RconPassword, input.Name, input.Public)
	if err != nil {
		return nil, err
	}
	return &model.GameServer{
		ID:     string(gs.ID),
		IP:     gs.Ip,
		Port:   int(gs.Port),
		Name:   gs.DisplayName,
		Public: gs.IsPublic,
	}, nil
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: GetUser - getUser"))
}

// GetTeam is the resolver for the getTeam field.
func (r *queryResolver) GetTeam(ctx context.Context, id string) (*model.Team, error) {
	team, err := r.TeamUsecase.GetTeam(ctx, entity.TeamID(id))
	if err != nil {
		return nil, err
	}
	return &model.Team{
		ID:     string(team.ID),
		Name:   team.Name,
		Tag:    team.Tag,
		Flag:   team.Flag,
		Logo:   team.Logo,
		Public: team.Public,
	}, nil
}

// GetTeamsByUser is the resolver for the getTeamsByUser field.
func (r *queryResolver) GetTeamsByUser(ctx context.Context) ([]*model.Team, error) {
	token, err := g5ctx.GetUserToken(ctx)
	if err != nil {
		return nil, err
	}
	teams, err := r.TeamUsecase.GetTeamsByUser(ctx, token.UserID)
	if err != nil {
		return nil, err
	}

	ret := make([]*model.Team, 0, len(teams))
	for _, team := range teams {
		ret = append(ret, &model.Team{
			ID:     string(team.ID),
			Name:   team.Name,
			Tag:    team.Tag,
			Flag:   team.Flag,
			Logo:   team.Logo,
			Public: team.Public,
		})
	}
	return ret, nil
}

// GetMatch is the resolver for the getMatch field.
func (r *queryResolver) GetMatch(ctx context.Context, id string) (*model.Match, error) {
	match, err := r.MatchUsecase.GetMatch(ctx, entity.MatchID(id))
	if err != nil {
		return nil, err
	}

	team1players := make([]*model.Player, 0, len(match.Team1.Players))
	for _, player := range match.Team1.Players {
		team1players = append(team1players, &model.Player{
			ID:      string(player.ID),
			TeamID:  string(player.TeamID),
			SteamID: strconv.Itoa(int(player.SteamID)),
			Name:    player.Name,
		})
	}

	team2players := make([]*model.Player, 0, len(match.Team2.Players))
	for _, player := range match.Team2.Players {
		team2players = append(team2players, &model.Player{
			ID:      string(player.ID),
			TeamID:  string(player.TeamID),
			SteamID: strconv.Itoa(int(player.SteamID)),
			Name:    player.Name,
		})
	}

	mapstats := make([]*model.MapStats, 0, len(match.Mapstats))
	for _, mapstat := range match.Mapstats {
		playerStats := make([]*model.PlayerStats, 0, len(mapstat.PlayerStats))
		for _, playerStat := range mapstat.PlayerStats {
			playerStats = append(playerStats, &model.PlayerStats{
				ID:               string(playerStat.ID),
				MatchID:          string(playerStat.MatchID),
				MapstatsID:       string(playerStat.MapID),
				SteamID:          strconv.Itoa(int(playerStat.SteamID)),
				Name:             playerStat.Name,
				Kills:            int(playerStat.Kills),
				Assists:          int(playerStat.Assists),
				Deaths:           int(playerStat.Deaths),
				RoundsPlayed:     int(playerStat.RoundsPlayed),
				FlashBangAssists: int(playerStat.FlashbangAssists),
				Suicides:         int(playerStat.Suicides),
				HeadshotKills:    int(playerStat.HeadShotKills),
				Damage:           int(playerStat.Damage),
				BombPlants:       int(playerStat.BombPlants),
				BombDefuses:      int(playerStat.BombDefuses),
				V1:               int(playerStat.V1),
				V2:               int(playerStat.V2),
				V3:               int(playerStat.V3),
				V4:               int(playerStat.V4),
				V5:               int(playerStat.V5),
				K1:               int(playerStat.K1),
				K2:               int(playerStat.K2),
				K3:               int(playerStat.K3),
				K4:               int(playerStat.K4),
				K5:               int(playerStat.K5),
				FirstDeathT:      int(playerStat.FirstDeathT),
				FirstDeathCt:     int(playerStat.FirstDeathCT),
				FirstKillT:       int(playerStat.FirstKillT),
				FirstKillCt:      int(playerStat.FirstKillCT),
			})
		}

		mapstats = append(mapstats, &model.MapStats{
			ID:          string(mapstat.ID),
			MatchID:     string(mapstat.MatchID),
			MapNumber:   int(mapstat.MapNumber),
			MapName:     mapstat.MapName,
			StartedAt:   mapstat.StartTime,
			EndedAt:     mapstat.EndTime,
			Winner:      (*string)(mapstat.Winner),
			Team1score:  int(mapstat.Team1Score),
			Team2score:  int(mapstat.Team2Score),
			Playerstats: playerStats,
		})
	}

	return &model.Match{
		ID:     string(match.ID),
		UserID: string(match.UserID),
		Server: &model.GameServer{
			ID:     string(match.GameServer.ID),
			IP:     match.GameServer.Ip,
			Port:   int(match.GameServer.Port),
			Name:   match.GameServer.DisplayName,
			Public: match.GameServer.IsPublic,
		},
		Team1: &model.Team{
			ID:      string(match.Team1.ID),
			UserID:  string(match.Team1.UserID),
			Name:    match.Team1.Name,
			Flag:    match.Team1.Flag,
			Tag:     match.Team1.Tag,
			Logo:    match.Team1.Logo,
			Public:  match.Team1.Public,
			Players: team1players,
		},
		Team2: &model.Team{
			ID:      string(match.Team2.ID),
			UserID:  string(match.Team2.UserID),
			Name:    match.Team2.Name,
			Flag:    match.Team2.Flag,
			Tag:     match.Team2.Tag,
			Logo:    match.Team2.Logo,
			Public:  match.Team2.Public,
			Players: team2players,
		},
		Winner:     (*string)(match.Winner),
		StartedAt:  match.StartTime,
		EndedAt:    match.EndTime,
		MaxMaps:    int(match.MaxMaps),
		Title:      match.Title,
		SkipVeto:   match.SkipVeto,
		Team1Score: int(match.Team1Score),
		Team2Score: int(match.Team2Score),
		Forfeit:    match.Forfeit,
		MapStats:   mapstats,
	}, nil
}

// GetServer is the resolver for the getServer field.
func (r *queryResolver) GetServer(ctx context.Context, id string) (*model.GameServer, error) {
	panic(fmt.Errorf("not implemented: GetServer - getServer"))
}

// GetPublicServers is the resolver for the getPublicServers field.
func (r *queryResolver) GetPublicServers(ctx context.Context) ([]*model.GameServer, error) {
	panic(fmt.Errorf("not implemented: GetPublicServers - getPublicServers"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
