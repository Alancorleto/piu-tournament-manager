-- name: CreatePlayer :one
INSERT INTO players (id, nickname, name, team_name, country_code, city, created_at, modified_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    NOW(),
    NOW()
)
RETURNING *;