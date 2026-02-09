-- name: ListCharts :many
SELECT *
FROM charts
ORDER BY song_id ASC;

-- name: CreateChart :one
INSERT INTO charts (
    id,
    song_id,
    mode,
    level,
    player_count
)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4
)
ON CONFLICT (song_id, mode, level, player_count) DO UPDATE SET song_id = charts.song_id
RETURNING *;

-- name: GetChart :one
SELECT *
FROM charts
WHERE id = $1;

-- name: UpdateChart :one
UPDATE charts
SET
    song_id = COALESCE(sqlc.narg('SongID'), song_id),
    mode = COALESCE(sqlc.narg('Mode'), mode),
    level = COALESCE(sqlc.narg('Level'), level),
    player_count = COALESCE(sqlc.narg('PlayerCount'), player_count)
WHERE id = $1
RETURNING *;

-- name: DeleteChart :exec
DELETE FROM charts
WHERE id = $1;