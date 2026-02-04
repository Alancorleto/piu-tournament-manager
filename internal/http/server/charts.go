package server

import (
	"fmt"
	"net/http"

	"github.com/alancorleto/piu-tournament-manager/internal/http/codec/json"
	"github.com/alancorleto/piu-tournament-manager/internal/http/dto"
	"github.com/alancorleto/piu-tournament-manager/internal/http/mapper"
)

func (s *Server) CreateChart(w http.ResponseWriter, r *http.Request) {
	requestParams, err := json.ParseRequestParameters[dto.CreateChartRequest](r)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding parameters: %s", err))
		return
	}

	createChartParams := mapper.CreateChartParams(requestParams)
	chart, err := s.db.CreateChart(
		r.Context(),
		createChartParams,
	)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error creating chart: %s", err))
		return
	}

	response := mapper.ChartResponse(chart)
	json.RespondWithJSON(w, http.StatusCreated, response)
}

func (s *Server) ListCharts(w http.ResponseWriter, r *http.Request) {
	charts, err := s.db.ListCharts(r.Context())
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error listing charts: %s", err))
		return
	}

	response := make([]dto.ChartResponse, len(charts))
	for i, chart := range charts {
		response[i] = mapper.ChartResponse(chart)
	}

	json.RespondWithJSON(w, http.StatusOK, response)
}

func (s *Server) UpdateChart(w http.ResponseWriter, r *http.Request) {
	chartID, err := mustPathUUID(r, "chart_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	requestParams, err := json.ParseRequestParameters[dto.UpdateChartRequest](r)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding parameters: %s", err))
		return
	}

	updateChartParams := mapper.UpdateChartParams(chartID, requestParams)
	chart, err := s.db.UpdateChart(
		r.Context(),
		updateChartParams,
	)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error updating chart: %s", err))
		return
	}

	response := mapper.ChartResponse(chart)
	json.RespondWithJSON(w, http.StatusOK, response)
}

func (s *Server) DeleteChart(w http.ResponseWriter, r *http.Request) {
	chartID, err := mustPathUUID(r, "chart_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = s.db.DeleteChart(
		r.Context(),
		chartID,
	)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error deleting chart: %s", err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
