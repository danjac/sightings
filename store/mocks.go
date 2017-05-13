package store

import "github.com/danjac/sightings"

type MockStore struct {
	Err      error
	Sighting *sightings.Sighting
	Page     *sightings.Page
}

func (st *MockStore) GetOne(_ string) (*sightings.Sighting, error)      { return st.Sighting, st.Err }
func (st *MockStore) GetAll(_ int64) (*sightings.Page, error)           { return st.Page, st.Err }
func (st *MockStore) Search(_ string, _ int64) (*sightings.Page, error) { return st.Page, st.Err }

func (st *MockStore) Insert(s *sightings.Sighting) error { return st.Err }
func (st *MockStore) Close() error                       { return nil }

// create mock store with sensible defaults
func NewMockStore() *MockStore {
	return &MockStore{
		Sighting: &sightings.Sighting{},
		Page:     &sightings.Page{},
	}
}
