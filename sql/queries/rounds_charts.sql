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
SELECT
  cws.*,
  rc.order_index
FROM rounds_charts rc
JOIN charts_with_songs cws ON cws.chart_id = rc.chart_id
WHERE rc.round_id = $1;

-- name: ReplaceRoundChart :exec
UPDATE rounds_charts
SET chart_id = $3
WHERE round_id = $1
  AND order_index = $2;

-- name: GetRoundChartCount :one
SELECT COUNT(*) AS count
FROM rounds_charts
WHERE round_id = $1;
