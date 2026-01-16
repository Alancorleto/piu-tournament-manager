-- +goose Up
CREATE TABLE rounds(
    id UUID PRIMARY KEY,
    name TEXT,
    format INTEGER NOT NULL,
    levels TEXT,
    state INTEGER NOT NULL,
    category_id UUID NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    order_index INTEGER NOT NULL
);

-- +goose Down
DROP TABLE rounds;