package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alancorleto/piu-tournament-manager/internal/database"
	"github.com/alancorleto/piu-tournament-manager/internal/http/codec/json"
	"github.com/google/uuid"
)

type Player struct {
	ID                uuid.UUID `json:"id"`
	Nickname          string    `json:"nickname"`
	Name              *string   `json:"name"`
	TeamName          *string   `json:"team_name"`
	CountryCode       *string   `json:"country_code"`
	City              *string   `json:"city"`
	ProfilePictureUrl *string   `json:"profile_picture_url"`
	CreatedAt         time.Time `json:"created_at"`
	ModifiedAt        time.Time `json:"modified_at"`
}

func (s *Server) PostPlayersHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Nickname    string  `json:"nickname"`
		Name        *string `json:"name"`
		TeamName    *string `json:"team_name"`
		CountryCode *string `json:"country_code"`
		City        *string `json:"city"`
	}

	params, err := json.ParseRequestParameters[parameters](r)
	if err != nil {
		json.RespondWithError(w, 400, fmt.Sprintf("Error decoding parameters: %s", err))
		return
	}

	player, err := s.db.CreatePlayer(
		r.Context(),
		database.CreatePlayerParams{
			Nickname:    params.Nickname,
			Name:        toNullString(params.Name),
			TeamName:    toNullString(params.TeamName),
			CountryCode: toNullString(params.CountryCode),
			City:        toNullString(params.City),
		},
	)

	if err != nil {
		json.RespondWithError(w, 500, fmt.Sprintf("Error creating player: %s", err))
		return
	}

	json.RespondWithJSON(w, 201, databasePlayerToServerPlayer(player))
}

func databasePlayerToServerPlayer(dbPlayer database.Player) Player {
	return Player{
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
