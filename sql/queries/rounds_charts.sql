-- name: AddChartToRound :exec
INSERT INTO rounds_charts (round_id, chart_id)
VALUES ($1, $2)
ON CONFLICT (round_id, chart_id) DO NOTHING;

-- name: RemoveChartFromRound :exec
DELETE FROM rounds_charts
WHERE round_id = $1 AND chart_id = $2;

-- name: ListChartsInRound :many
SELECT c.*
FROM charts c
JOIN rounds_charts rc ON c.id = rc.chart_id
WHERE rc.round_id = $1
ORDER BY rc.chart_id ASC;
