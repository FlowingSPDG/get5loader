package users

import (
	"context"
	"database/sql"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	users_gen "github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/users/generated"
)

type usersRepositry struct {
	queries *users_gen.Queries
}

func NewUsersRepositry(db *sql.DB) database.UsersRepositry {
	queries := users_gen.New(db)
	return &usersRepositry{
		queries: queries,
	}
}

func NewUsersRepositryWithTx(db *sql.DB, tx *sql.Tx) database.UsersRepositry {
	queries := users_gen.New(db).WithTx(tx)
	return &usersRepositry{
		queries: queries,
	}
}

// CreateUser implements database.UsersRepositry.
func (ur *usersRepositry) CreateUser(ctx context.Context, steamID string, name string, admin bool, hash string) (*entity.User, error) {
	result, err := ur.queries.CreateUser(ctx, users_gen.CreateUserParams{
		SteamID:      steamID,
		Name:         name,
		Admin:        admin,
		PasswordHash: hash,
	})
	if err != nil {
		return nil, database.NewInternalError(err)
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return nil, database.NewInternalError(err)
	}

	user, err := ur.queries.GetUser(ctx, insertedID)
	if err != nil {
		return nil, database.NewInternalError(err)
	}

	return &entity.User{
		ID:        user.ID,
		SteamID:   user.SteamID,
		Name:      user.Name,
		Admin:     user.Admin,
		Hash:      user.PasswordHash,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil

}

// GetUser implements database.UsersRepositry.
func (ur *usersRepositry) GetUser(ctx context.Context, id int64) (*entity.User, error) {
	user, err := ur.queries.GetUser(ctx, id)
	if err != nil {
		return nil, database.NewInternalError(err)
	}
	return &entity.User{
		ID:        user.ID,
		SteamID:   user.SteamID,
		Name:      user.Name,
		Admin:     user.Admin,
		Hash:      user.PasswordHash,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
