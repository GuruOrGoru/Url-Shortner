package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	urlModel "github.com/guruorgoru/ushort/internal/model"
	"gorm.io/gorm"
)

func RootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := struct {
			Message string `json:"message"`
		}{
			Message: "Hello World!",
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Fatalln("Error while internal encoding: ", err)
		}
	}

}

func ShortenHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var originalUrl string
		if err := json.NewDecoder(r.Body).Decode(&originalUrl); err != nil {
			http.Error(w, "Invalid Request Body", http.StatusBadRequest)
			return
		}
		if originalUrl == "" {
			http.Error(w, "old_url is required", http.StatusBadRequest)
			return
		}
		response, err := urlModel.ShortenUrl(originalUrl, db)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Error while internal encoding.", http.StatusInternalServerError)
			return
		}
	}

}

func RedirectHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var shortUrl string
		if err := json.NewDecoder(r.Body).Decode(&shortUrl); err != nil {
			http.Error(w, "Invalid Request Body", http.StatusBadRequest)
			return
		}
		originalUrl, err := urlModel.GetOldUrl(shortUrl, db)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, "Url Not Found", http.StatusNotFound)
				return
			}
			http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, originalUrl, http.StatusFound)
	}

}
