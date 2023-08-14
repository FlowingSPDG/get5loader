-- name: GetPlayer :one
SELECT * FROM players
WHERE id = ? LIMIT 1;

-- name: GetPlayerBySteamID :one
SELECT * FROM players
WHERE steam_id = ? LIMIT 1;

-- name: GetPlayersByTeam :many
SELECT * FROM players
WHERE team_id = ?;

-- name: GetPlayersByTeams :many
SELECT * FROM players
WHERE team_id IN (sqlc.slice('team_ids'));

-- name: AddPlayer :execresult
INSERT INTO players (
  id,
  team_id,
  steam_id,
  name
) VALUES (
  ?,
  ?,
  ?,
  ?
);