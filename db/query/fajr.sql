-- name: UpsertFajrDone :exec
INSERT INTO fajr (user_id, done)
VALUES ($1, TRUE);

-- name: CheckFajrExists :one
SELECT EXISTS (
    SELECT 1
    FROM fajr
    WHERE user_id = $1
      AND DATE(created_at) = CURRENT_DATE
);

-- name: GetFajrData :many
select created_at from fajr where user_id = $1 and done = true;
