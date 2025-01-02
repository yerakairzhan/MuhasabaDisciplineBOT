-- name: UpsertTafsirDone :exec
INSERT INTO tafsir (user_id, done)
VALUES ($1, TRUE);

-- name: CheckTafsirExists :one
SELECT EXISTS (
    SELECT 1
    FROM tafsir
    WHERE user_id = $1
      AND DATE(created_at) = CURRENT_DATE
);

-- name: GetTafsirData :many
select created_at from tafsir where user_id = $1 and done = true;

