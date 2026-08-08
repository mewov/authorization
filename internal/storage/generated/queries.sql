-- name: CreateUser :one
INSERT INTO users (
    id,
    username,
    password_hash
)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;


-- name: GetUserByID :one
SELECT
    id,
    username,
    password_hash,
    created_at,
    updated_at
FROM users
WHERE id = $1
LIMIT 1;


-- name: GetUserByUsername :one
SELECT
    id,
    username,
    password_hash,
    created_at,
    updated_at
FROM users
WHERE username = $1
LIMIT 1;


-- name: UpdateUserPassword :exec
UPDATE users
SET
    password_hash = $2,
    updated_at = NOW()
WHERE id = $1;


-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;


-- name: CreateSession :exec
INSERT INTO sessions (
    token_hash,
    user_id,
    expires_at
)
VALUES (
    $1,
    $2,
    $3
);


-- name: GetSession :one
SELECT
    token_hash,
    user_id,
    expires_at,
    created_at
FROM sessions
WHERE token_hash = $1
LIMIT 1;


-- name: DeleteSession :exec
DELETE FROM sessions
WHERE token_hash = $1;


-- name: DeleteUserSessions :exec
DELETE FROM sessions
WHERE user_id = $1;


-- name: DeleteExpiredSessions :exec
DELETE FROM sessions
WHERE expires_at <= NOW();