package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alancorleto/piu-tournament-manager/internal/http/codec/json"
	"github.com/alancorleto/piu-tournament-manager/internal/http/dto"
	"github.com/alancorleto/piu-tournament-manager/internal/http/mapper"
	"github.com/google/uuid"
)

type Player struct {
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

func (s *Server) PostPlayersHandler(w http.ResponseWriter, r *http.Request) {
	requestParams, err := json.ParseRequestParameters[dto.CreatePlayerRequest](r)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding parameters: %s", err))
		return
	}

	createPlayerParams := mapper.CreatePlayerParams(requestParams)
	player, err := s.db.CreatePlayer(
		r.Context(),
		createPlayerParams,
	)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error creating player: %s", err))
		return
	}

	response := mapper.PlayerResponse(player)
	json.RespondWithJSON(w, http.StatusCreated, response)
}
