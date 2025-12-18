package handlers

import (
	"net/http"

	"github.com/alancorleto/piu-tournament-manager/internal/http/codec/json"
)

type helloMessage struct {
	Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.RespondWithJSON(w, 200, helloMessage{Message: "Hello world!"})
}
