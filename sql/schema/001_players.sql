-- +goose Up
CREATE TABLE players(
    id UUID PRIMARY KEY,
    nickname TEXT NOT NULL,
    name TEXT NOT NULL,
    team_name TEXT,
    country_code CHAR(2) NOT NULL,
    city TEXT,
    created_at TIMESTAMP NOT NULL,
    modified_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE players;