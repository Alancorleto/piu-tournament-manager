package dto

import (
	"github.com/google/uuid"
)

type CreateRoundRequest struct {
	Name       *string   `json:"name"`
	Format     int32     `json:"format"`
	Levels     *string   `json:"levels"`
	State      int32     `json:"state"`
	CategoryID uuid.UUID `json:"category_id"`
	OrderIndex int32     `json:"order_index"`
}

type UpdateRoundRequest struct {
	Name   *string `json:"name"`
	Format *int32  `json:"format"`
	Levels *string `json:"levels"`
	State  *int32  `json:"state"`
}

type RoundResponse struct {
	ID         uuid.UUID `json:"id"`
	Name       *string   `json:"name,omitempty"`
	Format     int32     `json:"format"`
	Levels     *string   `json:"levels,omitempty"`
	State      int32     `json:"state"`
	CategoryID uuid.UUID `json:"category_id"`
	OrderIndex int32     `json:"order_index"`
}
