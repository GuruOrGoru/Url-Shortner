package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/guruorgoru/ushort/internal/handler"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.Recoverer)

	r.Route("/api/v1", getRoutes)
	return r
}

func getRoutes(r chi.Router) {
	r.Get("/", handler.RootHandler)
	r.Post("/shorten", handler.ShortenHandler)
	r.Get("/{short}", handler.RedirectHandler)
}
