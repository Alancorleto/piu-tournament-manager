package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreatePlayerRequest struct {
	Nickname    string  `json:"nickname"`
	Name        *string `json:"name"`
	TeamName    *string `json:"team_name"`
	CountryCode *string `json:"country_code"`
	City        *string `json:"city"`
}

type PlayerResponse struct {
	ID                uuid.UUID `json:"id"`
	Nickname          string    `json:"nickname"`
	Name              *string   `json:"name"`
	TeamName          *string   `json:"team_name"`
	CountryCode       *string   `json:"country_code"`
	City              *string   `json:"city"`
	ProfilePictureUrl *string   `json:"profile_picture_url"`
	CreatedAt         time.Time `json:"created_at"`
	ModifiedAt        time.Time `json:"modified_at"`
}
