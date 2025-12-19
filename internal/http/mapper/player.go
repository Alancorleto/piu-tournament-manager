package mapper

import (
	"database/sql"

	"github.com/alancorleto/piu-tournament-manager/internal/database"
	"github.com/alancorleto/piu-tournament-manager/internal/http/dto"
	"github.com/google/uuid"
)

// PlayerResponse maps a database.Player to a dto.PlayerResponse.
func PlayerResponse(dbPlayer database.Player) dto.PlayerResponse {
	return dto.PlayerResponse{
		ID:                dbPlayer.ID,
		Nickname:          dbPlayer.Nickname,
		Name:              fromNullString(dbPlayer.Name),
		TeamName:          fromNullString(dbPlayer.TeamName),
		CountryCode:       fromNullString(dbPlayer.CountryCode),
		City:              fromNullString(dbPlayer.City),
		ProfilePictureUrl: fromNullString(dbPlayer.ProfilePictureUrl),
		CreatedAt:         dbPlayer.CreatedAt,
		ModifiedAt:        dbPlayer.ModifiedAt,
	}
}

// CreatePlayerParams maps a dto.CreatePlayerRequest to database.CreatePlayerParams.
func CreatePlayerParams(req dto.CreatePlayerRequest) database.CreatePlayerParams {
	return database.CreatePlayerParams{
		Nickname:    req.Nickname,
		Name:        toNullString(req.Name),
		TeamName:    toNullString(req.TeamName),
		CountryCode: toNullString(req.CountryCode),
		City:        toNullString(req.City),
	}
}

func UpdatePlayerParams(id uuid.UUID, req dto.UpdatePlayerRequest) database.UpdatePlayerParams {
	return database.UpdatePlayerParams{
		ID:          id,
		Nickname:    toNullString(req.Nickname),
		Name:        toNullString(req.Name),
		TeamName:    toNullString(req.TeamName),
		CountryCode: toNullString(req.CountryCode),
		City:        toNullString(req.City),
	}
}

func toNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: *s, Valid: true}
}

func fromNullString(ns sql.NullString) *string {
	if !ns.Valid {
		return nil
	}
	return &ns.String
}

func ParseUUID(s string) uuid.UUID {
	parsed, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil
	}
	return parsed
}
