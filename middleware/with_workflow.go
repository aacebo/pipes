package middleware

import (
	"context"
	"net/http"
	"strconv"

	model "github.com/aacebo/pipes/models/workflow"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func WithWorkflow(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "workflow_id"))

		if err != nil {
			render.Status(r, 400)
			render.PlainText(w, r, "workflow_id must be an integer")
			return
		}

		v := model.FindByID(id)

		if v == nil {
			render.Status(r, 404)
			render.PlainText(w, r, "workflow not found")
			return
		}

		ctx := context.WithValue(r.Context(), "workflow", v)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
