-- +goose Up
-- +goose StatementBegin

-- 1. Create the new enum type
CREATE TYPE round_state AS ENUM (
  'not_started',
  'in_progress',
  'paused',
  'finished'
);

-- 2. Drop the old state column
ALTER TABLE rounds
DROP COLUMN state;

-- 3. Add the new state column with the enum type
ALTER TABLE rounds
ADD COLUMN current_state round_state NOT NULL DEFAULT 'not_started';

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

-- 1. Add the old state column back
ALTER TABLE rounds
ADD COLUMN state INTEGER NOT NULL DEFAULT 0;

-- 2. Drop the new state column
ALTER TABLE rounds
DROP COLUMN current_state;

-- 3. Drop the enum type
DROP TYPE round_state;

-- +goose StatementEnd
