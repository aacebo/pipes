package workflows

import (
	"net/http"

	"github.com/go-chi/render"
)

type H map[string]interface{}

func Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, H{"john": "doe"})
	}
}
