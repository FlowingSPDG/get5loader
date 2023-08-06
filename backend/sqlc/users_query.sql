-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: CreateUser :execresult
INSERT INTO users (
  steam_id, name, admin, password_hash
) VALUES (
  ?, ?, ?, ?
);