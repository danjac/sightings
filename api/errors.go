package api

import "net/http"

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
