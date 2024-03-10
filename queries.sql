-- name: GetPlayerByID :one
SELECT * FROM Players
WHERE id = ? LIMIT 1;

-- name: GetPlayerByMail :one
SELECT * FROM Players
WHERE email = ? LIMIT 1;

-- name: GetPlayerScores :many
SELECT * FROM Scores
WHERE player_id = ?
ORDER BY score;

-- name: RegisterPlayerScore :one
INSERT INTO Scores (score, player_id)
VALUES (?, ?)
RETURNING *;

-- name: CreatePlayer :one
INSERT INTO Players (name, email)
VALUES (?, ?)
RETURNING *;

-- name: RenamePlayer :one
UPDATE Players
SET name = ?
WHERE id = ?
RETURNING *;

-- name: CreateAuth :exec
INSERT INTO Auth (hash, player_id)
VALUES (?, ?);

-- name: GetAuth :one
SELECT * FROM Auth 
WHERE player_id = ?
LIMIT 1;

-- name: IsRegistered :one
SELECT EXISTS(SELECT 1 FROM Players WHERE email = ?) AS player_registered;
