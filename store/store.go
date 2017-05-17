package store

import (
	"github.com/danjac/sightings"
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

type Store struct {
	sightings.Reader
	sightings.Writer
	sightings.Closer
}

func NewStore(db *DB) sightings.Store {
	return &Store{
		Reader: &Reader{db},
		Writer: &Writer{db},
		Closer: db,
	}
}
