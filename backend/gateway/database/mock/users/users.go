package users

import (
	"context"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
)

type mockUsersRepositry struct {
	usersToReturnID      map[entity.UserID]*entity.User
	usersToReturnSteamID map[entity.SteamID]*entity.User
}

func NewMockUsersRepositry(
	usersToReturnID map[entity.UserID]*entity.User,
	usersToReturnSteamID map[entity.SteamID]*entity.User,
) database.UsersRepositry {
	return &mockUsersRepositry{
		usersToReturnID:      usersToReturnID,
		usersToReturnSteamID: usersToReturnSteamID,
	}
}

// CreateUser implements database.UsersRepositry.
func (mur *mockUsersRepositry) CreateUser(ctx context.Context, steamID entity.SteamID, name string, admin bool, hash string) error {
	return nil
}

// GetUser implements database.UsersRepositry.
func (mur *mockUsersRepositry) GetUser(ctx context.Context, id entity.UserID) (*entity.User, error) {
	v, ok := mur.usersToReturnID[id]
	if !ok {
		return nil, database.ErrNotFound
	}

	return v, nil
}

// GetUserBySteamID implements database.UsersRepositry.
func (mur *mockUsersRepositry) GetUserBySteamID(ctx context.Context, steamID entity.SteamID) (*entity.User, error) {
	v, ok := mur.usersToReturnSteamID[steamID]
	if !ok {
		return nil, database.ErrNotFound
	}

	return v, nil
}
