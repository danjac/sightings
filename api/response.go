package api

import (
	"database/sql"
	"fmt"
	"github.com/danjac/sightings/config"
	"github.com/danjac/sightings/models"
	"github.com/pressly/chi/render"
	"net/http"
)

type SightingsPageResponse struct {
	*models.SightingsPage
}

func (resp *SightingsPageResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewSightingsPageResponse(r *http.Request, page *models.SightingsPage) *SightingsPageResponse {
	resp := &SightingsPageResponse{page}

	for i, _ := range page.Sightings {
		s := &page.Sightings[i]
		s.URL = getSightingURL(r, s)
	}
	return resp
}

type SightingResponse struct {
	*models.Sighting
}

func (resp *SightingResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewSightingResponse(s *models.Sighting) *SightingResponse {
	return &SightingResponse{s}
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrRender(err error) render.Renderer {
	var (
		statusCode int
		statusText string
	)
	switch err.(type) {
	default:
		if err == sql.ErrNoRows {
			statusCode = http.StatusNotFound
			statusText = "Page not found: cannot find this item"
		} else {
			statusCode = http.StatusInternalServerError
			statusText = "Error rendering response."
		}
	case HTTPError:
		httpErr, _ := err.(HTTPError)
		statusCode = httpErr.StatusCode
		statusText = httpErr.StatusText()
	}

	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: statusCode,
		StatusText:     statusText,
		ErrorText:      err.Error(),
	}
}

func getScheme(r *http.Request) string {
	if r.TLS == nil {
		return "http"
	}
	return "https"
}

func getSightingURL(r *http.Request, s *models.Sighting) string {
	return fmt.Sprintf(
		"%s://%s%s/sightings/%d",
		getScheme(r),
		r.Host,
		config.ApiRoot,
		s.ID,
	)
}
