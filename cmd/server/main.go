package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/leesamuel423/url-shortener/internal/config"
	"github.com/leesamuel423/url-shortener/internal/database"
	"github.com/leesamuel423/url-shortener/internal/handlers"
	"github.com/leesamuel423/url-shortener/internal/repository"
	"github.com/leesamuel423/url-shortener/internal/service"
)

func main() {
	// Load configs
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// init db
	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// run migration
	if err := database.Migrate(db); err != nil {
		log.Fatal("Failed to run migration:", err)
	}

	urlRepo := repository.NewPostgresURLRepository(db)
	urlService := service.NewURLService(urlRepo, cfg)
	handler := handlers.NewHandler(urlService, cfg)

	// Setup routes
	router := handler.SetupRoutes()

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
