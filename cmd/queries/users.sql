-- name: CreateUser :exec
INSERT INTO users (id, created_at, updated_at, user_oauth_id, name, email) VALUES($1, $2, $3, $4, $5, $6);

-- name: GetUser :one
SELECT * FROM users WHERE user_oauth_id = $1;
