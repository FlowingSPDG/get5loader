package users

import (
	"context"
	"database/sql"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	users_gen "github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/users/generated"
)

type usersRepositry struct {
	dsn string
}

func NewUsersRepositry(dsn string) database.UsersRepositry {
	return &usersRepositry{
		dsn: dsn,
	}
}

func (ur *usersRepositry) open() (*sql.DB, error) {
	return sql.Open("mysql", ur.dsn)
}

// CreateUser implements database.UsersRepositry.
func (ur *usersRepositry) CreateUser(ctx context.Context, steamID string, name string, admin bool) (*entity.User, error) {
	db, err := ur.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	queries := users_gen.New(db)

	result, err := queries.CreateUser(ctx, users_gen.CreateUserParams{
		SteamID: steamID,
		Name:    name,
		Admin:   admin,
	})
	if err != nil {
		return nil, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user, err := queries.GetUser(ctx, insertedID)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:      user.ID,
		SteamID: user.SteamID,
		Name:    user.Name,
		Admin:   user.Admin,
	}, nil

}

// GetUser implements database.UsersRepositry.
func (ur *usersRepositry) GetUser(ctx context.Context, id int64) (*entity.User, error) {
	db, err := ur.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	queries := users_gen.New(db)

	user, err := queries.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID:      user.ID,
		SteamID: user.SteamID,
		Name:    user.Name,
		Admin:   user.Admin,
	}, nil
}
