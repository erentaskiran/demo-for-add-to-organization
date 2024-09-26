package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func newDb() (*sql.DB, error) {
	connString := "user=demo dbname=demo sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
