package server

import (
	"net/http"

	"github.com/alancorleto/piu-tournament-manager/internal/database"
	"github.com/alancorleto/piu-tournament-manager/internal/http/server/handlers"
)

type Server struct {
	db *database.Queries
}

// New returns an *http.Server configured with the package's handlers.
func New(addr string, db *database.Queries) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handlers.HelloHandler)
	mux.HandleFunc("/health", handlers.HealthHandler)

	return &http.Server{Addr: addr, Handler: mux}
}
