package usecase

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
)

type UserLogin interface {
	// HandleLoginSteamID returns jwt token if login is successful.
	IssueJWTBySteamID(ctx context.Context, steamID string, password string) (jwt string, err error)
}

type userLogin struct {
	key                 []byte
	repositoryConnector database.RepositoryConnector
}

func NewUserLogin(
	key []byte,
	repositoryConnector database.RepositoryConnector,
) UserLogin {
	return &userLogin{
		key:                 key,
		repositoryConnector: repositoryConnector,
	}
}

// HandleLoginSteamID implements UserLoginUsecase.
func (ul *userLogin) IssueJWTBySteamID(ctx context.Context, steamID string, password string) (string, error) {
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"steamID": user.SteamID,
		"admin":   user.Admin,
	})

	signed, err := token.SignedString(ul.key)
	if err != nil {
		return "", err
	}

	return signed, nil
}
