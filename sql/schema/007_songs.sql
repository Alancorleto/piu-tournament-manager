-- +goose Up
CREATE TABLE songs(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    title_url TEXT
);

-- +goose Down
DROP TABLE songs;