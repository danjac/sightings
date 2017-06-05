package api

import (
	"github.com/pressly/chi/render"
	"net/http"
)

var (
	errNotFound            = HTTPError{http.StatusNotFound}
	errUnprocessableEntity = HTTPError{http.StatusUnprocessableEntity}
)

type HTTPError struct {
	StatusCode int
}

func (e HTTPError) StatusText() string {
	return http.StatusText(e.StatusCode)
}

func (e HTTPError) Error() string {
	return e.StatusText()
}

// wraps handler so we can just return an error

type errHandlerFunc func(w http.ResponseWriter, r *http.Request) error

func errHandler(h errHandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			render.Render(w, r, ErrRender(err))
		}
	})
}
