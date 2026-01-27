-- name: AddPlayerToCategory :exec
INSERT INTO categories_players (category_id, player_id)
VALUES ($1, $2)
ON CONFLICT (category_id, player_id) DO NOTHING;

-- name: AddPlayersToCategoryBulk :exec
INSERT INTO categories_players (category_id, player_id)
SELECT $1, unnest(sqlc.narg('player_ids')::uuid[])
ON CONFLICT (category_id, player_id) DO NOTHING;

-- name: RemovePlayerFromCategory :exec
DELETE FROM categories_players
WHERE category_id = $1 AND player_id = $2;

-- name: ListPlayersInCategory :many
SELECT p.*
FROM players p
JOIN categories_players cp ON p.id = cp.player_id
WHERE cp.category_id = $1;
