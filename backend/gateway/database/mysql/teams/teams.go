package teams

import (
	"context"
	"database/sql"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	teams_gen "github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/teams/generated"
)

type teamsRepository struct {
	queries *teams_gen.Queries
}

func NewPlayersRepository(db *sql.DB) database.TeamsRepository {
	queries := teams_gen.New(db)
	return &teamsRepository{
		queries: queries,
	}
}

func NewPlayerRepositoryWithTx(db *sql.DB, tx *sql.Tx) database.TeamsRepository {
	queries := teams_gen.New(db).WithTx(tx)
	return &teamsRepository{
		queries: queries,
	}
}

// AddTeam implements database.TeamsRepository.
func (tr *teamsRepository) AddTeam(ctx context.Context, userID int64, name string, tag string, flag string, logo string) (*entity.Team, error) {
	res, err := tr.queries.AddTeam(ctx, teams_gen.AddTeamParams{
		UserID: userID,
		Name:   name,
		Tag:    tag,
		Flag:   flag,
		Logo:   logo,
	})
	if err != nil {
		return nil, err
	}

	insertedID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	team, err := tr.queries.GetTeam(ctx, insertedID)
	if err != nil {
		return nil, err
	}

	return &entity.Team{
		ID:     team.ID,
		UserID: team.UserID,
		Name:   team.Name,
		Tag:    team.Tag,
		Flag:   team.Flag,
		Logo:   team.Logo,
	}, nil
}

// GetPublicTeams implements database.TeamsRepository.
func (tr *teamsRepository) GetPublicTeams(ctx context.Context) ([]*entity.Team, error) {
	teams, err := tr.queries.GetPublicTeams(ctx)
	if err != nil {
		return nil, err
	}

	ret := make([]*entity.Team, 0, len(teams))
	for _, team := range teams {
		ret = append(ret, &entity.Team{
			ID:     team.ID,
			UserID: team.UserID,
			Name:   team.Name,
			Tag:    team.Tag,
			Flag:   team.Flag,
			Logo:   team.Logo,
		})
	}

	return ret, nil
}

// GetTeam implements database.TeamsRepository.
func (tr *teamsRepository) GetTeam(ctx context.Context, id int64) (*entity.Team, error) {
	team, err := tr.queries.GetTeam(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.Team{
		ID:     team.ID,
		UserID: team.UserID,
		Name:   team.Name,
		Tag:    team.Tag,
		Flag:   team.Flag,
		Logo:   team.Logo,
	}, nil
}

// GetTeamsByUser implements database.TeamsRepository.
func (tr *teamsRepository) GetTeamsByUser(ctx context.Context, userID int64) ([]*entity.Team, error) {
	teams, err := tr.queries.GetTeamByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	ret := make([]*entity.Team, 0, len(teams))
	for _, team := range teams {
		ret = append(ret, &entity.Team{
			ID:     team.ID,
			UserID: team.UserID,
			Name:   team.Name,
			Tag:    team.Tag,
			Flag:   team.Flag,
			Logo:   team.Logo,
		})
	}

	return ret, nil
}
