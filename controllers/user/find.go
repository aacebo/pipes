package user

import (
	"net/http"

	model "github.com/aacebo/pipes/models/user"
	"github.com/aacebo/pipes/page"
	"github.com/go-chi/render"
)

func Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := model.Find(page.New(r))
		render.JSON(w, r, users)
	}
}
