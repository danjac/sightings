package api

import (
	"database/sql"
	"github.com/danjac/sightings/config"
	"github.com/danjac/sightings/models"
	"github.com/danjac/sightings/store"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockReader struct {
	sighting *models.Sighting
	page     *models.Page
	err      error
}

func (r *mockReader) GetOne(_ string) (*models.Sighting, error) {
	return r.sighting, r.err
}

func (r *mockReader) Find(_ int64) (*models.Page, error) {
	return r.page, r.err
}

func (r *mockReader) Search(_ string, _ int64) (*models.Page, error) {
	return r.page, r.err
}

func testRequest(
	t *testing.T,
	cfg *config.AppConfig,
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

	cfg := &config.AppConfig{}
	cfg.Store = &store.DBStore{
		Reader: &mockReader{
			sighting: &models.Sighting{},
		},
	}

	testRequest(t, cfg, "GET", "/api/sightings/1234", http.StatusOK)
}

func TestGetSightingNotFound(t *testing.T) {

	cfg := &config.AppConfig{}

	cfg.Store = &store.DBStore{
		Reader: &mockReader{
			sighting: nil,
			err:      sql.ErrNoRows,
		},
	}

	testRequest(t, cfg, "GET", "/api/sightings/1234", http.StatusNotFound)
}

func TestListSightings(t *testing.T) {

	cfg := &config.AppConfig{}
	cfg.Store = &store.DBStore{
		Reader: &mockReader{
			page: &models.Page{},
		},
	}

	testRequest(t, cfg, "GET", "/api/sightings", http.StatusOK)

}
