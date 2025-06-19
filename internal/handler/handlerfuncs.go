package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	urlModel "github.com/guruorgoru/ushort/internal/model"
)

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	var originalUrl string
	if err := json.NewDecoder(r.Body).Decode(&originalUrl); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	response := urlModel.ShortenUrl(originalUrl)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error while internal encoding.", http.StatusInternalServerError)
		return
	}
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "short")
	originalUrl, err := urlModel.GetOldUrl(shortUrl)
	if err != nil {
		http.Error(w, "Url Not Found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, originalUrl, http.StatusFound)
}
