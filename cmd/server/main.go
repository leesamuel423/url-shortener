package main

import (
	"github.com/leesamuel423/url-shortener/internal/config"
	"github.com/leesamuel423/url-shortener/internal/database"
	"log"
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

}
