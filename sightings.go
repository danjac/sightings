package sightings

import (
	"fmt"
	"strings"
	"time"
)

type AppConfig struct {
	Store Store
	Port  string
}

func (cfg *AppConfig) Close() error {
	return cfg.Store.Close()
}

type Page struct {
	PageSize   int        `json:"pageSize"`
	Number     int64      `json:"number"`
	TotalRows  int64      `json:"totalRows"`
	TotalPages int64      `json:"totalPages"`
	Sightings  []Sighting `json:"sightings"`
}

type Sighting struct {
	ID          int64     `db:"id" json:"id"`
	OccurredAt  time.Time `db:"occurred_at" json:"occurredAt"`
	ReportedAt  time.Time `db:"reported_at" json:"reportedAt"`
	Description string    `db:"description" json:"description"`
	Shape       string    `db:"shape" json:"shape"`
	Location    string    `db:"location" json:"location"`
	Duration    string    `db:"duration" json:"duration"`
	Latitude    float64   `db:"latitude" json:"latitude"`
	Longitude   float64   `db:"longitude" json:"longitude"`
}

// Print human-readable of sighting info
func (s *Sighting) String() string {

	return strings.Join([]string{
		fmt.Sprintf("Location:\t\t %s (%v, %v)", s.Location, s.Latitude, s.Longitude),
		fmt.Sprintf("Reported at:\t\t %s", s.ReportedAt),
		fmt.Sprintf("Occurred at:\t\t %s", s.OccurredAt),
		fmt.Sprintf("Shape:\t\t\t %s", s.Shape),
		fmt.Sprintf("Duration:\t\t %s", s.Duration),
		s.Description,
	}, "\n")
}

// Inserts a sighting into the database
func (s *Sighting) Insert(i Inserter) error {
	return i.Insert(s)
}

type Inserter interface {
	Insert(*Sighting) error
}

type Writer interface {
	Inserter
}

type Finder interface {
	Find(int64) (*Page, error)
}

type Getter interface {
	GetOne(string) (*Sighting, error)
}

type Searcher interface {
	Search(string, int64) (*Page, error)
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
