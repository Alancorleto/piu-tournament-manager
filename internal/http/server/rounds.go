package server

import (
	"net/http"

	"github.com/alancorleto/piu-tournament-manager/internal/http/codec/json"
	"github.com/alancorleto/piu-tournament-manager/internal/http/dto"
	"github.com/alancorleto/piu-tournament-manager/internal/http/mapper"
	"github.com/google/uuid"
)

func (s *Server) CreateRound(w http.ResponseWriter, r *http.Request) {
	categoryIDString := r.PathValue("category_id")
	if categoryIDString == "" {
		json.RespondWithError(w, http.StatusBadRequest, "Missing category ID in URL")
		return
	}

	categoryID := mapper.ParseUUID(categoryIDString)
	if categoryID == uuid.Nil {
		json.RespondWithError(w, http.StatusBadRequest, "Invalid category ID format")
		return
	}

	requestParams, err := json.ParseRequestParameters[dto.CreateRoundRequest](r)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, "Error decoding parameters: "+err.Error())
		return
	}

	round, err := s.db.CreateRound(
		r.Context(),
		mapper.CreateRoundParams(requestParams, categoryID),
	)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, "Error creating round: "+err.Error())
		return
	}

	response := mapper.RoundResponse(round)
	json.RespondWithJSON(w, http.StatusCreated, response)
}

func (s *Server) ListRounds(w http.ResponseWriter, r *http.Request) {
	rounds, err := s.db.ListRounds(r.Context())
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, "Error listing rounds: "+err.Error())
		return
	}

	response := make([]dto.RoundResponse, len(rounds))
	for i, round := range rounds {
		response[i] = mapper.RoundResponse(round)
	}

	json.RespondWithJSON(w, http.StatusOK, response)
}

func (s *Server) UpdateRound(w http.ResponseWriter, r *http.Request) {
	roundIDString := r.PathValue("id")
	if roundIDString == "" {
		json.RespondWithError(w, http.StatusBadRequest, "Missing round ID in URL")
		return
	}

	roundID := mapper.ParseUUID(roundIDString)
	if roundID == uuid.Nil {
		json.RespondWithError(w, http.StatusBadRequest, "Invalid round ID format")
		return
	}

	requestParams, err := json.ParseRequestParameters[dto.UpdateRoundRequest](r)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, "Error decoding parameters: "+err.Error())
		return
	}

	round, err := s.db.UpdateRound(
		r.Context(),
		mapper.UpdateRoundParams(roundID, requestParams),
	)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, "Error updating round: "+err.Error())
		return
	}

	response := mapper.RoundResponse(round)
	json.RespondWithJSON(w, http.StatusOK, response)
}

func (s *Server) DeleteRound(w http.ResponseWriter, r *http.Request) {
	roundIDString := r.PathValue("round_id")
	if roundIDString == "" {
		json.RespondWithError(w, http.StatusBadRequest, "Missing round ID in URL")
		return
	}

	roundID := mapper.ParseUUID(roundIDString)
	if roundID == uuid.Nil {
		json.RespondWithError(w, http.StatusBadRequest, "Invalid round ID format")
		return
	}

	err := s.db.DeleteRound(
		r.Context(),
		roundID,
	)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, "Error deleting round: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) GetRound(w http.ResponseWriter, r *http.Request) {
	roundIDString := r.PathValue("round_id")
	if roundIDString == "" {
		json.RespondWithError(w, http.StatusBadRequest, "Missing round ID in URL")
		return
	}

	roundID := mapper.ParseUUID(roundIDString)
	if roundID == uuid.Nil {
		json.RespondWithError(w, http.StatusBadRequest, "Invalid round ID format")
		return
	}

	round, err := s.db.GetRound(
		r.Context(),
		roundID,
	)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, "Error retrieving round: "+err.Error())
		return
	}

	response := mapper.RoundResponse(round)
	json.RespondWithJSON(w, http.StatusOK, response)
}

func (s *Server) ListRoundsByCategory(w http.ResponseWriter, r *http.Request) {
	categoryIDString := r.PathValue("category_id")
	if categoryIDString == "" {
		json.RespondWithError(w, http.StatusBadRequest, "Missing category ID in URL")
		return
	}

	categoryID := mapper.ParseUUID(categoryIDString)
	if categoryID == uuid.Nil {
		json.RespondWithError(w, http.StatusBadRequest, "Invalid category ID format")
		return
	}

	rounds, err := s.db.ListRoundsByCategory(
		r.Context(),
		categoryID,
	)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, "Error listing rounds: "+err.Error())
		return
	}

	response := make([]dto.RoundResponse, len(rounds))
	for i, round := range rounds {
		response[i] = mapper.RoundResponse(round)
	}

	json.RespondWithJSON(w, http.StatusOK, response)
}
