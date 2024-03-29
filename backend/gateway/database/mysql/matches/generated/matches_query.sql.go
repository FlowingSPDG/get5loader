// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: matches_query.sql

package matches_gen

import (
	"context"
	"database/sql"
	"strings"
)

const addMatch = `-- name: AddMatch :execresult
INSERT INTO matches (
  id, user_id, server_id, team1_id, team2_id, start_time, end_time, max_maps, title, skip_veto, api_key, status
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
`

type AddMatchParams struct {
	ID        string
	UserID    string
	ServerID  string
	Team1ID   string
	Team2ID   string
	StartTime sql.NullTime
	EndTime   sql.NullTime
	MaxMaps   int32
	Title     string
	SkipVeto  bool
	ApiKey    string
	Status    int32
}

func (q *Queries) AddMatch(ctx context.Context, arg AddMatchParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addMatch,
		arg.ID,
		arg.UserID,
		arg.ServerID,
		arg.Team1ID,
		arg.Team2ID,
		arg.StartTime,
		arg.EndTime,
		arg.MaxMaps,
		arg.Title,
		arg.SkipVeto,
		arg.ApiKey,
		arg.Status,
	)
}

const cancelMatch = `-- name: CancelMatch :execresult
UPDATE matches
SET status = 4
WHERE id = ?
`

func (q *Queries) CancelMatch(ctx context.Context, id string) (sql.Result, error) {
	return q.db.ExecContext(ctx, cancelMatch, id)
}

const getMatch = `-- name: GetMatch :one
SELECT id, user_id, server_id, team1_id, team2_id, winner, cancelled, start_time, end_time, max_maps, title, skip_veto, api_key, team1_score, team2_score, forfeit, status FROM matches
WHERE id = ? LIMIT 1
`

func (q *Queries) GetMatch(ctx context.Context, id string) (Match, error) {
	row := q.db.QueryRowContext(ctx, getMatch, id)
	var i Match
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ServerID,
		&i.Team1ID,
		&i.Team2ID,
		&i.Winner,
		&i.Cancelled,
		&i.StartTime,
		&i.EndTime,
		&i.MaxMaps,
		&i.Title,
		&i.SkipVeto,
		&i.ApiKey,
		&i.Team1Score,
		&i.Team2Score,
		&i.Forfeit,
		&i.Status,
	)
	return i, err
}

const getMatchesByTeam = `-- name: GetMatchesByTeam :many
SELECT id, user_id, server_id, team1_id, team2_id, winner, cancelled, start_time, end_time, max_maps, title, skip_veto, api_key, team1_score, team2_score, forfeit, status FROM matches
WHERE team1_id = ? OR team2_id = ?
`

type GetMatchesByTeamParams struct {
	Team1ID string
	Team2ID string
}

