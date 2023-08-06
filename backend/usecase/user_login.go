package usecase

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/users"
)

type UserLogin interface {
	// HandleLoginSteamID returns jwt token if login is successful.
	IssueJWTBySteamID(ctx context.Context, steamID string, password string) (jwt string, err error)
}

type userLogin struct {
	key            []byte
	mysqlConnector database.DBConnector
}

func NewUserLogin(
	mysqlConnector database.DBConnector,
) UserLogin {
	return &userLogin{
		mysqlConnector: mysqlConnector,
	}
}

// HandleLoginSteamID implements UserLoginUsecase.
func (ul *userLogin) IssueJWTBySteamID(ctx context.Context, steamID string, password string) (string, error) {
	// TODO: エラーハンドリング
	db, err := ul.mysqlConnector.Connect()
	if err != nil {
		return "", err
	}
	defer db.Close()

	repository := users.NewUsersRepositry(db)
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
