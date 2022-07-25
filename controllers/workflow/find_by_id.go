package workflow

import (
	"net/http"

	model "github.com/aacebo/pipes/models/workflow"
	"github.com/go-chi/render"
)

func FindByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v := r.Context().Value("workflow").(*model.Workflow)
		render.JSON(w, r, v)
	}
}
