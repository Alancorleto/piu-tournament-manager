package dto

import "github.com/google/uuid"

type CreateSongRequest struct {
	Name     string  `json:"name"`
	TitleURL *string `json:"title_url,omitempty"`
}

type UpdateSongRequest struct {
	Name     *string `json:"name,omitempty"`
	TitleURL *string `json:"title_url,omitempty"`
}

type SongResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	TitleURL *string   `json:"title_url,omitempty"`
}
