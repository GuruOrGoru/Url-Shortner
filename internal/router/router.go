package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/guruorgoru/ushort/internal/handler"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	r.Post("/red", handler.RedirectHandler(db))
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/", handler.RootHandler())
		r.Post("/shorten", handler.ShortenHandler(db))
	})
	return r
}
