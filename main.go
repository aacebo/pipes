package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aacebo/pipes/controllers"
	"github.com/aacebo/pipes/database"
	"github.com/aacebo/pipes/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	log := logger.New("main")
	port := os.Getenv("PORT")
	r := chi.NewRouter()
	db := database.NewClient()
	defer db.Close()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Mount("/", controllers.NewRouter())

	log.Infof("listening on port %s...", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
