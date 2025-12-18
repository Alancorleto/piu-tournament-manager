-- +goose Up
CREATE TABLE players(
    id UUID PRIMARY KEY,
    nickname TEXT NOT NULL,
    name TEXT,
    team_name TEXT,
    country_code CHAR(2),
    city TEXT,
    profile_picture_url TEXT,
    created_at TIMESTAMP NOT NULL,
    modified_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE players;