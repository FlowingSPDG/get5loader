package teams

import (
	"context"
	"database/sql"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	teams_gen "github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/teams/generated"
)

type teamsRepository struct {
	dsn string
}

func NewPlayersRepository(dsn string) database.TeamsRepository {
	return &teamsRepository{
		dsn: dsn,
	}
}

func (pr *teamsRepository) open() (*sql.DB, error) {
	return sql.Open("mysql", pr.dsn)
}

// AddTeam implements database.TeamsRepository.
func (tr *teamsRepository) AddTeam(ctx context.Context, userID int64, name string, tag string, flag string, logo string) (*entity.Team, error) {
	db, err := tr.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	queries := teams_gen.New(db)

	res, err := queries.AddTeam(ctx, teams_gen.AddTeamParams{
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

	team, err := queries.GetTeam(ctx, insertedID)
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
	db, err := tr.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	queries := teams_gen.New(db)

	teams, err := queries.GetPublicTeams(ctx)
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
	db, err := tr.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	queries := teams_gen.New(db)

	team, err := queries.GetTeam(ctx, id)
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
	db, err := tr.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	queries := teams_gen.New(db)

	teams, err := queries.GetTeamByUserID(ctx, userID)
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
