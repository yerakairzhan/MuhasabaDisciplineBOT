-- name: UpsertDuhaDone :exec
INSERT INTO duha (user_id, done)
VALUES ($1, TRUE);

-- name: CheckDuhaExists :one
SELECT EXISTS (
    SELECT 1
    FROM duha
    WHERE user_id = $1
      AND DATE(created_at) = CURRENT_DATE
);

-- name: GetDuhaData :many
select created_at from duha where user_id = $1 and done = true;
