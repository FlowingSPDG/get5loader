// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: players_query.sql

package players_gen

import (
	"context"
)

const getPlayer = `-- name: GetPlayer :one
SELECT id, team_id, steam_id, name FROM players
WHERE id = ? LIMIT 1
`

func (q *Queries) GetPlayer(ctx context.Context, id int64) (Player, error) {
	row := q.db.QueryRowContext(ctx, getPlayer, id)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.TeamID,
		&i.SteamID,
		&i.Name,
	)
	return i, err
}

const getPlayerBySteamID = `-- name: GetPlayerBySteamID :one
SELECT id, team_id, steam_id, name FROM players
WHERE steam_id = ? LIMIT 1
`

func (q *Queries) GetPlayerBySteamID(ctx context.Context, steamID string) (Player, error) {
	row := q.db.QueryRowContext(ctx, getPlayerBySteamID, steamID)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.TeamID,
		&i.SteamID,
		&i.Name,
	)
	return i, err
}

const getPlayersByTeam = `-- name: GetPlayersByTeam :many
SELECT id, team_id, steam_id, name FROM players
WHERE team_id = ?
`

func (q *Queries) GetPlayersByTeam(ctx context.Context, teamID int64) ([]Player, error) {
	rows, err := q.db.QueryContext(ctx, getPlayersByTeam, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Player
	for rows.Next() {
		var i Player
		if err := rows.Scan(
			&i.ID,
			&i.TeamID,
			&i.SteamID,
			&i.Name,
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
