package api

import (
	"context"
	"github.com/danjac/sightings"
	"github.com/danjac/sightings/config"
	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
	"net/http"
	"strconv"
)

const sightingContextKey = "sighting"

type Resource struct {
	*config.AppConfig
}

func NewResource(cfg *config.AppConfig) *Resource {
	return &Resource{cfg}
}

func (rs *Resource) Routes() chi.Router {

	r := chi.NewRouter()

	r.Get("/", rs.List)

	r.With(rs.WithSighting).
		Route("/:id", func(r chi.Router) {
			r.Get("/", rs.Get)
			// delete, update....
		})

	return r
}

func (rs *Resource) WithSighting(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if _, err := strconv.ParseInt(id, 10, 64); err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		s, err := rs.Store.GetOne(id)

		if err != nil {
			rs.Error(w, r, err)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), sightingContextKey, s)))
	})
}

func (rs *Resource) List(w http.ResponseWriter, r *http.Request) {

	var (
		page       *sightings.Page
		pageNumber int64
		err        error
	)

	pageNumber, err = strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)

	if err != nil {
		pageNumber = 1
	}

	search := r.URL.Query().Get("s")

	if search == "" {
		page, err = rs.Store.GetAll(pageNumber)
	} else {
		page, err = rs.Store.Search(search, pageNumber)
	}

	if err != nil {
		rs.Error(w, r, err)
		return
	}

	rs.Render(w, r, NewPageResponse(page))
}

func (rs *Resource) Get(w http.ResponseWriter, r *http.Request) {

	s, ok := rs.FromRequest(r)

	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	rs.Render(w, r, NewSightingResponse(s))
}

func (rs *Resource) FromRequest(r *http.Request) (*sightings.Sighting, bool) {
	s, ok := r.Context().Value(sightingContextKey).(*sightings.Sighting)
	return s, ok
}

func (rs *Resource) Error(w http.ResponseWriter, r *http.Request, err error) {
	render.Render(w, r, ErrRender(err))
}

func (rs *Resource) Render(w http.ResponseWriter, r *http.Request, renderer render.Renderer) {
	if err := render.Render(w, r, renderer); err != nil {
		rs.Error(w, r, err)
	}
}
