package workflow

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type CreateWorkflow struct {
	Name *string `json:"name"`
}

func NewCreateWorkflow(r *http.Request) (*CreateWorkflow, error) {
	v := &CreateWorkflow{}
	err := render.DecodeJSON(r.Body, &v)

	if err != nil {
		return nil, err
	}

	if len(*v.Name) == 0 {
		return nil, errors.New("name is required")
	}

	return v, err
}
