-- name: CreateVideo :one
INSERT INTO videos(id, created_at, updated_at, title, description, duration_seconds)
VALUES (gen_random_uuid(), NOW(), NOW(), $1, $2, $3)
RETURNING *;

-- name: GetVideoByID :one
SELECT * FROM videos
WHERE id=$1;

-- name: GetVideos :many
SELECT * FROM videos
ORDER BY created_at $3
LIMIT $1 OFFSET $2;

-- name: GetVideosAsc :many
SELECT * FROM videos
ORDER BY created_at ASC
LIMIT $1 OFFSET $2;

-- name: GetVideosDesc :many
SELECT * FROM videos
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;
