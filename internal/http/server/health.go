package server

import (
	"net/http"

	"github.com/alancorleto/piu-tournament-manager/internal/http/codec/json"
)

type healthMessage struct {
	Status string `json:"status"`
}

func (s *Server) HealthHandler(w http.ResponseWriter, r *http.Request) {
	json.RespondWithJSON(w, 200, healthMessage{Status: "ok"})
}
