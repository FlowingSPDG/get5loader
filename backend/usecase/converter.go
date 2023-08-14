package usecase

import (
	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
)

// database系の構造体をentityに復元する

func convertGameServer(gs *database.GameServer) *entity.GameServer {
	return &entity.GameServer{
		UserID:       gs.UserID,
		ID:           gs.ID,
		Ip:           gs.Ip,
		Port:         gs.Port,
		RCONPassword: gs.RCONPassword,
		DisplayName:  gs.DisplayName,
		IsPublic:     gs.IsPublic,
		Status:       gs.Status,
	}
}

func convertGameServers(gss []*database.GameServer) []*entity.GameServer {
	ret := make([]*entity.GameServer, 0, len(gss))
	for _, gs := range gss {
		ret = append(ret, convertGameServer(gs))
	}
	return ret
}
func convertMapstat(ms *database.MapStat, playerstats []*database.PlayerStat) *entity.MapStat {
	return &entity.MapStat{
		ID:         ms.ID,
		MatchID:    ms.MatchID,
		MapNumber:  ms.MapNumber,
		MapName:    ms.MapName,
		StartTime:  ms.StartTime,
		EndTime:    ms.EndTime,
		Winner:     ms.Winner,
		Team1Score: ms.Team1Score,
		Team2Score: ms.Team2Score,
	}
}

func convertMapstats(mss []*database.MapStat, pss [][]*database.PlayerStat) []*entity.MapStat {
	ret := make([]*entity.MapStat, 0, len(mss))
	for i, ms := range mss {
		ret = append(ret, convertMapstat(ms, pss[i]))
	}
	return ret
}

func convertPlayerStat(ps *database.PlayerStat) *entity.PlayerStat {
	return &entity.PlayerStat{
		ID:               ps.ID,
		MatchID:          ps.MatchID,
		MapID:            ps.MapID,
		TeamID:           ps.TeamID,
		SteamID:          ps.SteamID,
		Name:             ps.Name,
		Kills:            ps.Kills,
		Assists:          ps.Assists,
		Deaths:           ps.Deaths,
		RoundsPlayed:     ps.RoundsPlayed,
		FlashbangAssists: ps.FlashbangAssists,
		Suicides:         ps.Suicides,
		HeadShotKills:    ps.HeadShotKills,
		Damage:           ps.Damage,
		BombPlants:       ps.BombPlants,
		BombDefuses:      ps.BombDefuses,
		V1:               ps.V1,
		V2:               ps.V2,
		V3:               ps.V3,
		V4:               ps.V4,
		V5:               ps.V5,
		K1:               ps.K1,
		K2:               ps.K2,
		K3:               ps.K3,
		K4:               ps.K4,
		K5:               ps.K5,
		FirstDeathCT:     ps.FirstDeathCT,
		FirstDeathT:      ps.FirstDeathT,
		FirstKillCT:      ps.FirstKillCT,
		FirstKillT:       ps.FirstKillT,
	}
}

func convertPlayerStats(pss []*database.PlayerStat) []*entity.PlayerStat {
	ret := make([]*entity.PlayerStat, 0, len(pss))
	for _, ps := range pss {
		ret = append(ret, convertPlayerStat(ps))
	}
	return ret
}

func convertTeam(t *database.Team, players []*database.Player) *entity.Team {
	return &entity.Team{
		ID:     t.ID,
		UserID: t.UserID,
		Name:   t.Name,
		Flag:   t.Flag,
		Tag:    t.Tag,
		Logo:   t.Logo,
		Public: t.Public,
	}
}

func convertTeams(ts []*database.Team, pss map[entity.TeamID][]*database.Player) []*entity.Team {
	ret := make([]*entity.Team, 0, len(ts))
	for _, t := range ts {
		ret = append(ret, convertTeam(t, pss[t.ID]))
	}
	return ret
}

func convertPlayer(p *database.Player) *entity.Player {
	return &entity.Player{
		ID:      p.ID,
		TeamID:  p.TeamID,
		SteamID: p.SteamID,
		Name:    p.Name,
	}
}

func convertPlayers(ps []*database.Player) []*entity.Player {
	ret := make([]*entity.Player, 0, len(ps))
	for _, p := range ps {
		ret = append(ret, convertPlayer(p))
	}
	return ret
}

func convertMatch(m *database.Match) *entity.Match {
	return &entity.Match{
		ID:         m.ID,
		UserID:     m.UserID,
		Team1ID:    m.Team1ID,
		Team2ID:    m.Team2ID,
		Winner:     m.Winner,
		StartTime:  m.StartTime,
		EndTime:    m.EndTime,
		MaxMaps:    m.MaxMaps,
		Title:      m.Title,
		SkipVeto:   m.SkipVeto,
		APIKey:     m.APIKey,
		Team1Score: m.Team1Score,
		Team2Score: m.Team2Score,
		Forfeit:    m.Forfeit,
		Status:     m.Status,
	}
}

func convertMatches(ms []*database.Match) []*entity.Match {
	ret := make([]*entity.Match, 0, len(ms))
	for _, m := range ms {
		ret = append(ret, convertMatch(m))
	}
	return ret
}

func convertGet5Team(t *database.Team, players []*database.Player) *entity.Get5Team {
	return &entity.Get5Team{
		ID:      t.ID,
		Name:    t.Name,
		Flag:    t.Flag,
		Tag:     t.Tag,
		Logo:    t.Logo,
		Players: convertPlayers(players),
	}
}

func convertGet5Match(m *database.Match, team1, team2 *database.Team, team1players, team2players []*database.Player) *entity.Get5Match {
	return &entity.Get5Match{
		ID:       m.ID,
		Team1:    *convertGet5Team(team1, team1players),
		Team2:    *convertGet5Team(team2, team2players),
		Winner:   m.Winner,
		MaxMaps:  m.MaxMaps,
		Title:    m.Title,
		SkipVeto: m.SkipVeto,
		APIKey:   m.APIKey,
	}
}

func convertUser(u *database.User) *entity.User {
	return &entity.User{
		ID:        u.ID,
		SteamID:   u.SteamID,
		Name:      u.Name,
		Admin:     u.Admin,
		Hash:      u.Hash,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
