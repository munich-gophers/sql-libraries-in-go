
-- name: CreateUser :exec
INSERT INTO users (name, age, birthday, created_at, updated_at) VALUES ($1, $2, $3, $4, $5);

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users SET name = $1 WHERE id = $2;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
