-- name: AddPlayerToRound :exec
INSERT INTO rounds_players (round_id, player_id, order_index)
SELECT
    $1,
    $2,
    COALESCE(MAX(order_index), -1) + 1
FROM rounds_players
WHERE round_id = $1
ON CONFLICT (round_id, player_id) DO NOTHING;

-- name: AddPlayersToRoundBulk :exec
WITH last_order AS (
    -- Get the current maximum order_index for the round
    SELECT COALESCE(MAX(order_index), 0) as value
    FROM rounds_players
    WHERE round_id = $1
)
INSERT INTO rounds_players (round_id, player_id, order_index)
  SELECT 
    $1, 
    new_player.id, 
    last_order.value + new_player.array_index
  FROM unnest(@player_ids::uuid[]) WITH ORDINALITY AS new_player(id, array_index)
  CROSS JOIN last_order;

-- name: RemovePlayerFromRound :one
DELETE FROM rounds_players
WHERE round_id = $1 AND player_id = $2
RETURNING order_index;

-- name: FixMissingIndexFromRoundPlayerOrder :exec
UPDATE rounds_players
SET order_index = order_index - 1
WHERE round_id = $1 AND order_index > $2;

-- name: UpdateRoundPlayersOrderBulk :exec
WITH incoming_order AS (
    SELECT 
        player_id, 
        new_position
    FROM unnest(@player_ids::uuid[]) WITH ORDINALITY AS u(player_id, new_position)
)
UPDATE rounds_players
SET order_index = incoming_order.new_position
FROM incoming_order
WHERE rounds_players.round_id = $1 
  AND rounds_players.player_id = incoming_order.player_id;

-- name: ListPlayersInRound :many
SELECT p.*
FROM players p
JOIN rounds_players cp ON p.id = cp.player_id
WHERE cp.round_id = $1
ORDER BY cp.order_index ASC;
