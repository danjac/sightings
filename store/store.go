package store

import (
	"github.com/danjac/sightings/models"
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

type Inserter interface {
	Insert(*models.Sighting) error
}

type Writer interface {
	Inserter
}

type Finder interface {
	Find(int64) (*models.Page, error)
}

type Getter interface {
	GetOne(string) (*models.Sighting, error)
}

type Searcher interface {
	Search(string, int64) (*models.Page, error)
}

type Reader interface {
	Finder
	Getter
	Searcher
}

type Closer interface {
	Close() error
}

type Store interface {
	Reader
	Writer
	Closer
}

type DBStore struct {
	Reader
	Writer
	Closer
}

func NewStore(db *DB) Store {
	return &DBStore{
		Reader: &DBReader{db},
		Writer: &DBWriter{db},
		Closer: db,
	}
}
