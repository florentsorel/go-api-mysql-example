package main

import (
	"github.com/go-chi/chi/v5"
)

func (app *application) routes() *chi.Mux {
	r := chi.NewRouter()

	r.NotFound(app.notFoundResponse)
	r.MethodNotAllowed(app.methodNotAllowedResponse)

	r.Get("/v1/actors/{id}", app.showActorHandler)
	r.Post("/v1/actors", app.createActorHandler)
	r.Get("/v1/actors", app.listActorHandler)

	return r
}
