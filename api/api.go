package api

import (
	"github.com/danjac/sightings/config"
	"github.com/goware/cors"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/pressly/chi/render"
	"net/http"
)

func NewRouter(cfg *config.Config) chi.Router {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)

	r.Use(render.SetContentType(render.ContentTypeJSON))

	cors := cors.Default()

	r.Use(cors.Handler)

	r.Route(cfg.Api.Path, func(r chi.Router) {

		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			render.PlainText(w, r, "pong")
		})

		r.Mount("/sightings", NewResource(cfg).Routes())
	})

	return r
}
