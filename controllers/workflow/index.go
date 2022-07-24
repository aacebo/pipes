package workflow

import (
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/workflows", Find())

	return r
}
