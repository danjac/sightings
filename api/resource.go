package api

import (
	"context"
	"errors"
	"github.com/danjac/sightings/config"
	"github.com/danjac/sightings/models"
	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
	"net/http"
	"strconv"
)

var invalidRoute = errors.New("Invalid route")
var missingContext = errors.New("Unprocessable entity")

const sightingContextKey = "sighting"

type Resource struct {
	*config.AppConfig
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func NewResource(cfg *config.AppConfig) *Resource {
	return &Resource{cfg}
}

func (rs *Resource) Routes() chi.Router {

	r := chi.NewRouter()

	r.Get("/", handle(rs.List))

	r.With(rs.WithSighting).
		Route("/:id", func(r chi.Router) {
			r.Get("/", handle(rs.Get))
			// delete, update....
		})

	return r
}

func (rs *Resource) WithSighting(next http.Handler) http.Handler {
	return handle(func(w http.ResponseWriter, r *http.Request) error {

		var (
			id  int64
			err error
		)

		if id, err = strconv.ParseInt(chi.URLParam(r, "id"), 10, 64); err != nil {
			return invalidRoute
		}

		s, err := rs.Store.Get(id)

		if err != nil {
			return err
		}

		next.ServeHTTP(w, r.WithContext(newContext(r.Context(), s)))
		return nil
	})
}

func (rs *Resource) List(w http.ResponseWriter, r *http.Request) error {

	var (
		page       *models.Page
		pageNumber int64
		err        error
	)

	pageNumber, err = strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)

	if err != nil {
		pageNumber = 1
	}

	search := r.URL.Query().Get("s")

	if search == "" {
		page, err = rs.Store.Find(pageNumber)
	} else {
		page, err = rs.Store.Search(search, pageNumber)
	}

	if err != nil {
		return err
	}

	return render.Render(w, r, NewPageResponse(page))

}

func (rs *Resource) Get(w http.ResponseWriter, r *http.Request) error {

	s, ok := fromContext(r.Context())

	if !ok {
		return missingContext
	}

	return render.Render(w, r, NewSightingResponse(s))

}

// wraps handler so we can just return an error
func handle(h HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			render.Render(w, r, ErrRender(err))
		}
	})
}

func newContext(ctx context.Context, s *models.Sighting) context.Context {
	return context.WithValue(ctx, sightingContextKey, s)
}

func fromContext(ctx context.Context) (*models.Sighting, bool) {
	s, ok := ctx.Value(sightingContextKey).(*models.Sighting)
	return s, ok
}
