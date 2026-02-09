-- +goose Up
CREATE TABLE rounds_charts(
    round_id UUID NOT NULL REFERENCES rounds(id) ON DELETE CASCADE,
    chart_id UUID NOT NULL REFERENCES charts(id) ON DELETE CASCADE,
    order_index INT NOT NULL,
    PRIMARY KEY(round_id, chart_id),
    UNIQUE (round_id, order_index)
);

-- +goose Down
DROP TABLE rounds_charts;