package usecase

import (
	"context"
	"errors"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
	"github.com/FlowingSPDG/get5loader/backend/service/jwt"
	hash "github.com/FlowingSPDG/get5loader/backend/service/password_hash"
)

type User interface {
	GetUser(ctx context.Context, id entity.UserID) (*entity.User, error)
	Register(ctx context.Context, steamID entity.SteamID, name string, admin bool, password string) (jwt string, err error)
	IssueJWT(ctx context.Context, userID entity.UserID, password string) (jwt string, err error)
	IssueJWTBySteamID(ctx context.Context, steamID entity.SteamID, password string) (jwt string, err error)
}

type user struct {
	jwtService     jwt.JWTService
	passwordHasher hash.PasswordHasher
}

func NewUser(jwtService jwt.JWTService, passwordHasher hash.PasswordHasher) User {
	return &user{
		jwtService:     jwtService,
		passwordHasher: passwordHasher,
	}
}

func (u *user) GetUser(ctx context.Context, id entity.UserID) (*entity.User, error) {
	repositoryConnector := database.GetConnection(ctx)

	userRepository := repositoryConnector.GetUserRepository()

	// ユーザーを取得
	user, err := userRepository.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return convertUser(user), nil
}

// Handle implements UserRegister.
func (u *user) Register(ctx context.Context, steamID entity.SteamID, name string, admin bool, password string) (string, error) {
	repositoryConnector := database.GetConnection(ctx)

	repository := repositoryConnector.GetUserRepository()
	_, err := repository.GetUserBySteamID(ctx, steamID)
	if err == nil {
		return "", errors.New("user already exists")
	} else if !database.IsNotFound(err) {
		return "", err
	}

	hash, err := u.passwordHasher.Hash(password)
	if err != nil {
		return "", err
	}

	if _, err := repository.CreateUser(ctx, steamID, name, admin, hash); err != nil {
		return "", err
	}

	user, err := repository.GetUserBySteamID(ctx, steamID)
	if err != nil {
		return "", err
	}

	signed, err := u.jwtService.IssueJWT(user.ID, user.SteamID, user.Admin)
	if err != nil {
		return "", err
	}

	return signed, nil
}

// HandleLoginSteamID implements UserLoginUsecase.
func (u *user) IssueJWT(ctx context.Context, userID entity.UserID, password string) (string, error) {
	// TODO: エラーハンドリング
	repositoryConnector := database.GetConnection(ctx)

	repository := repositoryConnector.GetUserRepository()
	user, err := repository.GetUser(ctx, userID)
	if err != nil {
		return "", err
	}

	if err := u.passwordHasher.Compare(user.Hash, password); err != nil {
		return "", err
	}

	signed, err := u.jwtService.IssueJWT(user.ID, user.SteamID, user.Admin)
	if err != nil {
		return "", err
	}

	return signed, nil
}

// HandleLoginSteamID implements UserLoginUsecase.
func (u *user) IssueJWTBySteamID(ctx context.Context, steamID entity.SteamID, password string) (string, error) {
	// TODO: エラーハンドリング
	repositoryConnector := database.GetConnection(ctx)

	repository := repositoryConnector.GetUserRepository()
	user, err := repository.GetUserBySteamID(ctx, steamID)
	if err != nil {
		return "", err
	}

	if err := u.passwordHasher.Compare(user.Hash, password); err != nil {
		return "", err
	}

	signed, err := u.jwtService.IssueJWT(user.ID, user.SteamID, user.Admin)
	if err != nil {
		return "", err
	}

	return signed, nil
}
