package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateTournamentRequest struct {
	Name      string    `json:"name"`
	Location  *string   `json:"location"`
	StartDate time.Time `json:"start_date"`
}

type UpdateTournamentRequest struct {
	Name      *string   `json:"name"`
	Location  *string   `json:"location"`
	StartDate time.Time `json:"start_date"`
}

type TournamentResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Location  *string   `json:"location"`
	StartDate time.Time `json:"start_date"`
}
