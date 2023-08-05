-- name: GetPlayerStats :one
SELECT * FROM player_stats
WHERE id = ? LIMIT 1;

-- name: GetPlayerStatsByMatch :many
SELECT * FROM player_stats
WHERE match_id = ?;

-- name: GetPlayerStatsByMap :one
SELECT * FROM player_stats
WHERE map_id = ?;

-- name: GetPlayerStatsByTeam :many
SELECT * FROM player_stats
WHERE team_id = ?;

-- name: GetPlayerStatsBySteamID :many
SELECT * FROM player_stats
WHERE steam_id = ?;