// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: mapstats_query.sql

package mapstats_gen

import (
	"context"
	"strings"
)

const getMapStat = `-- name: GetMapStat :one
SELECT id, match_id, map_number, map_name, start_time, end_time, winner, team1_score, team2_score, forfeit FROM map_stats
WHERE id = ? LIMIT 1
`

func (q *Queries) GetMapStat(ctx context.Context, id string) (MapStat, error) {
	row := q.db.QueryRowContext(ctx, getMapStat, id)
	var i MapStat
	err := row.Scan(
		&i.ID,
		&i.MatchID,
		&i.MapNumber,
		&i.MapName,
		&i.StartTime,
		&i.EndTime,
		&i.Winner,
		&i.Team1Score,
		&i.Team2Score,
		&i.Forfeit,
	)
	return i, err
}

const getMapStatsByMatch = `-- name: GetMapStatsByMatch :many
SELECT id, match_id, map_number, map_name, start_time, end_time, winner, team1_score, team2_score, forfeit FROM map_stats
WHERE match_id = ?
`

func (q *Queries) GetMapStatsByMatch(ctx context.Context, matchID string) ([]MapStat, error) {
	rows, err := q.db.QueryContext(ctx, getMapStatsByMatch, matchID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MapStat
	for rows.Next() {
		var i MapStat
		if err := rows.Scan(
			&i.ID,
			&i.MatchID,
			&i.MapNumber,
			&i.MapName,
			&i.StartTime,
			&i.EndTime,
			&i.Winner,
			&i.Team1Score,
			&i.Team2Score,
			&i.Forfeit,
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

const getMapStatsByMatchAndMap = `-- name: GetMapStatsByMatchAndMap :one
SELECT id, match_id, map_number, map_name, start_time, end_time, winner, team1_score, team2_score, forfeit FROM map_stats
WHERE match_id = ? AND map_number = ? LIMIT 1
`

type GetMapStatsByMatchAndMapParams struct {
	MatchID   string
	MapNumber uint32
}

func (q *Queries) GetMapStatsByMatchAndMap(ctx context.Context, arg GetMapStatsByMatchAndMapParams) (MapStat, error) {
	row := q.db.QueryRowContext(ctx, getMapStatsByMatchAndMap, arg.MatchID, arg.MapNumber)
	var i MapStat
	err := row.Scan(
		&i.ID,
		&i.MatchID,
		&i.MapNumber,
		&i.MapName,
		&i.StartTime,
		&i.EndTime,
		&i.Winner,
		&i.Team1Score,
		&i.Team2Score,
		&i.Forfeit,
	)
	return i, err
}

const getMapStatsByMatches = `-- name: GetMapStatsByMatches :many
SELECT id, match_id, map_number, map_name, start_time, end_time, winner, team1_score, team2_score, forfeit FROM map_stats
WHERE match_id IN (/*SLICE:match_ids*/?)
`

func (q *Queries) GetMapStatsByMatches(ctx context.Context, matchIds []string) ([]MapStat, error) {
	query := getMapStatsByMatches
	var queryParams []interface{}
	if len(matchIds) > 0 {
		for _, v := range matchIds {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:match_ids*/?", strings.Repeat(",?", len(matchIds))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:match_ids*/?", "NULL", 1)
	}
	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MapStat
	for rows.Next() {
		var i MapStat
		if err := rows.Scan(
			&i.ID,
			&i.MatchID,
			&i.MapNumber,
			&i.MapName,
			&i.StartTime,
			&i.EndTime,
			&i.Winner,
			&i.Team1Score,
			&i.Team2Score,
			&i.Forfeit,
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
