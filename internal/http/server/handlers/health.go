package handlers

import (
	"net/http"

	"github.com/alancorleto/piu-tournament-manager/internal/http/codec/json"
)

type healthMessage struct {
	Status string `json:"status"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.RespondWithJSON(w, 200, healthMessage{Status: "ok"})
}
