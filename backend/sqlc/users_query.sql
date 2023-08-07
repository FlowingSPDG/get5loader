-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: GetUserBySteamID :one
SELECT * FROM users
WHERE steam_id = ? LIMIT 1;

-- name: CreateUser :execresult
INSERT INTO users (
  id ,steam_id, name, admin, password_hash
) VALUES (
  ?, ?, ?, ?, ?
);