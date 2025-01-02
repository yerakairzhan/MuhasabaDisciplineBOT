-- name: GetUsers :many
SELECT * FROM users;

-- name: InsertUser :exec
INSERT INTO users (user_id, username)
VALUES ($1, $2)
    ON CONFLICT (user_id) DO NOTHING;
