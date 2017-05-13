package api

import (
	"database/sql"
	"github.com/danjac/sightings/config"
	"github.com/danjac/sightings/store"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
	cfg.Store = store.NewMockStore()

	testRequest(t, cfg, "GET", "/api/sightings/1234", http.StatusOK)
}

func TestGetSightingNotFound(t *testing.T) {

	mockStore := store.NewMockStore()
	mockStore.Sighting = nil
	mockStore.Err = sql.ErrNoRows

	cfg := &config.Config{}
	cfg.Store = mockStore

	testRequest(t, cfg, "GET", "/api/sightings/1234", http.StatusNotFound)
}

func TestListSightings(t *testing.T) {

	cfg := &config.Config{}
	cfg.Store = store.NewMockStore()

	testRequest(t, cfg, "GET", "/api/sightings", http.StatusOK)

}
