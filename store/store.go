package store

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(connection string) (*DB, error) {
	db, err := sqlx.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return newDB(db), nil

}
