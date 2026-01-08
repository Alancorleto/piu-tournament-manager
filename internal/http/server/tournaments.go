package server

import (
	"fmt"
	"net/http"

	"github.com/alancorleto/piu-tournament-manager/internal/http/codec/json"
	"github.com/alancorleto/piu-tournament-manager/internal/http/dto"
	"github.com/alancorleto/piu-tournament-manager/internal/http/mapper"
	"github.com/google/uuid"
)

func (s *Server) CreateTournament(w http.ResponseWriter, r *http.Request) {
	requestParams, err := json.ParseRequestParameters[dto.CreateTournamentRequest](r)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding parameters: %s", err))
		return
	}

	createTournamentParams := mapper.CreateTournamentParams(requestParams)
	tournament, err := s.db.CreateTournament(
		r.Context(),
		createTournamentParams,
	)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error creating tournament: %s", err))
		return
	}

	response := mapper.TournamentResponse(tournament)
	json.RespondWithJSON(w, http.StatusCreated, response)
}

func (s *Server) ListTournaments(w http.ResponseWriter, r *http.Request) {
	tournaments, err := s.db.ListTournaments(r.Context())
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error listing tournaments: %s", err))
		return
	}

	response := make([]dto.TournamentResponse, len(tournaments))
	for i, tournament := range tournaments {
		response[i] = mapper.TournamentResponse(tournament)
	}

	json.RespondWithJSON(w, http.StatusOK, response)
}

func (s *Server) UpdateTournament(w http.ResponseWriter, r *http.Request) {
	tournamentIDString := r.PathValue("id")
	if tournamentIDString == "" {
		json.RespondWithError(w, http.StatusBadRequest, "Missing tournament ID in URL")
		return
	}

	tournamentID := mapper.ParseUUID(tournamentIDString)
	if tournamentID == uuid.Nil {
		json.RespondWithError(w, http.StatusBadRequest, "Invalid tournament ID format")
		return
	}

	requestParams, err := json.ParseRequestParameters[dto.UpdateTournamentRequest](r)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding parameters: %s", err))
		return
	}

	updateTournamentParams := mapper.UpdateTournamentParams(tournamentID, requestParams)
	tournament, err := s.db.UpdateTournament(
		r.Context(),
		updateTournamentParams,
	)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error updating tournament: %s", err))
		return
	}

	response := mapper.TournamentResponse(tournament)
	json.RespondWithJSON(w, http.StatusOK, response)
}

func (s *Server) DeleteTournament(w http.ResponseWriter, r *http.Request) {
	tournamentIDString := r.PathValue("id")
	if tournamentIDString == "" {
		json.RespondWithError(w, http.StatusBadRequest, "Missing tournament ID in URL")
		return
	}

	tournamentID := mapper.ParseUUID(tournamentIDString)
	if tournamentID == uuid.Nil {
		json.RespondWithError(w, http.StatusBadRequest, "Invalid tournament ID format")
		return
	}

	err := s.db.DeleteTournament(r.Context(), tournamentID)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error deleting tournament: %s", err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) GetTournament(w http.ResponseWriter, r *http.Request) {
	tournamentIDString := r.PathValue("id")
	if tournamentIDString == "" {
		json.RespondWithError(w, http.StatusBadRequest, "Missing tournament ID in URL")
		return
	}

	tournamentID := mapper.ParseUUID(tournamentIDString)
	if tournamentID == uuid.Nil {
		json.RespondWithError(w, http.StatusBadRequest, "Invalid tournament ID format")
		return
	}

	tournament, err := s.db.GetTournament(r.Context(), tournamentID)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error retrieving tournament: %s", err))
		return
	}

	response := mapper.TournamentResponse(tournament)
	json.RespondWithJSON(w, http.StatusOK, response)
}
