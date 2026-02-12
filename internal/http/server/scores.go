package server

import (
	"fmt"
	"net/http"

	"github.com/alancorleto/piu-tournament-manager/internal/http/codec/json"
	"github.com/alancorleto/piu-tournament-manager/internal/http/dto"
	"github.com/alancorleto/piu-tournament-manager/internal/http/mapper"
)

func (s *Server) CreateScore(w http.ResponseWriter, r *http.Request) {
	roundID, err := mustPathUUID(r, "round_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	requestParams, err := json.ParseRequestParameters[dto.CreateScoreRequest](r)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding parameters: %s", err))
		return
	}

	createScoreParams := mapper.CreateScoreParams(requestParams, roundID)
	score, err := s.db.CreateScore(
		r.Context(),
		createScoreParams,
	)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error creating score: %s", err))
		return
	}

	response := mapper.ScoreResponse(score)
	json.RespondWithJSON(w, http.StatusCreated, response)
}

func (s *Server) ListScores(w http.ResponseWriter, r *http.Request) {
	roundID, err := mustPathUUID(r, "round_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	scores, err := s.db.ListScores(r.Context(), roundID)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error listing scores: %s", err))
		return
	}

	response := make([]dto.ScoreResponse, len(scores))
	for i, score := range scores {
		response[i] = mapper.ScoreResponse(score)
	}

	json.RespondWithJSON(w, http.StatusOK, response)
}

func (s *Server) GetScore(w http.ResponseWriter, r *http.Request) {
	scoreID, err := mustPathUUID(r, "score_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	score, err := s.db.GetScore(r.Context(), scoreID)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error getting score: %s", err))
		return
	}

	response := mapper.ScoreResponse(score)
	json.RespondWithJSON(w, http.StatusOK, response)
}

func (s *Server) DeleteScore(w http.ResponseWriter, r *http.Request) {
	scoreID, err := mustPathUUID(r, "score_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = s.db.DeleteScore(r.Context(), scoreID)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error deleting score: %s", err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) UpdateScore(w http.ResponseWriter, r *http.Request) {
	scoreID, err := mustPathUUID(r, "score_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	requestParams, err := json.ParseRequestParameters[dto.UpdateScoreRequest](r)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding parameters: %s", err))
		return
	}

	updateScoreParams := mapper.UpdateScoreParams(scoreID, requestParams)
	score, err := s.db.UpdateScore(
		r.Context(),
		updateScoreParams,
	)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error updating score: %s", err))
		return
	}

	response := mapper.ScoreResponse(score)
	json.RespondWithJSON(w, http.StatusOK, response)
}
