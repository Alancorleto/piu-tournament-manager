-- name: ListRounds :many
SELECT *
FROM rounds
WHERE category_id = $1
ORDER BY order_index ASC;

-- name: CreateRound :one
INSERT INTO rounds (
    id,
    name,
    format,
    levels,
    qualifiers_count,
    current_state,
    category_id,
    order_index
)
SELECT
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    COALESCE(MAX(order_index), -1) + 1
FROM rounds
WHERE category_id = $6
RETURNING *;

-- name: GetRound :one
SELECT *
FROM rounds
WHERE id = $1;

-- name: UpdateRound :one
UPDATE rounds
SET
    name = COALESCE(sqlc.narg('Name'), name),
    format = COALESCE(sqlc.narg('Format'), format),
    levels = COALESCE(sqlc.narg('Levels'), levels),
    qualifiers_count = COALESCE(sqlc.narg('QualifiersCount'), qualifiers_count),
    current_state = COALESCE(sqlc.narg('CurrentState'), current_state),
    category_id = COALESCE(sqlc.narg('CategoryId'), category_id),
    order_index = COALESCE(sqlc.narg('OrderIndex'), order_index)
WHERE id = $1
RETURNING *;

-- name: DeleteRound :exec
DELETE FROM rounds
WHERE id = $1;
