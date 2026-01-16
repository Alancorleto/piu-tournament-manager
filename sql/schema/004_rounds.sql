-- +goose Up
CREATE TABLE rounds(
    id UUID PRIMARY KEY,
    name TEXT,
    format INTEGER NOT NULL,
    levels TEXT,
    qualifiers_count INTEGER NOT NULL,
    state INTEGER NOT NULL,
    category_id UUID NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    order_index INTEGER NOT NULL,
    UNIQUE(category_id, order_index)
);

-- +goose Down
DROP TABLE rounds;