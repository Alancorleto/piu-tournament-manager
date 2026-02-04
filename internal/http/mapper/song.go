package mapper

import (
	"github.com/alancorleto/piu-tournament-manager/internal/database"
	"github.com/alancorleto/piu-tournament-manager/internal/http/dto"
	"github.com/google/uuid"
)

func SongResponse(dbSong database.Song) dto.SongResponse {
	return dto.SongResponse{
		ID:       dbSong.ID,
		Name:     dbSong.Name,
		TitleURL: fromNullString(dbSong.TitleUrl),
	}
}

func CreateSongParams(req dto.CreateSongRequest) database.CreateSongParams {
	return database.CreateSongParams{
		Name:     req.Name,
		TitleUrl: toNullString(req.TitleURL),
	}
}

func UpdateSongParams(id uuid.UUID, req dto.UpdateSongRequest) database.UpdateSongParams {
	return database.UpdateSongParams{
		ID:       id,
		Name:     toNullString(req.Name),
		TitleUrl: toNullString(req.TitleURL),
	}
}
