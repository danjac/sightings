package api

import (
	"context"
	"fmt"
	"github.com/danjac/sightings/config"
	"github.com/danjac/sightings/models"
	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
	"net/http"
	"strconv"
)

const sightingContextKey = "sighting"

type Resource struct {
	*config.Config
}

func NewResource(cfg *config.Config) *Resource {
	return &Resource{cfg}
}

func (rs *Resource) Routes() chi.Router {

	r := chi.NewRouter()

	r.Get("/", rs.List())

	r.With(rs.WithSighting).
		Route("/:id", func(r chi.Router) {
			r.Get("/", rs.Get())
		})

	return r
}

func (rs *Resource) WithSighting(next http.Handler) http.Handler {
	return errHandler(func(w http.ResponseWriter, r *http.Request) error {
		var (
			id  int64
			err error
		)

		if id, err = strconv.ParseInt(chi.URLParam(r, "id"), 10, 64); err != nil {
			return errNotFound
		}

		s, err := rs.Repo.Get(id)

		if err != nil {
			return err
		}

		next.ServeHTTP(w, r.WithContext(newContext(r.Context(), s)))
		return nil
	})
}

func (rs *Resource) List() http.HandlerFunc {
	return errHandler(func(w http.ResponseWriter, r *http.Request) error {

		var (
			page       *models.SightingsPage
			pageNumber int64
			err        error
		)

		pageNumber, err = strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)

		if err != nil {
			pageNumber = 1
		}

		search := r.URL.Query().Get("s")

		if search == "" {
			page, err = rs.Repo.Find(pageNumber)
		} else {
			page, err = rs.Repo.Search(search, pageNumber)
		}

		if err != nil {
			return err
		}

		// assign URLs to each sighting

		for i, _ := range page.Sightings {
			s := &page.Sightings[i]
			s.URL = rs.getSightingURL(r, s)
		}

		return render.Render(w, r, NewSightingsPageResponse(r, page))
	})
}

func (rs *Resource) Get() http.HandlerFunc {
	return errHandler(func(w http.ResponseWriter, r *http.Request) error {

		s, ok := fromContext(r.Context())

		if !ok {
			return errUnprocessableEntity
		}

		return render.Render(w, r, NewSightingResponse(s))
	})
}

func (rs *Resource) getSightingURL(r *http.Request, s *models.Sighting) string {
	return fmt.Sprintf(
		"%s://%s%s/sightings/%d",
		getScheme(r),
		r.Host,
		rs.Api.Path,
		s.ID,
	)
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

func newContext(ctx context.Context, s *models.Sighting) context.Context {
	return context.WithValue(ctx, sightingContextKey, s)
}

func fromContext(ctx context.Context) (*models.Sighting, bool) {
	s, ok := ctx.Value(sightingContextKey).(*models.Sighting)
	return s, ok
}

func getScheme(r *http.Request) string {
	if r.TLS == nil {
		return "http"
	}
	return "https"
}
