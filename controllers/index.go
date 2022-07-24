package controllers

import (
	"net/http"
	"time"

	"github.com/aacebo/pipes/common"
	"github.com/aacebo/pipes/controllers/workflow"
	"github.com/aacebo/pipes/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	now := time.Now()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, common.Any{
			"uptime": utils.FormatTimeFromNow(&now),
		})
	})

	r.Mount("/", workflow.NewRouter())

	return r
}
