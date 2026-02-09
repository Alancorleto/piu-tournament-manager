-- name: AddChartToRound :exec
INSERT INTO rounds_charts (round_id, chart_id, order_index)
SELECT
    $1,
    $2,
    COALESCE(MAX(order_index), -1) + 1
FROM rounds_charts
WHERE round_id = $1
ON CONFLICT (round_id, chart_id) DO NOTHING;

-- name: RemoveChartFromRound :one
DELETE FROM rounds_charts
WHERE round_id = $1 AND chart_id = $2
RETURNING order_index;

-- name: FixMissingIndexFromRoundChartOrder :exec
UPDATE rounds_charts
SET order_index = order_index - 1
WHERE round_id = $1 AND order_index > $2;

-- name: ListChartsInRound :many
SELECT c.*
FROM charts c
JOIN rounds_charts rc ON c.id = rc.chart_id
WHERE rc.round_id = $1
ORDER BY rc.order_index ASC;

-- name: ReplaceRoundChart :exec
WITH removed_chart AS (
    DELETE FROM rounds_charts
    WHERE round_id = $1 AND order_index = $2
    RETURNING chart_id
) INSERT INTO rounds_charts (round_id, chart_id, order_index)
VALUES ($1, $3, $2);

-- name: GetRoundChartCount :one
SELECT COUNT(*) AS count
FROM rounds_charts
WHERE round_id = $1;
