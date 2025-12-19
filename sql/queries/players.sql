-- name: CreatePlayer :one
INSERT INTO players (id, nickname, name, team_name, country_code, city, created_at, modified_at)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    $5,
    NOW(),
    NOW()
)
RETURNING *;

-- name: ListPlayers :many
SELECT *
FROM players
ORDER BY created_at DESC;

-- name: UpdatePlayer :one
UPDATE players
SET
    nickname = COALESCE(sqlc.narg('Nickname'), nickname),
    name = COALESCE(sqlc.narg('Name'), name),
    team_name = COALESCE(sqlc.narg('TeamName'), team_name),
    country_code = COALESCE(sqlc.narg('CountryCode'), country_code),
    city = COALESCE(sqlc.narg('City'), city),
    modified_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeletePlayer :exec
DELETE FROM players
WHERE id = $1;
