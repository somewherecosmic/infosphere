-- name: CreateUser :exec
INSERT INTO users (username, email, password_hash)
VALUES ($1, $2, $3);

-- name: FindUserByHandle :one
SELECT username, email, password_hash FROM users
WHERE email = $1 OR username = $1;