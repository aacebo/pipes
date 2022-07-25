package workflow

import (
	"database/sql"
	"log"
	"time"

	"github.com/aacebo/pipes/database"
)

var db = database.NewClient()

type Workflow struct {
	ID        *int       `json:"id"`
	Name      *string    `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func NewWorkflow(name string) *Workflow {
	now := time.Now()
	v := Workflow{
		Name:      &name,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	return &v
}

func Find() []*Workflow {
	rows, err := db.Query(
		`
			SELECT *
			FROM workflows
			WHERE deleted_at IS NULL
		`,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	arr := []*Workflow{}

	for rows.Next() {
		v := Workflow{}
		err := rows.Scan(&v.ID, &v.Name, &v.CreatedAt, &v.UpdatedAt, &v.DeletedAt)

		if err != nil {
			log.Fatal(err)
		}

		arr = append(arr, &v)
	}

	return arr
}

func FindByID(id int) *Workflow {
	v := Workflow{}
	err := db.QueryRow(
		`
			SELECT *
			FROM workflows
			WHERE id = $1
		`,
		id,
	).Scan(&v.ID, &v.Name, &v.CreatedAt, &v.UpdatedAt, &v.DeletedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		log.Fatal(err)
	}

	return &v
}

func (self *Workflow) Save() {
	if self.ID == nil {
		self.Create()
	}
}

func (self *Workflow) Create() {
	err := db.QueryRow(
		`
			INSERT INTO workflows (name, created_at, updated_at)
			VALUES ($1, $2, $3)
			RETURNING id
		`,
		self.Name, self.CreatedAt, self.UpdatedAt,
	).Scan(&self.ID)

	if err != nil {
		log.Fatal(err)
	}
}
