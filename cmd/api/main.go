package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/guruorgoru/ushort/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln(err)
	}
	portstr := os.Getenv("PORT")
	if portstr == "" {
		log.Fatalln("Port not set in env")
	}

	router := router.NewRouter()

	server := &http.Server{
		Addr:         ":" + portstr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Server started on port", portstr)
	log.Fatalln(server.ListenAndServe())
}
