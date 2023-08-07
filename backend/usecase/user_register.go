package usecase

import (
	"context"
	"errors"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	"github.com/FlowingSPDG/get5-web-go/backend/service/jwt"
	hash "github.com/FlowingSPDG/get5-web-go/backend/service/password_hash"
)

type UserRegister interface {
	// Handle registers a user and issue jwt token.
	RegisterUser(ctx context.Context, steamID entity.SteamID, name string, admin bool, password string) (jwt string, err error)
}

type userRegister struct {
	jwtService          jwt.JWTService
	passwordHasher      hash.PasswordHasher
	repositoryConnector database.RepositoryConnector
}

func NewUserRegister(
	jwtService jwt.JWTService,
	passwordHasher hash.PasswordHasher,
	repositoryConnector database.RepositoryConnector,
) UserRegister {
	return &userRegister{
		jwtService:          jwtService,
		passwordHasher:      passwordHasher,
		repositoryConnector: repositoryConnector,
	}
}

// Handle implements UserRegister.
func (ur *userRegister) RegisterUser(ctx context.Context, steamID entity.SteamID, name string, admin bool, password string) (string, error) {
	if err := ur.repositoryConnector.Open(); err != nil {
		return "", err
	}
	defer ur.repositoryConnector.Close()

	repository := ur.repositoryConnector.GetUserRepository()
	if _, err := repository.GetUserBySteamID(ctx, steamID); err == nil {
		return "", errors.New("user already exists")
	} else {
		if database.IsInternal(err) {
			return "", err
		}
	}

	hash, err := ur.passwordHasher.Hash(password)
	if err != nil {
		return "", err
	}

	if err := repository.CreateUser(ctx, steamID, name, admin, hash); err != nil {
		return "", err
	}

	user, err := repository.GetUserBySteamID(ctx, steamID)
	if err != nil {
		return "", err
	}

	signed, err := ur.jwtService.IssueJWT(user)
	if err != nil {
		return "", err
	}

	return signed, nil
}
