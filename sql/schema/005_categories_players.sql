-- +goose Up
CREATE TABLE categories_players(
    category_id UUID NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    player_id UUID NOT NULL REFERENCES players(id) ON DELETE CASCADE,
    PRIMARY KEY(category_id, player_id)
);

-- +goose Down
DROP TABLE categories_players;