package server

import (
	"net/http"

	"github.com/alancorleto/piu-tournament-manager/internal/http/server/handlers"
)

// New returns an *http.Server configured with the package's handlers.
func New(addr string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handlers.HelloHandler)
	mux.HandleFunc("/health", handlers.HealthHandler)

	return &http.Server{Addr: addr, Handler: mux}
}
