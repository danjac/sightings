package api

import (
	"database/sql"
	"github.com/danjac/sightings/config"
	"github.com/danjac/sightings/models"
	"github.com/danjac/sightings/repo"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockRepo struct {
	repo.Reader
	repo.Writer
	repo.Closer
}

type mockReader struct {
	sighting *models.Sighting
	page     *models.Page
	err      error
}

func (r *mockReader) Get(_ int64) (*models.Sighting, error) {
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
	cfg *config.Config,
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

	cfg := &config.Config{}
	cfg.Repo = &mockRepo{
		Reader: &mockReader{
			sighting: &models.Sighting{},
		},
	}

	testRequest(t, cfg, "GET", "/api/sightings/1234", http.StatusOK)
}

func TestGetSightingNotFound(t *testing.T) {

	cfg := &config.Config{}

	cfg.Repo = &mockRepo{
		Reader: &mockReader{
			sighting: nil,
			err:      sql.ErrNoRows,
		},
	}

	testRequest(t, cfg, "GET", "/api/sightings/1234", http.StatusNotFound)
}

func TestListSightings(t *testing.T) {

	cfg := &config.Config{}
	cfg.Repo = &mockRepo{
		Reader: &mockReader{
			page: &models.Page{},
		},
	}

	testRequest(t, cfg, "GET", "/api/sightings", http.StatusOK)

}
