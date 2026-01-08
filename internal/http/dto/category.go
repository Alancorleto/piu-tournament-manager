package dto

import (
	"github.com/google/uuid"
)

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type UpdateCategoryRequest struct {
	Name *string `json:"name"`
}

type CategoryResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	TournamentID uuid.UUID `json:"tournament_id"`
}
