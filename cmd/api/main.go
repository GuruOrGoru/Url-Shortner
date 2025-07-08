package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	urlModel "github.com/guruorgoru/ushort/internal/model"
	"github.com/guruorgoru/ushort/internal/router"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false,
		Logger:      logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalln("Failed to establish connection to the database")
	}
	err = db.AutoMigrate(&urlModel.Url{})
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			log.Println("Database table already exists, skipping migration")
		} else {
			log.Fatal("Failed to migrate database:", err)
		}
	} else {
		log.Println("Database migration completed successfully")
	}
	log.Println("Database migration completed successfully")

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
