package workflow

import (
	"net/http"

	"github.com/aacebo/pipes/common"
	"github.com/go-chi/render"
)

func Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, common.Any{"john": "doe"})
	}
}
