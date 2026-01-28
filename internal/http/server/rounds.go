package server

import (
	"net/http"

	"github.com/alancorleto/piu-tournament-manager/internal/http/codec/json"
	"github.com/alancorleto/piu-tournament-manager/internal/http/dto"
	"github.com/alancorleto/piu-tournament-manager/internal/http/mapper"
)

func (s *Server) CreateRound(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tournamentID, err := mustPathUUID(r, "tournament_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	categoryID, err := mustPathUUID(r, "category_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	req, err := json.ParseRequestParameters[dto.CreateRoundRequest](r)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	err = s.validateCategoryInTournament(ctx, categoryID, tournamentID)
	if err != nil {
		json.RespondWithError(w, http.StatusNotFound, "category not found in tournament")
		return
	}

	round, err := s.db.CreateRound(
		ctx,
		mapper.CreateRoundParams(req, categoryID),
	)
	if err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, "failed to create round")
		return
	}

	response := mapper.RoundResponse(round)

	json.RespondWithJSON(w, http.StatusCreated, response)
}

func (s *Server) ListRounds(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tournamentID, err := mustPathUUID(r, "tournament_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	categoryID, err := mustPathUUID(r, "category_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = s.validateCategoryInTournament(ctx, categoryID, tournamentID)
	if err != nil {
		json.RespondWithError(w, http.StatusNotFound, "category not found in tournament")
		return
	}

	rounds, err := s.db.ListRounds(ctx, categoryID)
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
	ctx := r.Context()

	tournamentID, err := mustPathUUID(r, "tournament_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	categoryID, err := mustPathUUID(r, "category_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	roundID, err := mustPathUUID(r, "round_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	requestParams, err := json.ParseRequestParameters[dto.UpdateRoundRequest](r)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, "Error decoding parameters: "+err.Error())
		return
	}

	err = s.validateCategoryInTournament(ctx, categoryID, tournamentID)
	if err != nil {
		json.RespondWithError(w, http.StatusNotFound, "category not found in tournament")
		return
	}

	round, err := s.db.UpdateRound(
		ctx,
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
	ctx := r.Context()

	tournamentID, err := mustPathUUID(r, "tournament_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	categoryID, err := mustPathUUID(r, "category_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	roundID, err := mustPathUUID(r, "round_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = s.validateCategoryInTournament(ctx, categoryID, tournamentID)
	if err != nil {
		json.RespondWithError(w, http.StatusNotFound, "category not found in tournament")
		return
	}

	err = s.db.DeleteRound(
		ctx,
		roundID,
	)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, "Error deleting round: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) GetRound(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tournamentID, err := mustPathUUID(r, "tournament_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	categoryID, err := mustPathUUID(r, "category_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	roundID, err := mustPathUUID(r, "round_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = s.validateCategoryInTournament(ctx, categoryID, tournamentID)
	if err != nil {
		json.RespondWithError(w, http.StatusNotFound, "category not found in tournament")
		return
	}

	round, err := s.db.GetRound(
		ctx,
		roundID,
	)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, "Error retrieving round: "+err.Error())
		return
	}

	response := mapper.RoundResponse(round)
	json.RespondWithJSON(w, http.StatusOK, response)
}

// func (s *Server) AddPlayersToRoundBulk(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()

// 	roundID, err := mustPathUUID(r, "round_id")
// 	if err != nil {
// 		json.RespondWithError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	req, err := json.ParseRequestParameters[dto.AddPlayersToRoundBulkRequest](r)
// 	if err != nil {
// 		json.RespondWithError(w, http.StatusBadRequest, "invalid request body")
// 		return
// 	}

// 	// Initiate transaction to avoid race conditions
// 	tx, err := s.dbConn.BeginTx(ctx, nil)
// 	if err != nil {
// 		json.RespondWithError(w, http.StatusInternalServerError, "failed to begin transaction")
// 		return
// 	}
// 	defer tx.Rollback()

// 	// Use WithTx to get a Queries instance that uses the transaction
// 	txQueries := s.db.WithTx(tx)

// 	maxIndex, err := txQueries.GetMaxOrderIndexInRound(ctx, roundID)
// 	if err != nil {
// 		json.RespondWithError(w, http.StatusInternalServerError, "failed to get max order index for round: "+err.Error())
// 		return
// 	}

// 	err = txQueries.AddPlayersToRoundBulk(
// 		ctx,
// 		mapper.AddPlayersToRoundBulkParams(roundID, req),
// 	)
// 	if err != nil {
// 		json.RespondWithError(w, http.StatusInternalServerError, "failed to add players to round: "+err.Error())
// 		return
// 	}

// 	// Commit transaction
// 	if err = tx.Commit(); err != nil {
// 		json.RespondWithError(w, http.StatusInternalServerError, "failed to commit transaction")
// 		return
// 	}

// 	w.WriteHeader(http.StatusNoContent)
// }
