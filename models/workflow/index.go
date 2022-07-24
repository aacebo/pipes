package workflow

import "time"

type Workflow struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func NewWorkflow(name string) *Workflow {
	v := Workflow{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &v
}

func (self *Workflow) Save() {

}
