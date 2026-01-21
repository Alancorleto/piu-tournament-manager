package mapper

import (
	"github.com/alancorleto/piu-tournament-manager/internal/database"
	"github.com/alancorleto/piu-tournament-manager/internal/http/dto"
	"github.com/google/uuid"
)

// CategoryResponse maps a database.Category to a dto.CategoryResponse.
func CategoryResponse(dbCategory database.Category) dto.CategoryResponse {
	return dto.CategoryResponse{
		ID:           dbCategory.ID,
		Name:         dbCategory.Name,
		TournamentID: dbCategory.TournamentID,
	}
}

func CreateCategoryParams(req dto.CreateCategoryRequest, tournamentID uuid.UUID) database.CreateCategoryParams {
	return database.CreateCategoryParams{
		Name:         req.Name,
		TournamentID: tournamentID,
	}
}

func UpdateCategoryParams(id uuid.UUID, req dto.UpdateCategoryRequest) database.UpdateCategoryParams {
	return database.UpdateCategoryParams{
		ID:   id,
		Name: toNullString(req.Name),
	}
}

func ValidateCategoryBelongsToTournamentParams(categoryID, tournamentID uuid.UUID) database.ValidateCategoryBelongsToTournamentParams {
	return database.ValidateCategoryBelongsToTournamentParams{
		ID:           categoryID,
		TournamentID: tournamentID,
	}
}
