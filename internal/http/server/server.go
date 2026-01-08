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
	mux.HandleFunc("GET /api/players/{id}", server.GetPlayer)
	mux.HandleFunc("PATCH /api/players/{id}", server.UpdatePlayer)
	mux.HandleFunc("DELETE /api/players/{id}", server.DeletePlayer)

	mux.HandleFunc("POST /api/tournaments", server.CreateTournament)
	mux.HandleFunc("GET /api/tournaments", server.ListTournaments)
	mux.HandleFunc("PATCH /api/tournaments/{id}", server.UpdateTournament)
	mux.HandleFunc("DELETE /api/tournaments/{id}", server.DeleteTournament)

	mux.HandleFunc("POST /api/tournaments/{tournament_id}/categories", server.CreateCategory)
	mux.HandleFunc("GET /api/categories", server.ListCategories)
	mux.HandleFunc("GET /api/tournaments/{tournament_id}/categories", server.ListCategoriesByTournament)
	mux.HandleFunc("PATCH /api/categories/{id}", server.UpdateCategory)
	mux.HandleFunc("DELETE /api/categories/{id}", server.DeleteCategory)

	return &server
}
