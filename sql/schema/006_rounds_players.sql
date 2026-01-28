-- +goose Up
CREATE TABLE rounds_players(
    round_id UUID NOT NULL REFERENCES rounds(id) ON DELETE CASCADE,
    player_id UUID NOT NULL REFERENCES players(id) ON DELETE CASCADE,
    order_index INT NOT NULL,
    PRIMARY KEY(round_id, player_id),
    UNIQUE (round_id, order_index)
);

-- +goose Down
DROP TABLE rounds_players;