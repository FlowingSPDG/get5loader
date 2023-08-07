package api

// 受け取ったデータを特定の形式に変換して返すpresenter

import (
	"time"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
)

type player struct {
	ID      entity.PlayerID `json:"id"`
	TeamID  entity.TeamID   `json:"team_id"`
	SteamID entity.SteamID  `json:"steam_id"`
	Name    string          `json:"name"`
}

type team struct {
	ID         entity.TeamID `json:"id"`
	UserID     entity.UserID `son:"user_id"`
	Name       string        `son:"name"`
	Tag        string        `son:"tag"`
	Flag       string        `son:"flag"`
	Logo       string        `son:"logo"`
	Players    []player      `json:"players"`
	PublicTeam bool          `son:"public_team"`
}

type mapStats struct {
	ID         entity.MapStatsID `json:"id"`
	MatchID    entity.MatchID    `json:"match_id"`
	MapNumber  int               `json:"map_number"`
	MapName    string            `json:"map_name"`
	StartTime  time.Time         `json:"-"`
	EndTime    time.Time         `json:"-"`
	Winner     string            `json:"winner"`
	Team1Score int               `json:"team1_score"`
	Team2Score int               `json:"team2_score"`
}

type match struct {
	ID         entity.MatchID      `json:"id"`
	UserID     entity.UserID       `json:"user_id"`
	Team1ID    entity.TeamID       `json:"team1"`
	Team2ID    entity.TeamID       `json:"team2"`
	Winner     *entity.TeamID      `json:"winner"`
	ServerID   entity.GameServerID `json:"server_id"`
	Cancelled  bool                `json:"cancelled"`
	StartTime  *time.Time          `json:"start_time"`
	EndTime    *time.Time          `json:"end_time"`
	MaxMaps    int                 `json:"max_maps"`
	Title      string              `json:"title"`
	SkipVeto   bool                `json:"skip_veto"`
	Team1Score int                 `json:"team1_score"`
	Team2Score int                 `json:"team2_score"`
	Forfeit    bool                `json:"forfeit"`

	MapStatsIDs []entity.MapStatsID `json:"map_stats"`

	Status string `json:"status"`
}
