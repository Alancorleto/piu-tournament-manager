-- name: CreateScore :one
INSERT INTO scores (
    id,
    round_id,
    player_id,
    chart_id,
    score,
    perfect,
    great,
    good,
    bad,
    miss,
    max_combo,
    kcal,
    grade,
    stage_pass,
    video_url
) VALUES (
    gen_random_uuid(),
    sqlc.arg('RoundID'),
    sqlc.arg('PlayerID'),
    sqlc.arg('ChartID'),
    sqlc.arg('Score'),
    COALESCE(sqlc.narg('Perfect'), 0),
    COALESCE(sqlc.narg('Great'), 0),
    COALESCE(sqlc.narg('Good'), 0),
    COALESCE(sqlc.narg('Bad'), 0),
    COALESCE(sqlc.narg('Miss'), 0),
    COALESCE(sqlc.narg('MaxCombo'), 0),
    COALESCE(sqlc.narg('Kcal'), 0),
    COALESCE(sqlc.narg('Grade')::grade, 'F'::grade),
    COALESCE(sqlc.narg('StagePass'), false),
    COALESCE(sqlc.narg('VideoUrl'), '')
)
RETURNING *;

-- name: ListScores :many
SELECT * FROM scores WHERE round_id = sqlc.arg('RoundID') ORDER BY score DESC;

-- name: UpdateScore :one
UPDATE scores SET
    score = COALESCE(sqlc.narg('Score'), score),
    perfect = COALESCE(sqlc.narg('Perfect'), perfect),
    great = COALESCE(sqlc.narg('Great'), great),
    good = COALESCE(sqlc.narg('Good'), good),
    bad = COALESCE(sqlc.narg('Bad'), bad),
    miss = COALESCE(sqlc.narg('Miss'), miss),
    max_combo = COALESCE(sqlc.narg('MaxCombo'), max_combo),
    kcal = COALESCE(sqlc.narg('Kcal'), kcal),
    grade = COALESCE(sqlc.narg('Grade')::grade, grade),
    stage_pass = COALESCE(sqlc.narg('StagePass'), stage_pass),
    video_url = COALESCE(sqlc.narg('VideoUrl'), video_url)
WHERE id = sqlc.arg('ID')
RETURNING *;

-- name: DeleteScore :exec
DELETE FROM scores WHERE id = sqlc.arg('ID');

-- name: GetScore :one
SELECT * FROM scores WHERE id = sqlc.arg('ID');