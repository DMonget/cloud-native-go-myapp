package router

import (
	"github.com/go-chi/chi"

	"myapp/server/app"
)

// New : create the router for our server, and register the routes.
func New() *chi.Mux {
	r := chi.NewRouter()

	r.MethodFunc("GET", "/", app.HandleIndex)

	return r
}