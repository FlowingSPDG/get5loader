-- name: GetGameServer :one
SELECT * FROM game_servers
WHERE id = ? LIMIT 1;

-- name: AddGameServer :execresult
INSERT INTO game_servers (
  id, user_id, ip, port, rcon_password, display_name, is_public
) VALUES (
  ?, ?, ?, ?, ?, ?, ?
);

-- name: GetGameServersByUser :many
SELECT * FROM game_servers
WHERE user_id = ?;

-- name: GetGameServersByUsers :many
SELECT * FROM game_servers
WHERE user_id IN (sqlc.slice('user_ids'));

-- name: GetPublicGameServers :many
SELECT * FROM game_servers
WHERE is_public = TRUE ;

-- name: DeleteGameServer :execresult
DELETE FROM game_servers
WHERE id = ?;