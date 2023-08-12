package users

import (
	"context"
	"database/sql"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	users_gen "github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/users/generated"
	"github.com/FlowingSPDG/get5-web-go/backend/service/uuid"
)

type usersRepositry struct {
	uuidGenerator uuid.UUIDGenerator
	queries       *users_gen.Queries
}

func NewUsersRepositry(uuidGenerator uuid.UUIDGenerator, db *sql.DB) database.UsersRepositry {
	queries := users_gen.New(db)
	return &usersRepositry{
		uuidGenerator: uuidGenerator,
		queries:       queries,
	}
}

func NewUsersRepositryWithTx(uuidGenerator uuid.UUIDGenerator, db *sql.DB, tx *sql.Tx) database.UsersRepositry {
	queries := users_gen.New(db).WithTx(tx)
	return &usersRepositry{
		uuidGenerator: uuidGenerator,
		queries:       queries,
	}
}

// CreateUser implements database.UsersRepositry.
func (ur *usersRepositry) CreateUser(ctx context.Context, steamID entity.SteamID, name string, admin bool, hash []byte) error {
	if _, err := ur.queries.CreateUser(ctx, users_gen.CreateUserParams{
		ID:           ur.uuidGenerator.Generate(),
		SteamID:      uint64(steamID),
		Name:         name,
		Admin:        admin,
		PasswordHash: hash,
	}); err != nil {
		return database.NewInternalError(err)
	}

	return nil
}

// GetUser implements database.UsersRepositry.
func (ur *usersRepositry) GetUser(ctx context.Context, id entity.UserID) (*entity.User, error) {
	user, err := ur.queries.GetUser(ctx, string(id))
	if err != nil {
		return nil, database.NewInternalError(err)
	}
	return &entity.User{
		ID:        entity.UserID(user.ID),
		SteamID:   entity.SteamID(user.SteamID),
		Name:      user.Name,
		Admin:     user.Admin,
		Hash:      user.PasswordHash,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// GetUserBySteamID implements database.UsersRepositry.
func (ur *usersRepositry) GetUserBySteamID(ctx context.Context, steamID entity.SteamID) (*entity.User, error) {
	user, err := ur.queries.GetUserBySteamID(ctx, uint64(steamID))
	if err != nil {
		return nil, database.NewInternalError(err)
	}
	return &entity.User{
		ID:        entity.UserID(user.ID),
		SteamID:   entity.SteamID(user.SteamID),
		Name:      user.Name,
		Admin:     user.Admin,
		Hash:      user.PasswordHash,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
