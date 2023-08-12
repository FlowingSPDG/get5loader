package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
	"github.com/FlowingSPDG/get5loader/backend/service/jwt"
	hash "github.com/FlowingSPDG/get5loader/backend/service/password_hash"
)

type User interface {
	Register(ctx context.Context, steamID entity.SteamID, name string, admin bool, password string) (jwt string, err error)
	IssueJWT(ctx context.Context, userID entity.UserID, password string) (jwt string, err error)
	IssueJWTBySteamID(ctx context.Context, steamID entity.SteamID, password string) (jwt string, err error)
}

type user struct {
	jwtService          jwt.JWTService
	passwordHasher      hash.PasswordHasher
	repositoryConnector database.RepositoryConnector
}

func NewUser(jwtService jwt.JWTService, passwordHasher hash.PasswordHasher, repositoryConnector database.RepositoryConnector) User {
	return &user{
		jwtService:          jwtService,
		passwordHasher:      passwordHasher,
		repositoryConnector: repositoryConnector,
	}
}

// Handle implements UserRegister.
func (u *user) Register(ctx context.Context, steamID entity.SteamID, name string, admin bool, password string) (string, error) {
	if err := u.repositoryConnector.Open(); err != nil {
		return "", err
	}
	defer u.repositoryConnector.Close()

	log.Println("registering user")

	repository := u.repositoryConnector.GetUserRepository()
	_, err := repository.GetUserBySteamID(ctx, steamID)
	if err == nil {
		return "", errors.New("user already exists")
	} else if !database.IsNotFound(err) {
		return "", err
	}

	log.Println("generating hash")
	hash, err := u.passwordHasher.Hash(password)
	if err != nil {
		return "", err
	}

	log.Println("creating hash")
	if _, err := repository.CreateUser(ctx, steamID, name, admin, hash); err != nil {
		return "", err
	}

	user, err := repository.GetUserBySteamID(ctx, steamID)
	if err != nil {
		return "", err
	}

	signed, err := u.jwtService.IssueJWT(&entity.User{
		ID:        user.ID,
		SteamID:   user.SteamID,
		Name:      user.Name,
		Admin:     user.Admin,
		Hash:      user.Hash,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
	if err != nil {
		return "", err
	}

	return signed, nil
}

// HandleLoginSteamID implements UserLoginUsecase.
func (u *user) IssueJWT(ctx context.Context, userID entity.UserID, password string) (string, error) {
	// TODO: エラーハンドリング
	if err := u.repositoryConnector.Open(); err != nil {
		return "", err
	}
	defer u.repositoryConnector.Close()

	repository := u.repositoryConnector.GetUserRepository()
	user, err := repository.GetUser(ctx, userID)
	if err != nil {
		return "", err
	}

	if err := u.passwordHasher.Compare(user.Hash, password); err != nil {
		return "", err
	}

	signed, err := u.jwtService.IssueJWT(&entity.User{
		ID:        user.ID,
		SteamID:   user.SteamID,
		Name:      user.Name,
		Admin:     user.Admin,
		Hash:      user.Hash,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
	if err != nil {
		return "", err
	}

	return signed, nil
}

// HandleLoginSteamID implements UserLoginUsecase.
func (u *user) IssueJWTBySteamID(ctx context.Context, steamID entity.SteamID, password string) (string, error) {
	// TODO: エラーハンドリング
	if err := u.repositoryConnector.Open(); err != nil {
		return "", err
	}
	defer u.repositoryConnector.Close()

	repository := u.repositoryConnector.GetUserRepository()
	user, err := repository.GetUserBySteamID(ctx, steamID)
	if err != nil {
		return "", err
	}

	if err := u.passwordHasher.Compare(user.Hash, password); err != nil {
		return "", err
	}

	signed, err := u.jwtService.IssueJWT(&entity.User{
		ID:        user.ID,
		SteamID:   user.SteamID,
		Name:      user.Name,
		Admin:     user.Admin,
		Hash:      user.Hash,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
	if err != nil {
		return "", err
	}

	return signed, nil
}
