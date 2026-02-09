-- +goose Up
CREATE VIEW charts_with_songs AS
SELECT
  c.id AS chart_id,
  c.mode,
  c.level,
  c.player_count,
  s.id AS song_id,
  s.name AS song_name,
  s.title_url AS song_title_url
FROM charts c
JOIN songs s ON s.id = c.song_id;

-- +goose Down
DROP VIEW charts_with_songs;