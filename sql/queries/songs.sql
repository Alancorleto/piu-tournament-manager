-- name: ListSongs :many
SELECT *
FROM songs
ORDER BY name ASC;

-- name: CreateSong :one
INSERT INTO songs (
    id,
    name,
    title_url
)
VALUES (
    gen_random_uuid(),
    $1,
    $2
)
RETURNING *;

-- name: GetSong :one
SELECT *
FROM songs
WHERE id = $1;

-- name: UpdateSong :one
UPDATE songs
SET
    name = COALESCE(sqlc.narg('Name'), name),
    title_url = COALESCE(sqlc.narg('TitleUrl'), title_url)
WHERE id = $1
RETURNING *;

-- name: DeleteSong :exec
DELETE FROM songs
WHERE id = $1;
