package workflow

import (
	"net/http"

	"github.com/aacebo/pipes/models/workflow"
	"github.com/go-chi/render"
)

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := workflow.NewCreateWorkflow(r)

		if err != nil {
			render.Status(r, 400)
			render.PlainText(w, r, err.Error())
			return
		}

		entity := workflow.NewWorkflow(*body.Name)
		entity.Save()
		render.JSON(w, r, entity)
	}
}
