-- name: AddPlayerToRound :exec
INSERT INTO rounds_players (round_id, player_id, order_index)
SELECT
    $1,
    $2,
    COALESCE(MAX(order_index), -1) + 1
FROM rounds_players
WHERE round_id = $1
ON CONFLICT (round_id, player_id) DO NOTHING;

-- name: GetMaxOrderIndexInRound :one
SELECT COALESCE(MAX(order_index), 0)
FROM rounds_players
WHERE round_id = $1
FOR UPDATE;

-- name: AddPlayersToRoundBulk :exec
INSERT INTO rounds_players (round_id, player_id, order_index)
SELECT
  sqlc.arg(round_id)::uuid        AS round_id,
  unnest(sqlc.arg(player_ids)::uuid[]) AS player_id,
  generate_series(
    sqlc.arg(order_base)::int,
    (sqlc.arg(order_base) + array_length(sqlc.arg(player_ids), 1) - 1)::int
  ) AS order_index;

-- name: RemovePlayerFromRound :exec
DELETE FROM rounds_players
WHERE round_id = $1 AND player_id = $2;

-- name: ListPlayersInRound :many
SELECT p.*
FROM players p
JOIN rounds_players cp ON p.id = cp.player_id
WHERE cp.round_id = $1;
