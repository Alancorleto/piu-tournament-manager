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

-- name: GetPlayers :many
SELECT *
FROM players
ORDER BY created_at DESC;