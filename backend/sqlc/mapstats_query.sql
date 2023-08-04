-- name: GetMapStats :one
SELECT * FROM map_stats
WHERE id = ? LIMIT 1;

-- name: GetMapStatsByMatch :many
SELECT * FROM map_stats
WHERE match_id = ?;