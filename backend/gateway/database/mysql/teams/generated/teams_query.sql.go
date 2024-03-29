// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: teams_query.sql

package teams_gen

import (
	"context"
	"database/sql"
	"strings"
)

const addTeam = `-- name: AddTeam :execresult
INSERT INTO teams (
  id, user_id, name, flag, logo, tag, public_team
) VALUES (
  ?, ?, ?, ?, ?, ?, ?
)
`

type AddTeamParams struct {
	ID         string
	UserID     string
	Name       string
	Flag       string
	Logo       string
	Tag        string
	PublicTeam sql.NullBool
}

func (q *Queries) AddTeam(ctx context.Context, arg AddTeamParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addTeam,
		arg.ID,
		arg.UserID,
		arg.Name,
		arg.Flag,
		arg.Logo,
		arg.Tag,
		arg.PublicTeam,
	)
}

const getPublicTeams = `-- name: GetPublicTeams :many
SELECT id, user_id, name, flag, logo, tag, public_team FROM teams
WHERE public_team = TRUE
`

func (q *Queries) GetPublicTeams(ctx context.Context) ([]Team, error) {
	rows, err := q.db.QueryContext(ctx, getPublicTeams)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Team
	for rows.Next() {
		var i Team
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.Flag,
			&i.Logo,
			&i.Tag,
			&i.PublicTeam,
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

const getTeam = `-- name: GetTeam :one
SELECT id, user_id, name, flag, logo, tag, public_team FROM teams
WHERE id = ? LIMIT 1
`

func (q *Queries) GetTeam(ctx context.Context, id string) (Team, error) {
	row := q.db.QueryRowContext(ctx, getTeam, id)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Flag,
		&i.Logo,
		&i.Tag,
		&i.PublicTeam,
	)
	return i, err
}

const getTeamByUserID = `-- name: GetTeamByUserID :many
SELECT id, user_id, name, flag, logo, tag, public_team FROM teams
WHERE user_id = ?
`

func (q *Queries) GetTeamByUserID(ctx context.Context, userID string) ([]Team, error) {
	rows, err := q.db.QueryContext(ctx, getTeamByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Team
	for rows.Next() {
		var i Team
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.Flag,
			&i.Logo,
			&i.Tag,
			&i.PublicTeam,
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

const getTeams = `-- name: GetTeams :many
SELECT id, user_id, name, flag, logo, tag, public_team FROM teams
WHERE id IN (/*SLICE:ids*/?)
`

func (q *Queries) GetTeams(ctx context.Context, ids []string) ([]Team, error) {
	query := getTeams
	var queryParams []interface{}
	if len(ids) > 0 {
		for _, v := range ids {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:ids*/?", strings.Repeat(",?", len(ids))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:ids*/?", "NULL", 1)
	}
	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Team
	for rows.Next() {
		var i Team
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.Flag,
			&i.Logo,
			&i.Tag,
			&i.PublicTeam,
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

const getTeamsByUsers = `-- name: GetTeamsByUsers :many
SELECT id, user_id, name, flag, logo, tag, public_team FROM teams
WHERE user_id IN (/*SLICE:user_ids*/?)
`

func (q *Queries) GetTeamsByUsers(ctx context.Context, userIds []string) ([]Team, error) {
	query := getTeamsByUsers
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
	var items []Team
	for rows.Next() {
		var i Team
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.Flag,
			&i.Logo,
			&i.Tag,
			&i.PublicTeam,
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
