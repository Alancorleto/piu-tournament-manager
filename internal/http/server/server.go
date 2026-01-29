package server

import (
	"database/sql"
	"net/http"

	"github.com/alancorleto/piu-tournament-manager/internal/database"
)

type Server struct {
	http.Server
	db     *database.Queries
	dbConn *sql.DB
}

// New returns a *Server configured with the package's handlers.
func New(addr string, dbConn *sql.DB, db *database.Queries) *Server {
	mux := http.NewServeMux()

	server := Server{
		Server: http.Server{
			Addr:    addr,
			Handler: mux,
		},
		db:     db,
		dbConn: dbConn,
	}

	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("frontend"))))

	mux.HandleFunc("GET /api/health", server.GetHealth)
	mux.HandleFunc("POST /api/players", server.CreatePlayer)
	mux.HandleFunc("GET /api/players", server.ListPlayers)
	mux.HandleFunc("GET /api/players/{player_id}", server.GetPlayer)
	mux.HandleFunc("PATCH /api/players/{player_id}", server.UpdatePlayer)
	mux.HandleFunc("DELETE /api/players/{player_id}", server.DeletePlayer)

	mux.HandleFunc("POST /api/tournaments", server.CreateTournament)
	mux.HandleFunc("GET /api/tournaments", server.ListTournaments)
	mux.HandleFunc("PATCH /api/tournaments/{tournament_id}", server.UpdateTournament)
	mux.HandleFunc("DELETE /api/tournaments/{tournament_id}", server.DeleteTournament)

	mux.HandleFunc("POST /api/tournaments/{tournament_id}/categories", server.CreateCategory)
	mux.HandleFunc("GET /api/tournaments/{tournament_id}/categories/{category_id}", server.GetCategory)
	mux.HandleFunc("GET /api/tournaments/{tournament_id}/categories", server.ListCategories)
	mux.HandleFunc("PATCH /api/tournaments/{tournament_id}/categories/{category_id}", server.UpdateCategory)
	mux.HandleFunc("DELETE /api/tournaments/{tournament_id}/categories/{category_id}", server.DeleteCategory)

	mux.HandleFunc("PUT /api/tournaments/{tournament_id}/categories/{category_id}/players/{player_id}", server.AddPlayerToCategory)
	mux.HandleFunc("POST /api/tournaments/{tournament_id}/categories/{category_id}/players/bulk", server.AddPlayersToCategoryBulk)
	mux.HandleFunc("DELETE /api/tournaments/{tournament_id}/categories/{category_id}/players/{player_id}", server.RemovePlayerFromCategory)
	mux.HandleFunc("GET /api/tournaments/{tournament_id}/categories/{category_id}/players", server.ListPlayersInCategory)

	mux.HandleFunc("POST /api/tournaments/{tournament_id}/categories/{category_id}/rounds", server.CreateRound)
	mux.HandleFunc("GET /api/tournaments/{tournament_id}/categories/{category_id}/rounds/{round_id}", server.GetRound)
	mux.HandleFunc("GET /api/tournaments/{tournament_id}/categories/{category_id}/rounds", server.ListRounds)
	mux.HandleFunc("PATCH /api/tournaments/{tournament_id}/categories/{category_id}/rounds/{round_id}", server.UpdateRound)
	mux.HandleFunc("DELETE /api/tournaments/{tournament_id}/categories/{category_id}/rounds/{round_id}", server.DeleteRound)

	mux.HandleFunc("GET /api/tournaments/{tournament_id}/categories/{category_id}/rounds/{round_id}/players", server.ListPlayersInRound)
	mux.HandleFunc("POST /api/tournaments/{tournament_id}/categories/{category_id}/rounds/{round_id}/players/bulk", server.AddPlayersToRoundBulk)
	mux.HandleFunc("PUT /api/tournaments/{tournament_id}/categories/{category_id}/rounds/{round_id}/players/bulk", server.UpdateRoundPlayersOrderBulk)
	mux.HandleFunc("DELETE /api/tournaments/{tournament_id}/categories/{category_id}/rounds/{round_id}/players/{player_id}", server.RemovePlayerFromRound)

	return &server
}
