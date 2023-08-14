package dataloaders

import (
	"github.com/graph-gophers/dataloader/v7"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/usecase"
)

type Loaders struct {
	PlayersByTeamID        dataloader.Interface[entity.TeamID, []*entity.Player]
	MapStatsByMatchID      dataloader.Interface[entity.MatchID, []*entity.MapStat]
	PlayerStatsByMapStatID dataloader.Interface[entity.MapStatsID, []*entity.PlayerStat]
	TeamsByTeamID          dataloader.Interface[entity.TeamID, *entity.Team]
	TeamByUserID           dataloader.Interface[entity.UserID, []*entity.Team]
	ServersByUserID        dataloader.Interface[entity.UserID, []*entity.GameServer]
	MatchByUserID          dataloader.Interface[entity.UserID, []*entity.Match]
}

func (l *Loaders) ClearAll() {
	l.PlayersByTeamID.ClearAll()
	l.MapStatsByMatchID.ClearAll()
	l.PlayerStatsByMapStatID.ClearAll()
	l.TeamsByTeamID.ClearAll()
	l.TeamByUserID.ClearAll()
	l.ServersByUserID.ClearAll()
	l.MatchByUserID.ClearAll()
}

func NewLoaders(
	player usecase.Player,
	match usecase.Match,
	team usecase.Team,
	mapstats usecase.Mapstat,
	playerstats usecase.PlayerStat,
	server usecase.GameServer,
) *Loaders {
	teamPlayersLoader := &teamPlayersLoader{player: player}
	matchMapstatLoader := &matchMapstat{mapstat: mapstats}
	mapStatPlayerstatsLoader := &mapstatPlayerStats{playerstat: playerstats}
	teamsLoader := &teamsLoader{team: team}
	userTeamLoader := &userTeamLoader{team: team}
	userServersLoader := &userServersLoader{server: server}
	userMatchLoader := &userMatchLoader{match: match}
	return &Loaders{
		PlayersByTeamID:        dataloader.NewBatchedLoader[entity.TeamID, []*entity.Player](teamPlayersLoader.BatchGetPlayers),
		MapStatsByMatchID:      dataloader.NewBatchedLoader[entity.MatchID, []*entity.MapStat](matchMapstatLoader.BatchGetMapStats),
		PlayerStatsByMapStatID: dataloader.NewBatchedLoader[entity.MapStatsID, []*entity.PlayerStat](mapStatPlayerstatsLoader.BatchGetMapstats),
		TeamsByTeamID:          dataloader.NewBatchedLoader[entity.TeamID, *entity.Team](teamsLoader.BatchGetTeams),
		TeamByUserID:           dataloader.NewBatchedLoader[entity.UserID, []*entity.Team](userTeamLoader.BatchGetTeams),
		ServersByUserID:        dataloader.NewBatchedLoader[entity.UserID, []*entity.GameServer](userServersLoader.BatchGetServers),
		MatchByUserID:          dataloader.NewBatchedLoader[entity.UserID, []*entity.Match](userMatchLoader.BatchGetMatches),
	}
}
