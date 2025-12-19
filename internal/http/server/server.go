package server

import (
	"net/http"

	"github.com/alancorleto/piu-tournament-manager/internal/database"
)

type Server struct {
	http.Server
	db *database.Queries
}

// New returns a *Server configured with the package's handlers.
func New(addr string, db *database.Queries) *Server {
	mux := http.NewServeMux()

	server := Server{
		Server: http.Server{
			Addr:    addr,
			Handler: mux,
		},
		db: db,
	}

	mux.HandleFunc("GET /api/health", server.GetHealth)
	mux.HandleFunc("POST /api/players", server.CreatePlayer)

	return &server
}
