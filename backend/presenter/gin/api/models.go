package api

import (
	"time"
)

type player struct {
	ID      int64  `json:"id"`
	TeamID  int64  `json:"team_id"`
	SteamID string `json:"steam_id"`
	Name    string `json:"name"`
}

type team struct {
	ID         int      `json:"id"`
	UserID     int      `son:"user_id"`
	Name       string   `son:"name"`
	Tag        string   `son:"tag"`
	Flag       string   `son:"flag"`
	Logo       string   `son:"logo"`
	Players    []player `json:"players"`
	PublicTeam bool     `son:"public_team"`
}

type mapStats struct {
	ID         int       `json:"id"`
	MatchID    int       `json:"match_id"`
	MapNumber  int       `json:"map_number"`
	MapName    string    `json:"map_name"`
	StartTime  time.Time `json:"-"`
	EndTime    time.Time `json:"-"`
	Winner     int       `json:"winner"`
	Team1Score int       `json:"team1_score"`
	Team2Score int       `json:"team2_score"`
}

type match struct {
	ID         int        `json:"id"`
	UserID     int        `json:"user_id"`
	Team1ID    int        `json:"team1"`
	Team2ID    int        `json:"team2"`
	Winner     int        `json:"winner"`
	Cancelled  bool       `json:"cancelled"`
	StartTime  *time.Time `json:"start_time"`
	EndTime    *time.Time `json:"end_time"`
	MaxMaps    int        `json:"max_maps"`
	Title      string     `json:"title"`
	SkipVeto   bool       `json:"skip_veto"`
	Team1Score int        `json:"team1_score"`
	Team2Score int        `json:"team2_score"`
	Forfeit    bool       `json:"forfeit"`
	ServerID   int        `json:"server_id"`

	MapStatsIDs []int `json:"map_stats"`

	Status string `json:"status"`
}
