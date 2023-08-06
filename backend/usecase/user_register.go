package usecase

import (
	"context"

	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserRegister interface {
	// Handle registers a user and issue jwt token.
	RegisterUser(ctx context.Context, steamID string, name string, admin bool, password string) (jwt string, err error)
}

type userRegister struct {
	key                 []byte
	repositoryConnector database.RepositoryConnector
}

func NewUserRegister(
	key []byte,
	repositoryConnector database.RepositoryConnector,
) UserRegister {
	return &userRegister{
		key:                 key,
		repositoryConnector: repositoryConnector,
	}
}

// Handle implements UserRegister.
func (ur *userRegister) RegisterUser(ctx context.Context, steamID string, name string, admin bool, password string) (string, error) {
	if err := ur.repositoryConnector.Open(); err != nil {
		return "", err
	}
	defer ur.repositoryConnector.Close()

	repository := ur.repositoryConnector.GetUserRepository()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user, err := repository.CreateUser(ctx, steamID, name, admin, string(hash))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"steamID": user.SteamID,
		"admin":   user.Admin,
	})

	signed, err := token.SignedString(ur.key)
	if err != nil {
		return "", err
	}

	return signed, nil
}
