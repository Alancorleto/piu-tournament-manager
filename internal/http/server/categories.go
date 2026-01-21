package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/alancorleto/piu-tournament-manager/internal/http/codec/json"
	"github.com/alancorleto/piu-tournament-manager/internal/http/dto"
	"github.com/alancorleto/piu-tournament-manager/internal/http/mapper"
	"github.com/google/uuid"
)

func (s *Server) CreateCategory(w http.ResponseWriter, r *http.Request) {
	tournamentIDString := r.PathValue("tournament_id")
	if tournamentIDString == "" {
		json.RespondWithError(w, http.StatusBadRequest, "Missing tournament ID in URL")
		return
	}

	tournamentID := mapper.ParseUUID(tournamentIDString)
	if tournamentID == uuid.Nil {
		json.RespondWithError(w, http.StatusBadRequest, "Invalid tournament ID format")
		return
	}

	requestParams, err := json.ParseRequestParameters[dto.CreateCategoryRequest](r)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding parameters: %s", err))
		return
	}

	category, err := s.db.CreateCategory(
		r.Context(),
		mapper.CreateCategoryParams(requestParams, tournamentID),
	)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error creating category: %s", err))
		return
	}

	response := mapper.CategoryResponse(category)
	json.RespondWithJSON(w, http.StatusCreated, response)
}

func (s *Server) ListCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := s.db.ListCategories(r.Context())
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error listing categories: %s", err))
		return
	}

	response := make([]dto.CategoryResponse, len(categories))
	for i, category := range categories {
		response[i] = mapper.CategoryResponse(category)
	}

	json.RespondWithJSON(w, http.StatusOK, response)
}

func (s *Server) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	categoryIDString := r.PathValue("id")
	if categoryIDString == "" {
		json.RespondWithError(w, http.StatusBadRequest, "Missing category ID in URL")
		return
	}

	categoryID := mapper.ParseUUID(categoryIDString)
	if categoryID == uuid.Nil {
		json.RespondWithError(w, http.StatusBadRequest, "Invalid category ID format")
		return
	}

	requestParams, err := json.ParseRequestParameters[dto.UpdateCategoryRequest](r)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding parameters: %s", err))
		return
	}

	updateCategoryParams := mapper.UpdateCategoryParams(categoryID, requestParams)
	category, err := s.db.UpdateCategory(
		r.Context(),
		updateCategoryParams,
	)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error updating category: %s", err))
		return
	}

	response := mapper.CategoryResponse(category)
	json.RespondWithJSON(w, http.StatusOK, response)
}

func (s *Server) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	categoryIDString := r.PathValue("id")
	if categoryIDString == "" {
		json.RespondWithError(w, http.StatusBadRequest, "Missing category ID in URL")
		return
	}

	categoryID := mapper.ParseUUID(categoryIDString)
	if categoryID == uuid.Nil {
		json.RespondWithError(w, http.StatusBadRequest, "Invalid category ID format")
		return
	}

	err := s.db.DeleteCategory(r.Context(), categoryID)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error deleting category: %s", err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) GetCategory(w http.ResponseWriter, r *http.Request) {
	categoryIDString := r.PathValue("id")
	if categoryIDString == "" {
		json.RespondWithError(w, http.StatusBadRequest, "Missing category ID in URL")
		return
	}

	categoryID := mapper.ParseUUID(categoryIDString)
	if categoryID == uuid.Nil {
		json.RespondWithError(w, http.StatusBadRequest, "Invalid category ID format")
		return
	}

	category, err := s.db.GetCategory(r.Context(), categoryID)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error retrieving category: %s", err))
		return
	}

	response := mapper.CategoryResponse(category)
	json.RespondWithJSON(w, http.StatusOK, response)
}

func (s *Server) ListCategoriesByTournament(w http.ResponseWriter, r *http.Request) {
	tournamentIDString := r.PathValue("tournament_id")
	if tournamentIDString == "" {
		json.RespondWithError(w, http.StatusBadRequest, "Missing tournament ID in URL")
		return
	}

	tournamentID := mapper.ParseUUID(tournamentIDString)
	if tournamentID == uuid.Nil {
		json.RespondWithError(w, http.StatusBadRequest, "Invalid tournament ID format")
		return
	}

	categories, err := s.db.ListCategoriesByTournament(r.Context(), tournamentID)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error listing categories from tournament: %s", err))
		return
	}

	response := make([]dto.CategoryResponse, len(categories))
	for i, category := range categories {
		response[i] = mapper.CategoryResponse(category)
	}

	json.RespondWithJSON(w, http.StatusOK, response)
}

func (s *Server) validateCategoryInTournament(
	ctx context.Context,
	categoryID,
	tournamentID uuid.UUID,
) error {
	_, err := s.db.ValidateCategoryBelongsToTournament(
		ctx,
		mapper.ValidateCategoryBelongsToTournamentParams(
			categoryID,
			tournamentID,
		),
	)
	return err
}
