package graph

import (
	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/graph/model"
)

// entity をmodelに変換する処理

func convertGameServer(gs *entity.GameServer) *model.GameServer {
	return &model.GameServer{
		ID:     string(gs.ID),
		IP:     gs.Ip,
		Port:   int(gs.Port),
		Name:   gs.DisplayName,
		Public: gs.IsPublic,
	}
}

func convertGameServers(gss []*entity.GameServer) []*model.GameServer {
	ret := make([]*model.GameServer, 0, len(gss))
	for _, gs := range gss {
		ret = append(ret, convertGameServer(gs))
	}
	return ret
}

func convertTeam(team *entity.Team) *model.Team {
	return &model.Team{
		ID:     string(team.ID),
		UserID: "",
		Name:   team.Name,
		Flag:   team.Flag,
		Tag:    team.Tag,
		Logo:   team.Logo,
		Public: team.Public,
	}
}

func convertTeams(teams []*entity.Team) []*model.Team {
	ret := make([]*model.Team, 0, len(teams))
	for _, team := range teams {
		ret = append(ret, convertTeam(team))
	}
	return ret
}

func convertPlayer(player *entity.Player) *model.Player {
	return &model.Player{
		ID:      string(player.ID),
		TeamID:  string(player.TeamID),
		SteamID: uint64(player.SteamID),
		Name:    player.Name,
	}
}

func convertPlayers(players []*entity.Player) []*model.Player {
	ret := make([]*model.Player, 0, len(players))
	for _, player := range players {
		ret = append(ret, convertPlayer(player))
	}
	return ret
}

func convertMatch(match *entity.Match) *model.Match {
	return &model.Match{
		ID:         string(match.ID),
		UserID:     string(match.UserID),
		Team1:      &model.Team{},
		Team1Id:    string(match.Team1ID),
		Team2:      &model.Team{},
		Team2Id:    string(match.Team2ID),
		Winner:     string(match.Winner),
		StartedAt:  match.StartTime,
		EndedAt:    match.EndTime,
		MaxMaps:    int(match.MaxMaps),
		Title:      match.Title,
		SkipVeto:   match.SkipVeto,
		Team1Score: int(match.Team1Score),
		Team2Score: int(match.Team2Score),
		Forfeit:    match.Forfeit,
		MapStats:   []*model.MapStats{},
	}
}

func convertMatches(matches []*entity.Match) []*model.Match {
	ret := make([]*model.Match, 0, len(matches))
	for _, match := range matches {
		ret = append(ret, convertMatch(match))
	}
	return ret
}

func convertMapstat(mapstat *entity.MapStat) *model.MapStats {
	return &model.MapStats{
		ID:         string(mapstat.ID),
		MatchID:    string(mapstat.MatchID),
		MapNumber:  int(mapstat.MapNumber),
		MapName:    mapstat.MapName,
		StartedAt:  mapstat.StartTime,
		EndedAt:    mapstat.EndTime,
		Winner:     (*string)(mapstat.Winner),
		Team1score: int(mapstat.Team1Score),
		Team2score: int(mapstat.Team2Score),
	}
}

func convertMapStats(mapstats []*entity.MapStat) []*model.MapStats {
	ret := make([]*model.MapStats, 0, len(mapstats))
	for _, mapstat := range mapstats {
		ret = append(ret, convertMapstat(mapstat))
	}
	return ret
}

func convertPlayerstat(playerstat *entity.PlayerStat) *model.PlayerStats {
	return &model.PlayerStats{
		ID:               string(playerstat.ID),
		MatchID:          string(playerstat.MatchID),
		MapstatsID:       string(playerstat.MapID),
		SteamID:          uint64(playerstat.SteamID),
		Name:             playerstat.Name,
		Kills:            int(playerstat.Kills),
		Assists:          int(playerstat.Assists),
		Deaths:           int(playerstat.Deaths),
		RoundsPlayed:     int(playerstat.RoundsPlayed),
		FlashBangAssists: int(playerstat.FlashbangAssists),
		Suicides:         int(playerstat.Suicides),
		HeadshotKills:    int(playerstat.HeadShotKills),
		Damage:           int(playerstat.Damage),
		BombPlants:       int(playerstat.BombPlants),
		BombDefuses:      int(playerstat.BombDefuses),
		V1:               int(playerstat.V1),
		V2:               int(playerstat.V2),
		V3:               int(playerstat.V3),
		V4:               int(playerstat.V4),
		V5:               int(playerstat.V5),
		K1:               int(playerstat.K1),
		K2:               int(playerstat.K2),
		K3:               int(playerstat.K3),
		K4:               int(playerstat.K4),
		K5:               int(playerstat.K5),
		FirstDeathT:      int(playerstat.FirstDeathT),
		FirstDeathCt:     int(playerstat.FirstDeathCT),
		FirstKillT:       int(playerstat.FirstKillT),
		FirstKillCt:      int(playerstat.FirstKillCT),
	}
}

func convertPlayerstats(playerstats []*entity.PlayerStat) []*model.PlayerStats {
	ret := make([]*model.PlayerStats, 0, len(playerstats))
	for _, playerstat := range playerstats {
		ret = append(ret, convertPlayerstat(playerstat))
	}
	return ret
}

func convertUser(user *entity.User) *model.User {
	return &model.User{
		ID:      string(user.ID),
		SteamID: uint64(user.SteamID),
		Name:    user.Name,
		Admin:   false,
	}
}
