-- name: GetMatch :one
SELECT * FROM matches
WHERE id = ? LIMIT 1;

-- name: AddMatch :execresult
INSERT INTO matches (
  id, user_id, server_id, team1_id, team2_id, start_time, end_time, max_maps, title, skip_veto, api_key, status
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: GetMatchesByUser :many
SELECT * FROM matches
WHERE user_id = ?;

-- name: GetMatchesByTeam :many
SELECT * FROM matches
WHERE team1_id = ? OR team2_id = ?;

-- name: GetMatchesByWinner :many
SELECT * FROM matches
WHERE winner = ?;

-- name: UpdateTeam1Score :execresult
UPDATE matches
SET team1_score = ?
WHERE id = ?;

-- name: UpdateTeam2Score :execresult
UPDATE matches
SET team2_score = ?
WHERE id = ?;

-- name: UpdateMatchWinner :execresult
UPDATE matches
SET winner = ?, end_time = ?, forfeit = ?, team1_score = ?, team2_score = ?
WHERE id = ?;

-- name: CancelMatch :execresult
UPDATE matches
SET status = 4
WHERE id = ?;

-- name: StartMatch :execresult
UPDATE matches
SET start_time = ?, status = 2
WHERE id = ?;