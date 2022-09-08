package middleware

import (
	"context"
	"net/http"
	"strconv"

	model "github.com/aacebo/pipes/models/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func WithUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "user_id"))

		if err != nil {
			render.Status(r, 400)
			render.PlainText(w, r, "user_id must be an integer")
			return
		}

		v := model.FindByID(id)

		if v == nil {
			render.Status(r, 404)
			render.PlainText(w, r, "user not found")
			return
		}

		ctx := context.WithValue(r.Context(), "user", v)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
