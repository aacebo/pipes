package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var client *sql.DB

func NewClient() *sql.DB {
	if client != nil {
		return client
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_CONNECTION_STRING"))

	if err != nil {
		log.Fatal(err)
	}

	client = db
	return db
}
