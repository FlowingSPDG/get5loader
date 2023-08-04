-- name: GetTeam :one
SELECT * FROM teams
WHERE id = ? LIMIT 1;

-- name: GetTeamByUserID :one
SELECT * FROM teams
WHERE user_id = ?;

-- name: GetPublicTeams :many
SELECT * FROM teams
WHERE public_team = TRUE;