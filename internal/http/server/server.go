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

	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("frontend"))))
	mux.HandleFunc("GET /api/health", server.GetHealth)
	mux.HandleFunc("POST /api/players", server.CreatePlayer)
	mux.HandleFunc("GET /api/players", server.ListPlayers)
	mux.HandleFunc("PATCH /api/players/{id}", server.UpdatePlayer)
	mux.HandleFunc("DELETE /api/players/{id}", server.DeletePlayer)

	return &server
}
