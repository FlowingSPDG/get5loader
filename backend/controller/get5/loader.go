package get5

import (
	"context"
	"strconv"

	got5 "github.com/FlowingSPDG/Got5"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
	"github.com/FlowingSPDG/get5loader/backend/usecase"
)

type matchLoader struct {
	repositoryConnector database.RepositoryConnector
}

func NewGot5MatchLoader(repositoryConnector database.RepositoryConnector) got5.MatchLoader {
	return &matchLoader{
		repositoryConnector: repositoryConnector,
	}
}

type match struct {
	m *entity.Match
}

func (match *match) ToG5Format() got5.Match {
	team1Players := map[string]string{}
	for _, player := range match.m.Team1.Players {
		team1Players[strconv.Itoa(int(player.SteamID))] = player.Name
	}
	team2Players := map[string]string{}
	for _, player := range match.m.Team2.Players {
		team2Players[strconv.Itoa(int(player.SteamID))] = player.Name
	}

	return got5.Match{
		MatchTitle:           match.m.Title,
		MatchID:              string(match.m.ID),
		ClinchSeries:         false,
		NumMaps:              int(match.m.MaxMaps),
		Scrim:                false,
		Wingman:              false,
		PlayersPerTeam:       5,
		CoachesPerTeam:       0,
		CoachesMustReady:     false,
		MinPlayersToReady:    5,
		MinSpectatorsToReady: 0,
		SkipVeto:             match.m.SkipVeto,
		VetoFirst:            "random",
		VetoMode:             "",
		SideType:             "standard",
		MapSides:             []string{},
		Spectators:           got5.Spectators{},
		Maplist: []string{
			"de_inferno",
			"de_mirage",
			"de_nuke",
			"de_overpass",
			"de_vertigo",
			"de_ancient",
			"de_anubis",
		},
		FavoredPercentageTeam1: 0,
		FavoredPercentageText:  "",
		Team1: got5.Team{
			ID:          string(match.m.Team1.ID),
			Players:     team1Players,
			Coaches:     map[string]string{},
			Name:        match.m.Team1.Name,
			Tag:         match.m.Team1.Tag,
			Flag:        match.m.Team1.Flag,
			Logo:        match.m.Team1.Logo,
			SeriesScore: 0,
			MatchText:   "",
		},
		Team2: got5.Team{
			ID:          string(match.m.Team2.ID),
			Players:     team2Players,
			Coaches:     map[string]string{},
			Name:        match.m.Team2.Name,
			Tag:         match.m.Team2.Tag,
			Flag:        match.m.Team2.Flag,
			Logo:        match.m.Team2.Logo,
			SeriesScore: 0,
			MatchText:   "",
		},
		Cvars: map[string]string{},
	}
}

// Load implements got5.MatchLoader.
func (ml *matchLoader) Load(ctx context.Context, mid string) (got5.G5Match, error) {
	uc := usecase.NewMatch(ml.repositoryConnector)
	m, err := uc.GetMatch(ctx, entity.MatchID(mid))
	if err != nil {
		return nil, err
	}
	return &match{
		m: m,
	}, nil

}
