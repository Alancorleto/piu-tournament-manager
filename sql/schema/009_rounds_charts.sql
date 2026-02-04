-- +goose Up
CREATE TABLE rounds_charts(
    round_id UUID NOT NULL REFERENCES rounds(id) ON DELETE CASCADE,
    chart_id UUID NOT NULL REFERENCES charts(id) ON DELETE CASCADE,
    PRIMARY KEY(round_id, chart_id)
);

-- +goose Down
DROP TABLE rounds_charts;