package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func CreateDb() *sql.DB {
	// Initialize the database
	connStr := "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
