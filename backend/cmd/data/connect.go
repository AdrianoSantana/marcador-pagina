package data

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var connStr string = "postgresql://postgres:password@localhost/record-books?sslmode=disable"

func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