func (q *Queries) GetMatchesByTeam(ctx context.Context, arg GetMatchesByTeamParams) ([]Match, error) {
	rows, err := q.db.QueryContext(ctx, getMatchesByTeam, arg.Team1ID, arg.Team2ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Match
	for rows.Next() {
		var i Match
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ServerID,
			&i.Team1ID,
			&i.Team2ID,
			&i.Winner,
			&i.Cancelled,
			&i.StartTime,
			&i.EndTime,
			&i.MaxMaps,
			&i.Title,
			&i.SkipVeto,
			&i.ApiKey,
			&i.Team1Score,
			&i.Team2Score,
			&i.Forfeit,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMatchesByUser = `-- name: GetMatchesByUser :many
SELECT id, user_id, server_id, team1_id, team2_id, winner, cancelled, start_time, end_time, max_maps, title, skip_veto, api_key, team1_score, team2_score, forfeit, status FROM matches
WHERE user_id = ?
`

func (q *Queries) GetMatchesByUser(ctx context.Context, userID string) ([]Match, error) {
	rows, err := q.db.QueryContext(ctx, getMatchesByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Match
	for rows.Next() {
		var i Match
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ServerID,
			&i.Team1ID,
			&i.Team2ID,
			&i.Winner,
			&i.Cancelled,
			&i.StartTime,
			&i.EndTime,
			&i.MaxMaps,
			&i.Title,
			&i.SkipVeto,
			&i.ApiKey,
			&i.Team1Score,
			&i.Team2Score,
			&i.Forfeit,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMatchesByUsers = `-- name: GetMatchesByUsers :many
SELECT id, user_id, server_id, team1_id, team2_id, winner, cancelled, start_time, end_time, max_maps, title, skip_veto, api_key, team1_score, team2_score, forfeit, status FROM matches
WHERE user_id IN (/*SLICE:user_ids*/?)
`

func (q *Queries) GetMatchesByUsers(ctx context.Context, userIds []string) ([]Match, error) {
	query := getMatchesByUsers
	var queryParams []interface{}
	if len(userIds) > 0 {
		for _, v := range userIds {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:user_ids*/?", strings.Repeat(",?", len(userIds))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:user_ids*/?", "NULL", 1)
	}
	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Match
	for rows.Next() {
		var i Match
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ServerID,
			&i.Team1ID,
			&i.Team2ID,
			&i.Winner,
			&i.Cancelled,
			&i.StartTime,
			&i.EndTime,
			&i.MaxMaps,
			&i.Title,
			&i.SkipVeto,
			&i.ApiKey,
			&i.Team1Score,
			&i.Team2Score,
			&i.Forfeit,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMatchesByWinner = `-- name: GetMatchesByWinner :many
SELECT id, user_id, server_id, team1_id, team2_id, winner, cancelled, start_time, end_time, max_maps, title, skip_veto, api_key, team1_score, team2_score, forfeit, status FROM matches
WHERE winner = ?
`

func (q *Queries) GetMatchesByWinner(ctx context.Context, winner sql.NullString) ([]Match, error) {
	rows, err := q.db.QueryContext(ctx, getMatchesByWinner, winner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Match
	for rows.Next() {
		var i Match
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ServerID,
			&i.Team1ID,
			&i.Team2ID,
			&i.Winner,
			&i.Cancelled,
			&i.StartTime,
			&i.EndTime,
			&i.MaxMaps,
			&i.Title,
			&i.SkipVeto,
			&i.ApiKey,
			&i.Team1Score,
			&i.Team2Score,
			&i.Forfeit,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const startMatch = `-- name: StartMatch :execresult
UPDATE matches
SET start_time = ?, status = 2
WHERE id = ?
`

type StartMatchParams struct {
	StartTime sql.NullTime
	ID        string
}

func (q *Queries) StartMatch(ctx context.Context, arg StartMatchParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, startMatch, arg.StartTime, arg.ID)
}

const updateMatchWinner = `-- name: UpdateMatchWinner :execresult
UPDATE matches
SET winner = ?, end_time = ?, forfeit = ?, team1_score = ?, team2_score = ?
WHERE id = ?
`

type UpdateMatchWinnerParams struct {
	Winner     sql.NullString
	EndTime    sql.NullTime
	Forfeit    sql.NullBool
	Team1Score uint32
	Team2Score uint32
	ID         string
}

func (q *Queries) UpdateMatchWinner(ctx context.Context, arg UpdateMatchWinnerParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateMatchWinner,
		arg.Winner,
		arg.EndTime,
		arg.Forfeit,
		arg.Team1Score,
		arg.Team2Score,
		arg.ID,
	)
}

const updateTeam1Score = `-- name: UpdateTeam1Score :execresult
UPDATE matches
SET team1_score = ?
WHERE id = ?
`

type UpdateTeam1ScoreParams struct {
	Team1Score uint32
	ID         string
}

func (q *Queries) UpdateTeam1Score(ctx context.Context, arg UpdateTeam1ScoreParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateTeam1Score, arg.Team1Score, arg.ID)
}

const updateTeam2Score = `-- name: UpdateTeam2Score :execresult
UPDATE matches
SET team2_score = ?
WHERE id = ?
`

type UpdateTeam2ScoreParams struct {
	Team2Score uint32
	ID         string
}

func (q *Queries) UpdateTeam2Score(ctx context.Context, arg UpdateTeam2ScoreParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateTeam2Score, arg.Team2Score, arg.ID)
}
