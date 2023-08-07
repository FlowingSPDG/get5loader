package usecase

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	"github.com/FlowingSPDG/get5-web-go/backend/service/jwt"
)

type UserLogin interface {
	// HandleLoginSteamID returns jwt token if login is successful.
	IssueJWTBySteamID(ctx context.Context, steamID entity.SteamID, password string) (jwt string, err error)
}

type userLogin struct {
	jwtService          jwt.JWTService
	repositoryConnector database.RepositoryConnector
}

func NewUserLogin(
	jwtService jwt.JWTService,
	repositoryConnector database.RepositoryConnector,
) UserLogin {
	return &userLogin{
		jwtService:          jwtService,
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

	if err := bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password)); err != nil {
		return "", err
	}

	signed, err := ul.jwtService.IssueJWT(user)
	if err != nil {
		return "", err
	}

	return signed, nil
}
