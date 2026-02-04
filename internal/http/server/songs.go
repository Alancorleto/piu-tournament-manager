package server

import (
	"fmt"
	"net/http"

	"github.com/alancorleto/piu-tournament-manager/internal/http/codec/json"
	"github.com/alancorleto/piu-tournament-manager/internal/http/dto"
	"github.com/alancorleto/piu-tournament-manager/internal/http/mapper"
)

func (s *Server) CreateSong(w http.ResponseWriter, r *http.Request) {
	requestParams, err := json.ParseRequestParameters[dto.CreateSongRequest](r)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding parameters: %s", err))
		return
	}

	createSongParams := mapper.CreateSongParams(requestParams)
	song, err := s.db.CreateSong(
		r.Context(),
		createSongParams,
	)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error creating song: %s", err))
		return
	}

	response := mapper.SongResponse(song)
	json.RespondWithJSON(w, http.StatusCreated, response)
}

func (s *Server) ListSongs(w http.ResponseWriter, r *http.Request) {
	songs, err := s.db.ListSongs(r.Context())
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error listing songs: %s", err))
		return
	}

	response := make([]dto.SongResponse, len(songs))
	for i, song := range songs {
		response[i] = mapper.SongResponse(song)
	}

	json.RespondWithJSON(w, http.StatusOK, response)
}

func (s *Server) UpdateSong(w http.ResponseWriter, r *http.Request) {
	songID, err := mustPathUUID(r, "song_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	requestParams, err := json.ParseRequestParameters[dto.UpdateSongRequest](r)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding parameters: %s", err))
		return
	}

	updateSongParams := mapper.UpdateSongParams(songID, requestParams)

	song, err := s.db.UpdateSong(
		r.Context(),
		updateSongParams,
	)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error updating song: %s", err))
		return
	}

	response := mapper.SongResponse(song)
	json.RespondWithJSON(w, http.StatusOK, response)
}

func (s *Server) DeleteSong(w http.ResponseWriter, r *http.Request) {
	songID, err := mustPathUUID(r, "song_id")
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = s.db.DeleteSong(r.Context(), songID)
	if err != nil {
		json.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error deleting song: %s", err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
