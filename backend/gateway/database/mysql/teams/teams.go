package teams

import (
	"context"
	"database/sql"
	"errors"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
	teams_gen "github.com/FlowingSPDG/get5loader/backend/gateway/database/mysql/teams/generated"
	"github.com/FlowingSPDG/get5loader/backend/service/uuid"
)

type teamsRepository struct {
	uuidGenerator uuid.UUIDGenerator
	queries       *teams_gen.Queries
}

func NewTeamsRepository(uuidGenerator uuid.UUIDGenerator, db *sql.DB) database.TeamsRepository {
	queries := teams_gen.New(db)
	return &teamsRepository{
		uuidGenerator: uuidGenerator,
		queries:       queries,
	}
}

func NewTeamsRepositoryWithTx(uuidGenerator uuid.UUIDGenerator, db *sql.DB, tx *sql.Tx) database.TeamsRepository {
	queries := teams_gen.New(db).WithTx(tx)
	return &teamsRepository{
		uuidGenerator: uuidGenerator,
		queries:       queries,
	}
}

// AddTeam implements database.TeamsRepository.
func (tr *teamsRepository) AddTeam(ctx context.Context, userID entity.UserID, name string, tag string, flag string, logo string, public bool) (entity.TeamID, error) {
	id := tr.uuidGenerator.Generate()
	if _, err := tr.queries.AddTeam(ctx, teams_gen.AddTeamParams{
		ID:     id,
		UserID: string(userID),
		Name:   name,
		Tag:    tag,
		Flag:   flag,
		Logo:   logo,
		PublicTeam: sql.NullBool{
			Bool:  public,
			Valid: true,
		},
	}); err != nil {
		return "", database.NewInternalError(err)
	}

	return entity.TeamID(id), nil
}

// GetPublicTeams implements database.TeamsRepository.
func (tr *teamsRepository) GetPublicTeams(ctx context.Context) ([]*database.Team, error) {
	teams, err := tr.queries.GetPublicTeams(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []*database.Team{}, nil
		}
		return nil, database.NewInternalError(err)
	}

	ret := make([]*database.Team, 0, len(teams))
	for _, team := range teams {
		ret = append(ret, &database.Team{
			ID:     entity.TeamID(team.ID),
			Name:   team.Name,
			Tag:    team.Tag,
			Flag:   team.Flag,
			Logo:   team.Logo,
			Public: team.PublicTeam.Valid && team.PublicTeam.Bool,
		})
	}

	return ret, nil
}

// GetTeam implements database.TeamsRepository.
func (tr *teamsRepository) GetTeam(ctx context.Context, id entity.TeamID) (*database.Team, error) {
	team, err := tr.queries.GetTeam(ctx, string(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.NewNotFoundError(err)
		}
		return nil, database.NewInternalError(err)
	}

	return &database.Team{
		ID:     entity.TeamID(team.ID),
		UserID: entity.UserID(team.UserID),
		Name:   team.Name,
		Tag:    team.Tag,
		Flag:   team.Flag,
		Logo:   team.Logo,
		Public: team.PublicTeam.Valid && team.PublicTeam.Bool,
	}, nil
}

// GetTeamsByUser implements database.TeamsRepository.
func (tr *teamsRepository) GetTeamsByUser(ctx context.Context, userID entity.UserID) ([]*database.Team, error) {
	teams, err := tr.queries.GetTeamByUserID(ctx, string(userID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []*database.Team{}, nil
		}
		return nil, database.NewInternalError(err)
	}

	ret := make([]*database.Team, 0, len(teams))
	for _, team := range teams {
		ret = append(ret, &database.Team{
			ID:     entity.TeamID(team.ID),
			UserID: entity.UserID(team.UserID),
			Name:   team.Name,
			Tag:    team.Tag,
			Flag:   team.Flag,
			Logo:   team.Logo,
			Public: team.PublicTeam.Valid && team.PublicTeam.Bool,
		})
	}

	return ret, nil
}
