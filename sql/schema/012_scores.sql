-- +goose Up
-- +goose StatementBegin

-- 1. Create the new enum type
CREATE TYPE grade AS ENUM (
    'F',
    'D',
    'C',
    'B',
    'A',
    'A_P',
    'AA',
    'AA_P',
    'AAA',
    'AAA_P',
    'S',
    'S_P',
    'SS',
    'SS_P',
    'SSS',
    'SSS_P'
);

-- 1. Create the table
CREATE TABLE scores(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    round_id UUID NOT NULL REFERENCES rounds(id) ON DELETE CASCADE,
    player_id UUID NOT NULL REFERENCES players(id) ON DELETE CASCADE,
    chart_id UUID NOT NULL REFERENCES charts(id) ON DELETE CASCADE,
    score INT NOT NULL,
    perfect INT,
    great INT,
    good INT,
    bad INT,
    miss INT,
    max_combo INT,
    kcal FLOAT,
    grade grade,
    stage_pass BOOLEAN,
    video_url TEXT,
    UNIQUE (round_id, player_id, chart_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE scores;
DROP TYPE grade;
-- +goose StatementEnd