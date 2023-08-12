package usecase

import (
	"context"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	"github.com/FlowingSPDG/get5-web-go/backend/service/jwt"
	hash "github.com/FlowingSPDG/get5-web-go/backend/service/password_hash"
)

type UserLogin interface {
	// HandleLoginSteamID returns jwt token if login is successful.
	IssueJWTBySteamID(ctx context.Context, steamID entity.SteamID, password string) (jwt string, err error)
}

type userLogin struct {
	jwtService          jwt.JWTService
	passwordHasher      hash.PasswordHasher
	repositoryConnector database.RepositoryConnector
}

func NewUserLogin(
	jwtService jwt.JWTService,
	passwordHasher hash.PasswordHasher,
	repositoryConnector database.RepositoryConnector,
) UserLogin {
	return &userLogin{
		jwtService:          jwtService,
		passwordHasher:      passwordHasher,
		repositoryConnector: repositoryConnector,
	}
}

// HandleLoginSteamID implements UserLoginUsecase.
func (ul *userLogin) IssueJWTBySteamID(ctx context.Context, steamID entity.SteamID, password string) (string, error) {
	// TODO: エラーハンドリング
	if err := ul.repositoryConnector.Open(); err != nil {
		return "", err
	}
	defer ul.repositoryConnector.Close()

	repository := ul.repositoryConnector.GetUserRepository()
	user, err := repository.GetUserBySteamID(ctx, steamID)
	if err != nil {
		return "", err
	}

	if err := ul.passwordHasher.Compare(user.Hash, password); err != nil {
		return "", err
	}

	signed, err := ul.jwtService.IssueJWT(user)
	if err != nil {
		return "", err
	}

	return signed, nil
}
