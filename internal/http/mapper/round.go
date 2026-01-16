package mapper

import (
	"github.com/alancorleto/piu-tournament-manager/internal/database"
	"github.com/alancorleto/piu-tournament-manager/internal/http/dto"
	"github.com/google/uuid"
)

// RoundResponse maps a database.Round to a dto.RoundResponse.
func RoundResponse(dbRound database.Round) dto.RoundResponse {
	return dto.RoundResponse{
		ID:              dbRound.ID,
		Name:            fromNullString(dbRound.Name),
		Format:          dbRound.Format,
		Levels:          fromNullString(dbRound.Levels),
		QualifiersCount: dbRound.QualifiersCount,
		State:           dbRound.State,
		CategoryID:      dbRound.CategoryID,
		OrderIndex:      dbRound.OrderIndex,
	}
}

func CreateRoundParams(req dto.CreateRoundRequest, categoryID uuid.UUID) database.CreateRoundParams {
	return database.CreateRoundParams{
		Name:            toNullString(req.Name),
		Format:          req.Format,
		Levels:          toNullString(req.Levels),
		QualifiersCount: req.QualifiersCount,
		State:           req.State,
		CategoryID:      categoryID,
	}
}

func UpdateRoundParams(id uuid.UUID, req dto.UpdateRoundRequest) database.UpdateRoundParams {
	return database.UpdateRoundParams{
		ID:              id,
		Name:            toNullString(req.Name),
		Format:          toNullInt32(req.Format),
		Levels:          toNullString(req.Levels),
		QualifiersCount: toNullInt32(req.QualifiersCount),
		State:           toNullInt32(req.State),
	}
}
