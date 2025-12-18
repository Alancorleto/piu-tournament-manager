package server

import (
	"net/http"

	"github.com/alancorleto/piu-tournament-manager/internal/http/codec/json"
)

type helloMessage struct {
	Message string `json:"message"`
}

func (s *Server) HelloHandler(w http.ResponseWriter, r *http.Request) {
	json.RespondWithJSON(w, 200, helloMessage{Message: "Hello world!"})
}
