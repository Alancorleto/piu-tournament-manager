-- name: CreateTournament :one
INSERT INTO tournaments (id, name, location, start_date)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3
)
RETURNING *;

-- name: ListTournaments :many
SELECT *
FROM tournaments
ORDER BY start_date DESC;

-- name: UpdateTournament :one
UPDATE tournaments
SET
    name = COALESCE(sqlc.narg('Name'), name),
    location = COALESCE(sqlc.narg('Location'), location),
    start_date = COALESCE(sqlc.narg('StartDate'), start_date)
WHERE id = $1
RETURNING *;

-- name: DeleteTournament :exec
DELETE FROM tournaments
WHERE id = $1;

-- name: GetTournament :one
SELECT *
FROM tournaments
WHERE id = $1;
