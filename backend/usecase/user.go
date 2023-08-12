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

func (u *user) GetUser(ctx context.Context, id entity.UserID) (*entity.User, error) {
	if err := u.repositoryConnector.Open(); err != nil {
		return nil, err
	}
	defer u.repositoryConnector.Close()

	userRepository := u.repositoryConnector.GetUserRepository()
	teamsRepository := u.repositoryConnector.GetTeamsRepository()
	gameServersRepository := u.repositoryConnector.GetGameServersRepository()
	matchesRepository := u.repositoryConnector.GetMatchesRepository()
	playersRepository := u.repositoryConnector.GetPlayersRepository()
	MapStatRepository := u.repositoryConnector.GetMapStatRepository()
	PlayerStatRepository := u.repositoryConnector.GetPlayerStatRepository()

	// ユーザーを取得
	user, err := userRepository.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	// チームを取得
	teams, err := teamsRepository.GetTeamsByUser(ctx, id)
	if err != nil {
		return nil, err
	}

	// チームのプレイヤーを取得
	players := make(map[entity.TeamID][]*database.Player)
	for _, team := range teams {
		playerForTeam, err := playersRepository.GetPlayersByTeam(ctx, team.ID)
		if err != nil {
			return nil, err
		}
		players[team.ID] = playerForTeam
	}

	// 変換処理
	ts := convertTeams(teams, players)

	// ゲームサーバーの取得処理
	gameServer, err := gameServersRepository.GetGameServersByUser(ctx, id)
	if err != nil {
		return nil, err
	}
	gs := convertGameServers(gameServer)

	// マッチの取得処理
	matches, err := matchesRepository.GetMatchesByUser(ctx, id)
	if err != nil {
		return nil, err
	}
	ms := make([]*entity.Match, 0, len(matches))
	for _, match := range matches {
		// Team1の取得と変換
		team1, err := teamsRepository.GetTeam(ctx, match.Team1ID)
		if err != nil {
			return nil, err
		}
		team1players, err := playersRepository.GetPlayersByTeam(ctx, team1.ID)
		if err != nil {
			return nil, err
		}

		// Team2の取得と変換
		team2, err := teamsRepository.GetTeam(ctx, match.Team2ID)
		if err != nil {
			return nil, err
		}
		team2players, err := playersRepository.GetPlayersByTeam(ctx, team2.ID)
		if err != nil {
			return nil, err
		}

		// mapstatsの取得
		mapstats, err := MapStatRepository.GetMapStatsByMatch(ctx, match.ID)
		if err != nil {
			return nil, err
		}
		maps := make([]*entity.MapStat, 0, len(mapstats))
		for _, mapstat := range mapstats {
			playerStats, err := PlayerStatRepository.GetPlayerStatsByMapstats(ctx, mapstat.ID)
			if err != nil {
				return nil, err
			}
			maps = append(maps, convertMapstat(mapstat, playerStats))
		}

		ms = append(ms, convertMatch(match, team1, team2, team1players, team2players, maps))
	}
	return convertUser(user, ts, gs, ms), nil
}

// Handle implements UserRegister.
func (u *user) Register(ctx context.Context, steamID entity.SteamID, name string, admin bool, password string) (string, error) {
	if err := u.repositoryConnector.Open(); err != nil {
		return "", err
	}
	defer u.repositoryConnector.Close()

	repository := u.repositoryConnector.GetUserRepository()
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

	signed, err := u.jwtService.IssueJWT(user.ID, user.SteamID, user.Admin)
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

	signed, err := u.jwtService.IssueJWT(user.ID, user.SteamID, user.Admin)
	if err != nil {
		return "", err
	}

	return signed, nil
}
