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

func (s *Server) AddPlayersToRoundBulk(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	roundID, err := mustPathUUID(r, "round_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	req, err := json.ParseRequestParameters[dto.AddPlayersToRoundBulkRequest](r)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	err = s.db.AddPlayersToRoundBulk(
		ctx,
		mapper.AddPlayersToRoundBulkParams(roundID, req),
	)
	if err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, "failed to add players to round: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) RemovePlayerFromRound(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	roundID, err := mustPathUUID(r, "round_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	playerID, err := mustPathUUID(r, "player_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Initiate transaction to avoid race conditions
	tx, err := s.dbConn.BeginTx(ctx, nil)
	if err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, "failed to begin transaction")
		return
	}
	defer tx.Rollback()

	// Use WithTx to get a Queries instance that uses the transaction
	txQueries := s.db.WithTx(tx)

	removedOrderIndex, err := txQueries.RemovePlayerFromRound(
		ctx,
		mapper.RemovePlayerFromRoundParams(roundID, playerID),
	)
	if err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, "failed to remove player from round: "+err.Error())
		return
	}

	err = txQueries.FixMissingIndexFromRoundPlayerOrder(
		ctx,
		mapper.FixMissingIndexFromRoundPlayerOrderParams(roundID, removedOrderIndex),
	)
	if err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, "failed to fix missing index in round player order: "+err.Error())
		return
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, "failed to commit transaction")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) ListPlayersInRound(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	roundID, err := mustPathUUID(r, "round_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	players, err := s.db.ListPlayersInRound(ctx, roundID)
	if err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, "failed to list players in round: "+err.Error())
		return
	}

	response := make([]dto.PlayerResponse, len(players))
	for i, player := range players {
		response[i] = mapper.PlayerResponse(player)
	}

	json.RespondWithJSON(w, http.StatusOK, response)
}

func (s *Server) UpdateRoundPlayersOrderBulk(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	roundID, err := mustPathUUID(r, "round_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	req, err := json.ParseRequestParameters[dto.UpdateRoundPlayersOrderBulkRequest](r)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	err = s.db.UpdateRoundPlayersOrderBulk(
		ctx,
		mapper.UpdateRoundPlayersOrderBulkParams(roundID, req),
	)
	if err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, "failed to update players order in round: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) AddChartToRound(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	roundID, err := mustPathUUID(r, "round_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	chartID, err := mustPathUUID(r, "chart_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = s.db.AddChartToRound(
		ctx,
		mapper.AddChartToRoundParams(roundID, chartID),
	)
	if err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, "failed to add chart to round: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) RemoveChartFromRound(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	roundID, err := mustPathUUID(r, "round_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	chartID, err := mustPathUUID(r, "chart_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Initiate transaction to avoid race conditions
	tx, err := s.dbConn.BeginTx(ctx, nil)
	if err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, "failed to begin transaction")
		return
	}
	defer tx.Rollback()

	// Use WithTx to get a Queries instance that uses the transaction
	txQueries := s.db.WithTx(tx)

	removedOrderIndex, err := txQueries.RemoveChartFromRound(
		ctx,
		mapper.RemoveChartFromRoundParams(roundID, chartID),
	)
	if err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, "failed to remove chart from round: "+err.Error())
		return
	}

	err = txQueries.FixMissingIndexFromRoundChartOrder(
		ctx,
		mapper.FixMissingIndexFromRoundChartOrderParams(roundID, removedOrderIndex),
	)
	if err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, "failed to fix missing index in round chart order: "+err.Error())
		return
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, "failed to commit transaction")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) ListChartsInRound(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	roundID, err := mustPathUUID(r, "round_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	charts, err := s.db.ListChartsInRound(ctx, roundID)
	if err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, "failed to list charts in round: "+err.Error())
		return
	}

	response := make([]dto.ChartInRoundResponse, len(charts))
	for i, chart := range charts {
		response[i] = mapper.ChartInRoundResponse(chart)
	}

	json.RespondWithJSON(w, http.StatusOK, response)
}

func (s *Server) ReplaceRoundChart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	roundID, err := mustPathUUID(r, "round_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	orderIndex, err := mustPathInt32(r, "order_index")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	req, err := json.ParseRequestParameters[dto.ReplaceRoundChartRequest](r)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	err = s.db.ReplaceRoundChart(
		ctx,
		mapper.ReplaceRoundChartParams(roundID, orderIndex, req.ChartID),
	)
	if err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, "failed to replace chart in round: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) StartRound(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	roundID, err := mustPathUUID(r, "round_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = s.db.StartRound(ctx, roundID)
	if err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, "failed to start round: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) CancelRoundStart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	roundID, err := mustPathUUID(r, "round_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = s.db.CancelRoundStart(ctx, roundID)
	if err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, "failed to cancel round start: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) PauseRound(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	roundID, err := mustPathUUID(r, "round_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = s.db.PauseRound(ctx, roundID)
	if err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, "failed to pause round: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) FinishRound(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	roundID, err := mustPathUUID(r, "round_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = s.db.FinishRound(ctx, roundID)
	if err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, "failed to finish round: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) ResumeRound(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	roundID, err := mustPathUUID(r, "round_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = s.db.ResumeRound(ctx, roundID)
	if err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, "failed to resume round: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
