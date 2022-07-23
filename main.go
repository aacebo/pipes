package main

import (
	"net/http"

	"github.com/aacebo/pipes/controllers/workflows"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
)

type H map[string]interface{}

func main() {
	godotenv.Load()
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, H{"hello": "world"})
	})

	r.Mount("/", workflows.NewRouter())

	http.ListenAndServe(":3000", r)
}
