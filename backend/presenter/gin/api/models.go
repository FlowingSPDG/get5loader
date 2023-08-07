package api

// 受け取ったデータを特定の形式に変換して返すpresenter
// このpackageでは極力根底型を使うようにする

import (
	"time"
)

type player struct {
	ID      string `json:"id"`
	TeamID  string `json:"team_id"`
	SteamID uint64 `json:"steam_id"`
	Name    string `json:"name"`
}

type team struct {
	ID         string   `json:"id"`
	UserID     string   `son:"user_id"`
	Name       string   `son:"name"`
	Tag        string   `son:"tag"`
	Flag       string   `son:"flag"`
	Logo       string   `son:"logo"`
	Players    []player `json:"players"`
	PublicTeam bool     `son:"public_team"`
}

type mapStats struct {
	ID         string    `json:"id"`
	MatchID    string    `json:"match_id"`
	MapNumber  int       `json:"map_number"`
	MapName    string    `json:"map_name"`
	StartTime  time.Time `json:"-"`
	EndTime    time.Time `json:"-"`
	Winner     string    `json:"winner"`
	Team1Score int       `json:"team1_score"`
	Team2Score int       `json:"team2_score"`
}

type match struct {
	ID         string     `json:"id"`
	UserID     string     `json:"user_id"`
	Team1ID    string     `json:"team1"`
	Team2ID    string     `json:"team2"`
	Winner     *string    `json:"winner"`
	Cancelled  bool       `json:"cancelled"`
	StartTime  *time.Time `json:"start_time"`
	EndTime    *time.Time `json:"end_time"`
	MaxMaps    int        `json:"max_maps"`
	Title      string     `json:"title"`
	SkipVeto   bool       `json:"skip_veto"`
	Team1Score int        `json:"team1_score"`
	Team2Score int        `json:"team2_score"`
	Forfeit    bool       `json:"forfeit"`
	ServerID   string     `json:"server_id"`

	MapStatsIDs []string `json:"map_stats"`

	Status string `json:"status"`
}
