package user

import (
	"github.com/aacebo/pipes/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/users", Find())
	r.With(middleware.WithUser).Get("/users/{user_id}", FindByID())
	r.Post("/users", Create())

	return r
}
