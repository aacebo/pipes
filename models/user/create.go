package user

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type CreateUser struct {
	Name *string `json:"name"`
}

func NewCreateUser(r *http.Request) (*CreateUser, error) {
	v := &CreateUser{}
	err := render.DecodeJSON(r.Body, &v)

	if err != nil {
		return nil, err
	}

	if len(*v.Name) == 0 {
		return nil, errors.New("name is required")
	}

	return v, err
}
