-- +goose Up
CREATE TABLE charts(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    song_id UUID NOT NULL REFERENCES songs(id) ON DELETE CASCADE,
    mode INT NOT NULL,
    level INT NOT NULL,
    player_count INT NOT NULL,
    UNIQUE (song_id, mode, level, player_count)
);

-- +goose Down
DROP TABLE charts;