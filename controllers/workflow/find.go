package workflow

import (
	"net/http"

	model "github.com/aacebo/pipes/models/workflow"
	"github.com/aacebo/pipes/page"
	"github.com/go-chi/render"
)

func Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		workflows := model.Find(page.New(r))
		render.JSON(w, r, workflows)
	}
}
