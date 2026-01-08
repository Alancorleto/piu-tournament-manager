-- name: CreateCategory :one
INSERT INTO categories (id, name, tournament_id)
VALUES (
    gen_random_uuid(),
    $1,
    $2
)
RETURNING *;

-- name: ListCategories :many
SELECT *
FROM categories;

-- name: UpdateCategory :one
UPDATE categories
SET
    name = COALESCE(sqlc.narg('Name'), name)
WHERE id = $1
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;

-- name: GetCategory :one
SELECT *
FROM categories
WHERE id = $1;

-- name: ListCategoriesByTournament :many
SELECT *
FROM categories
WHERE tournament_id = $1;