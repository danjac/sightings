package api

import (
	"database/sql"
	"github.com/danjac/sightings"
	"github.com/danjac/sightings/store"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockReader struct {
	sighting *sightings.Sighting
	page     *sightings.Page
	err      error
}

func (r *mockReader) GetOne(_ string) (*sightings.Sighting, error) {
	return r.sighting, r.err
}

func (r *mockReader) Find(_ int64) (*sightings.Page, error) {
	return r.page, r.err
}

func (r *mockReader) Search(_ string, _ int64) (*sightings.Page, error) {
	return r.page, r.err
}

func testRequest(
	t *testing.T,
	cfg *sightings.AppConfig,
	method string,
	url string,
	expectedStatus int) *httptest.ResponseRecorder {

	api := NewRouter(cfg, "/api")

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	api.ServeHTTP(w, req)

	if status := w.Code; status != expectedStatus {
		t.Errorf("Expected status %d, got %d: %v", expectedStatus, status, w.Body.String())
	}

	return w

}

func TestGetSighting(t *testing.T) {

	cfg := &sightings.AppConfig{}
	cfg.Store = &store.Store{
		Reader: &mockReader{
			sighting: &sightings.Sighting{},
		},
	}

	testRequest(t, cfg, "GET", "/api/sightings/1234", http.StatusOK)
}

func TestGetSightingNotFound(t *testing.T) {

	cfg := &sightings.AppConfig{}

	cfg.Store = &store.Store{
		Reader: &mockReader{
			sighting: nil,
			err:      sql.ErrNoRows,
		},
	}

	testRequest(t, cfg, "GET", "/api/sightings/1234", http.StatusNotFound)
}

func TestListSightings(t *testing.T) {

	cfg := &sightings.AppConfig{}
	cfg.Store = &store.Store{
		Reader: &mockReader{
			page: &sightings.Page{},
		},
	}

	testRequest(t, cfg, "GET", "/api/sightings", http.StatusOK)

}
