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
	mapStatsRepository := u.repositoryConnector.GetMapStatsRepository()
	playerStatsRepository := u.repositoryConnector.GetPlayerStatsRepository()

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
	players := make(map[entity.TeamID][]*entity.Player)
	for _, team := range teams {
		playerForTeam, err := playersRepository.GetPlayersByTeam(ctx, team.ID)
		if err != nil {
			return nil, err
		}

		ps := make([]*entity.Player, 0, len(playerForTeam))
		for _, player := range playerForTeam {
			ps = append(ps, &entity.Player{
				ID:      entity.PlayerID(player.ID),
				TeamID:  entity.TeamID(player.TeamID),
				SteamID: player.SteamID,
				Name:    player.Name,
			})
		}
		players[team.ID] = ps
	}

	// 変換処理
	ts := make([]*entity.Team, 0, len(teams))
	for _, team := range teams {
		ts = append(ts, &entity.Team{
			ID:      entity.TeamID(team.ID),
			UserID:  entity.UserID(team.UserID),
			Name:    team.Name,
			Flag:    team.Flag,
			Tag:     team.Tag,
			Logo:    team.Logo,
			Public:  team.Public,
			Players: players[team.ID],
		})
	}

	// ゲームサーバーの取得処理
	gameServer, err := gameServersRepository.GetGameServersByUser(ctx, id)
	if err != nil {
		return nil, err
	}
	gs := make([]*entity.GameServer, 0, len(gameServer))
	for _, server := range gameServer {
		gs = append(gs, &entity.GameServer{
			UserID:       entity.UserID(server.UserID),
			ID:           entity.GameServerID(server.ID),
			Ip:           server.Ip,
			Port:         server.Port,
			RCONPassword: server.RCONPassword,
			DisplayName:  server.DisplayName,
			IsPublic:     server.IsPublic,
			Status:       server.Status,
		})
	}

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
		t1p := make([]*entity.Player, 0, len(team1players))
		for _, player := range team1players {
			t1p = append(t1p, &entity.Player{
				ID:      player.ID,
				TeamID:  player.TeamID,
				SteamID: player.SteamID,
				Name:    player.Name,
			})
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
		t2p := make([]*entity.Player, 0, len(team2players))
		for _, player := range team2players {
			t2p = append(t2p, &entity.Player{
				ID:      player.ID,
				TeamID:  player.TeamID,
				SteamID: player.SteamID,
				Name:    player.Name,
			})
		}

		// mapstatsの取得
		mapstats, err := mapStatsRepository.GetMapStatsByMatch(ctx, match.ID)
		if err != nil {
			return nil, err
		}
		maps := make([]*entity.MapStats, 0, len(mapstats))
		for _, mapstat := range mapstats {
			playerStats, err := playerStatsRepository.GetPlayerStatsByMapstats(ctx, mapstat.ID)
			if err != nil {
				return nil, err
			}

			ps := make([]*entity.PlayerStats, 0, len(playerStats))
			for _, playerStat := range playerStats {
				ps = append(ps, &entity.PlayerStats{
					ID:               playerStat.ID,
					MatchID:          playerStat.MatchID,
					MapID:            playerStat.MapID,
					TeamID:           playerStat.TeamID,
					SteamID:          playerStat.SteamID,
					Name:             playerStat.Name,
					Kills:            playerStat.Kills,
					Assists:          playerStat.Assists,
					Deaths:           playerStat.Deaths,
					RoundsPlayed:     playerStat.RoundsPlayed,
					FlashbangAssists: playerStat.FlashbangAssists,
					Suicides:         playerStat.Suicides,
					HeadShotKills:    playerStat.HeadShotKills,
					Damage:           playerStat.Damage,
					BombPlants:       playerStat.BombPlants,
					BombDefuses:      playerStat.BombDefuses,
					V1:               playerStat.V1,
					V2:               playerStat.V2,
					V3:               playerStat.V3,
					V4:               playerStat.V4,
					V5:               playerStat.V5,
					K1:               playerStat.K1,
					K2:               playerStat.K2,
					K3:               playerStat.K3,
					K4:               playerStat.K4,
					K5:               playerStat.K5,
					FirstDeathCT:     playerStat.FirstDeathCT,
					FirstDeathT:      playerStat.FirstDeathT,
					FirstKillCT:      playerStat.FirstKillCT,
					FirstKillT:       playerStat.FirstKillT,
				})
			}
			maps = append(maps, &entity.MapStats{
				ID:          mapstat.ID,
				MatchID:     mapstat.MatchID,
				MapNumber:   mapstat.MapNumber,
				MapName:     mapstat.MapName,
				StartTime:   mapstat.StartTime,
				EndTime:     mapstat.EndTime,
				Winner:      mapstat.Winner,
				Team1Score:  mapstat.Team1Score,
				Team2Score:  mapstat.Team2Score,
				PlayerStats: ps,
			})
		}

		ms = append(ms, &entity.Match{
			ID:     entity.MatchID(match.ID),
			UserID: id,
			Team1: entity.Team{
				ID:      entity.TeamID(team1.ID),
				UserID:  entity.UserID(team1.UserID),
				Name:    team1.Name,
				Flag:    team1.Flag,
				Tag:     team1.Tag,
				Logo:    team1.Logo,
				Public:  team1.Public,
				Players: t1p,
			},
			Team2: entity.Team{
				ID:      entity.TeamID(team2.ID),
				UserID:  entity.UserID(team2.UserID),
				Name:    team2.Name,
				Flag:    team2.Flag,
				Tag:     team2.Tag,
				Logo:    team2.Logo,
				Public:  team2.Public,
				Players: t2p,
			},
			Winner:     match.Winner,
			StartTime:  match.StartTime,
			EndTime:    match.EndTime,
			MaxMaps:    match.MaxMaps,
			Title:      match.Title,
			SkipVeto:   match.SkipVeto,
			APIKey:     match.APIKey,
			Team1Score: match.Team1Score,
			Team2Score: match.Team2Score,
			Forfeit:    match.Forfeit,
			Status:     match.Status,
			Mapstats:   maps,
		})
	}

	return &entity.User{
		ID:        user.ID,
		SteamID:   user.SteamID,
		Name:      user.Name,
		Admin:     user.Admin,
		Hash:      user.Hash,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Teams:     ts,
		Servers:   gs,
		Matches:   ms,
	}, nil
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
