package main

import (
	"github.com/guruorgoru/ushort/internal/router"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
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
	dsn := os.Getenv("DSN_DB")
	if dsn == "" {
		log.Fatalln("Dsn Not set in set env")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to establish connection to the database")
	}

	router := router.NewRouter(db)

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
