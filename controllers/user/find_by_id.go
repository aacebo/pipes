package user

import (
	"net/http"

	model "github.com/aacebo/pipes/models/user"
	"github.com/go-chi/render"
)

func FindByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v := r.Context().Value("user").(*model.User)
		render.JSON(w, r, v)
	}
}
