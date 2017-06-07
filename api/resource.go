package api

import (
	"fmt"
	"github.com/danjac/sightings/config"
	"github.com/danjac/sightings/models"
	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
	"net/http"
	"strconv"
)

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

		next.ServeHTTP(w, r.WithContext(models.SightingToContext(r.Context(), s)))
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

		url := fmt.Sprintf("%s://%s%s/sightings/%%d",
			getScheme(r),
			r.Host,
			rs.Api.Path,
		)

		for i, _ := range page.Sightings {
			s := &page.Sightings[i]
			s.URL = fmt.Sprintf(url, s.ID)
		}

		return render.Render(w, r, NewSightingsPageResponse(r, page))
	})
}

func (rs *Resource) Get() http.HandlerFunc {
	return errHandler(func(w http.ResponseWriter, r *http.Request) error {

		s := models.SightingFromContext(r.Context())

		return render.Render(w, r, NewSightingResponse(s))
	})
}

func getScheme(r *http.Request) string {
	if r.TLS == nil {
		return "http"
	}
	return "https"
}
