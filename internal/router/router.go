package router

import (
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/guruorgoru/ushort/internal/handler"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.Recoverer)

	r.Route("/api/v1", getRoutes)
	r.Get("/", serveIndex)
	return r
}

func getRoutes(r chi.Router) {
	r.Get("/", handler.RootHandler)
	r.Post("/shorten", handler.ShortenHandler)
	r.Get("/{short}", handler.RedirectHandler)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	htmlPath, err := filepath.Abs("web/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	http.ServeFile(w, r, htmlPath)
}
