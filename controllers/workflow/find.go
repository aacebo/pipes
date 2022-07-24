package workflow

import (
	"net/http"

	"github.com/aacebo/pipes/models/workflow"
	"github.com/go-chi/render"
)

func Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		workflows := workflow.Find()
		render.JSON(w, r, workflows)
	}
}
