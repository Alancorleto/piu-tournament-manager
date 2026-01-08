-- +goose Up
CREATE TABLE categories(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE categories;