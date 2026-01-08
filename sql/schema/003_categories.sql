-- +goose Up
CREATE TABLE categories(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    tournament_id UUID NOT NULL REFERENCES tournaments(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE categories;