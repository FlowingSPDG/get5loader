-- name: GetMapStat :one
SELECT * FROM map_stats
WHERE id = ? LIMIT 1;

-- name: GetMapStatsByMatches :many
SELECT * FROM map_stats
WHERE match_id IN (sqlc.slice('match_ids'));

-- name: GetMapStatsByMatch :many
SELECT * FROM map_stats
WHERE match_id = ?;

-- name: GetMapStatsByMatchAndMap :one
SELECT * FROM map_stats
WHERE match_id = ? AND map_number = ? LIMIT 1;