-- name: GetTeam :one
SELECT * FROM teams
WHERE id = ? LIMIT 1;

-- name: GetTeamByUserID :many
SELECT * FROM teams
WHERE user_id = ?;

-- name: GetPublicTeams :many
SELECT * FROM teams
WHERE public_team = TRUE;

-- name: AddTeam :execresult
INSERT INTO teams (
  user_id, name, flag, logo, tag, public_team
) VALUES (
  ?, ?, ?, ?, ?, ?
);