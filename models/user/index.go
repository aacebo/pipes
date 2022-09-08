package user

import (
	"database/sql"
	"log"
	"time"

	"github.com/aacebo/pipes/database"
	_page "github.com/aacebo/pipes/page"
)

var db = database.NewClient()

type User struct {
	ID        *int       `json:"id"`
	Name      *string    `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func NewUser(name string) *User {
	now := time.Now()
	v := User{
		Name:      &name,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	return &v
}

func Find(page *_page.Page) []*User {
	rows, err := db.Query(
		`
			SELECT *
			FROM users
			WHERE deleted_at IS NULL
			OFFSET $1
			LIMIT $2
		`,
		page.GetOffset(),
		page.Size,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	arr := []*User{}

	for rows.Next() {
		v := User{}
		err := rows.Scan(&v.ID, &v.Name, &v.CreatedAt, &v.UpdatedAt, &v.DeletedAt)

		if err != nil {
			log.Fatal(err)
		}

		arr = append(arr, &v)
	}

	return arr
}

func FindByID(id int) *User {
	v := User{}
	err := db.QueryRow(
		`
			SELECT *
			FROM users
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

func (self *User) Save() {
	if self.ID == nil {
		self.Create()
	}
}

func (self *User) Create() {
	err := db.QueryRow(
		`
			INSERT INTO users (name, created_at, updated_at)
			VALUES ($1, $2, $3)
			RETURNING id
		`,
		self.Name, self.CreatedAt, self.UpdatedAt,
	).Scan(&self.ID)

	if err != nil {
		log.Fatal(err)
	}
}
