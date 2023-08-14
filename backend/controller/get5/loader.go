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
	match *entity.Get5Match
}

func (m *match) ToG5Format() got5.Match {
	team1Players := map[string]string{}
	for _, player := range m.match.Team1.Players {
		team1Players[strconv.Itoa(int(player.SteamID))] = player.Name
	}
	team2Players := map[string]string{}
	for _, player := range m.match.Team2.Players {
		team2Players[strconv.Itoa(int(player.SteamID))] = player.Name
	}

	return got5.Match{
		MatchTitle:           m.match.Title,
		MatchID:              string(m.match.ID),
		ClinchSeries:         false,
		NumMaps:              int(m.match.MaxMaps),
		Scrim:                false,
		Wingman:              false,
		PlayersPerTeam:       5,
		CoachesPerTeam:       0,
		CoachesMustReady:     false,
		MinPlayersToReady:    5,
		MinSpectatorsToReady: 0,
		SkipVeto:             m.match.SkipVeto,
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
			ID:          string(m.match.Team1.ID),
			Players:     team1Players,
			Coaches:     map[string]string{},
			Name:        m.match.Team1.Name,
			Tag:         m.match.Team1.Tag,
			Flag:        m.match.Team1.Flag,
			Logo:        m.match.Team1.Logo,
			SeriesScore: 0,
			MatchText:   "",
			FromFile:    "",
		},
		Team2: got5.Team{
			ID:          string(m.match.Team2.ID),
			Players:     team2Players,
			Coaches:     map[string]string{},
			Name:        m.match.Team2.Name,
			Tag:         m.match.Team2.Tag,
			Flag:        m.match.Team2.Flag,
			Logo:        m.match.Team2.Logo,
			SeriesScore: 0,
			MatchText:   "",
			FromFile:    "",
		},
		Cvars: map[string]string{},
	}
}

// Load implements got5.MatchLoader.
func (ml *matchLoader) Load(ctx context.Context, mid string) (got5.G5Match, error) {
	uc := usecase.NewGet5()
	m, err := uc.GetMatch(ctx, entity.MatchID(mid))
	if err != nil {
		return nil, err
	}
	return &match{
		match: m,
	}, nil

}
