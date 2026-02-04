package dto

import "github.com/google/uuid"

type CreateChartRequest struct {
	SongID      uuid.UUID `json:"song_id"`
	Mode        int32     `json:"mode"`
	Level       int32     `json:"level"`
	PlayerCount int32     `json:"player_count"`
}

type UpdateChartRequest struct {
	SongID      *uuid.UUID `json:"song_id,omitempty"`
	Mode        *int32     `json:"mode,omitempty"`
	Level       *int32     `json:"level,omitempty"`
	PlayerCount *int32     `json:"player_count,omitempty"`
}

type ChartResponse struct {
	ID          uuid.UUID `json:"id"`
	SongID      uuid.UUID `json:"song_id"`
	Mode        int32     `json:"mode"`
	Level       int32     `json:"level"`
	PlayerCount int32     `json:"player_count"`
}
