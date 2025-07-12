package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/guruorgoru/ushort/internal/handler"
	"gorm.io/gorm"
)

func skipNgrokWarning(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("ngrok-skip-browser-warning", "true")
		next.ServeHTTP(w, r)
	})
}

func NewRouter(db *gorm.DB) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.Recoverer)
	r.Use(skipNgrokWarning)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	r.Get("/{short}", handler.RedirectHandler(db))
	r.Get("/health", handler.HealthHandler())
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/", handler.RootHandler())
		r.Post("/shorten", handler.ShortenHandler(db))
	})
	return r
}
