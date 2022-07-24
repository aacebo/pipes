package main

import (
	"net/http"

	"github.com/aacebo/pipes/controllers"
	"github.com/aacebo/pipes/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := database.NewClient()
	r := chi.NewRouter()

	defer db.Close()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Mount("/", controllers.NewRouter())

	http.ListenAndServe(":3000", r)
}
