package mapper

import (
	"time"

	"github.com/alancorleto/piu-tournament-manager/internal/database"
	"github.com/alancorleto/piu-tournament-manager/internal/http/dto"
	"github.com/google/uuid"
)

// TournamentResponse maps a database.Tournament to a dto.TournamentResponse.
func TournamentResponse(dbTournament database.Tournament) dto.TournamentResponse {
	return dto.TournamentResponse{
		ID:        dbTournament.ID,
		Name:      dbTournament.Name,
		Location:  fromNullString(dbTournament.Location),
		StartDate: fromNullTime(dbTournament.StartDate).Unix(),
	}
}

func CreateTournamentParams(req dto.CreateTournamentRequest) database.CreateTournamentParams {
	reqTime := time.Unix(req.StartDate, 0).UTC()
	return database.CreateTournamentParams{
		Name:      req.Name,
		Location:  toNullString(req.Location),
		StartDate: toNullTime(&reqTime),
	}
}

func UpdateTournamentParams(id uuid.UUID, req dto.UpdateTournamentRequest) database.UpdateTournamentParams {
	reqTime := time.Unix(req.StartDate, 0).UTC()
	return database.UpdateTournamentParams{
		ID:        id,
		Name:      toNullString(req.Name),
		Location:  toNullString(req.Location),
		StartDate: toNullTime(&reqTime),
	}
}
