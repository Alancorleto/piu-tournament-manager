package server

import (
	"fmt"
	"net/http"

	"github.com/alancorleto/piu-tournament-manager/internal/http/codec/json"
	"github.com/alancorleto/piu-tournament-manager/internal/http/dto"
	"github.com/alancorleto/piu-tournament-manager/internal/http/mapper"
)

func (s *Server) CreatePlayer(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) ListPlayers(w http.ResponseWriter, r *http.Request) {
	players, err := s.db.ListPlayers(r.Context())
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error listing players: %s", err))
		return
	}

	response := make([]dto.PlayerResponse, len(players))
	for i, player := range players {
		response[i] = mapper.PlayerResponse(player)
	}

	json.RespondWithJSON(w, http.StatusOK, response)
}
