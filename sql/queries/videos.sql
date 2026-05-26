-- name: CreateVideo :one
INSERT INTO videos(id, created_at, updated_at, title, description, duration_seconds)
VALUES (gen_random_uuid(), NOW(), NOW(), $1, $2, $3)
RETURNING *;

-- name: GetVideoByID :one
SELECT * FROM videos
WHERE id=$1;
