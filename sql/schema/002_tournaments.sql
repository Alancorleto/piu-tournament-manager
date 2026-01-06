-- +goose Up
CREATE TABLE tournaments(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    location TEXT,
    start_date TIMESTAMP
);

-- +goose Down
DROP TABLE tournaments;