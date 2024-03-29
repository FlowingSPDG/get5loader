-- name: GetTeam :one
SELECT * FROM teams
WHERE id = ? LIMIT 1;

-- name: GetTeams :many
SELECT * FROM teams
WHERE id IN (sqlc.slice('ids'));

-- name: GetTeamByUserID :many
SELECT * FROM teams
WHERE user_id = ?;

-- name: GetTeamsByUsers :many
SELECT * FROM teams
WHERE user_id IN (sqlc.slice('user_ids'));

-- name: GetPublicTeams :many
SELECT * FROM teams
WHERE public_team = TRUE;

-- name: AddTeam :execresult
INSERT INTO teams (
  id, user_id, name, flag, logo, tag, public_team
) VALUES (
  ?, ?, ?, ?, ?, ?, ?
);