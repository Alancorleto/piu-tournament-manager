package server

import (
	"fmt"
	"net/http"

	"github.com/alancorleto/piu-tournament-manager/internal/http/codec/json"
	"github.com/alancorleto/piu-tournament-manager/internal/http/dto"
	"github.com/alancorleto/piu-tournament-manager/internal/http/mapper"
	"github.com/google/uuid"
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

func (s *Server) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	playerIDString := r.PathValue("id")
	if playerIDString == "" {
		json.RespondWithError(w, http.StatusBadRequest, "Missing player ID in URL")
		return
	}

	playerID := mapper.ParseUUID(playerIDString)
	if playerID == uuid.Nil {
		json.RespondWithError(w, http.StatusBadRequest, "Invalid player ID format")
		return
	}

	requestParams, err := json.ParseRequestParameters[dto.UpdatePlayerRequest](r)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding parameters: %s", err))
		return
	}

	updatePlayerParams := mapper.UpdatePlayerParams(playerID, requestParams)
	player, err := s.db.UpdatePlayer(
		r.Context(),
		updatePlayerParams,
	)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error updating player: %s", err))
		return
	}

	response := mapper.PlayerResponse(player)
	json.RespondWithJSON(w, http.StatusOK, response)
}

func (s *Server) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	playerIDString := r.PathValue("id")
	if playerIDString == "" {
		json.RespondWithError(w, http.StatusBadRequest, "Missing player ID in URL")
		return
	}

	playerID := mapper.ParseUUID(playerIDString)
	if playerID == uuid.Nil {
		json.RespondWithError(w, http.StatusBadRequest, "Invalid player ID format")
		return
	}

	err := s.db.DeletePlayer(r.Context(), playerID)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error deleting player: %s", err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) GetPlayer(w http.ResponseWriter, r *http.Request) {
	playerIDString := r.PathValue("id")
	if playerIDString == "" {
		json.RespondWithError(w, http.StatusBadRequest, "Missing player ID in URL")
		return
	}

	playerID := mapper.ParseUUID(playerIDString)
	if playerID == uuid.Nil {
		json.RespondWithError(w, http.StatusBadRequest, "Invalid player ID format")
		return
	}

	player, err := s.db.GetPlayer(r.Context(), playerID)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error retrieving player: %s", err))
		return
	}

	response := mapper.PlayerResponse(player)
	json.RespondWithJSON(w, http.StatusOK, response)
}
