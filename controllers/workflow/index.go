package workflow

import (
	"github.com/aacebo/pipes/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/workflows", Find())
	r.With(middleware.WithWorkflow).Get("/workflows/{workflow_id}", FindByID())
	r.Post("/workflows", Create())

	return r
}
