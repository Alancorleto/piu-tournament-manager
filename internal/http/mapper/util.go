package mapper

import (
	"database/sql"
	"time"

	"github.com/alancorleto/piu-tournament-manager/internal/database"
	"github.com/alancorleto/piu-tournament-manager/internal/http/dto"
	"github.com/google/uuid"
)

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

func toNullTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{Time: *t, Valid: true}
}

func fromNullTime(nt sql.NullTime) *time.Time {
	if !nt.Valid {
		return nil
	}
	return &nt.Time
}

func toNullInt32(i *int32) sql.NullInt32 {
	if i == nil {
		return sql.NullInt32{Valid: false}
	}
	return sql.NullInt32{Int32: *i, Valid: true}
}

func ParseUUID(s string) uuid.UUID {
	parsed, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil
	}
	return parsed
}

func toNullUUID(id *uuid.UUID) uuid.NullUUID {
	if id == nil {
		return uuid.NullUUID{Valid: false}
	}
	return uuid.NullUUID{UUID: *id, Valid: true}
}

func toNullRoundState(state *int32) database.NullRoundState {
	if state == nil {
		return database.NullRoundState{Valid: false}
	}
	return database.NullRoundState{RoundState: RoundStateFromDTO(dto.RoundState(*state)), Valid: true}
}
