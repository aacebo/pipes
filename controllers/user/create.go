package user

import (
	"net/http"

	model "github.com/aacebo/pipes/models/user"
	"github.com/go-chi/render"
)

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := model.NewCreateUser(r)

		if err != nil {
			render.Status(r, 400)
			render.PlainText(w, r, err.Error())
			return
		}

		entity := model.NewUser(*body.Name)
		entity.Save()
		render.Status(r, 201)
		render.JSON(w, r, entity)
	}
}
