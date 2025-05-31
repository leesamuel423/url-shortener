package database

import (
	"database/sql"
)

func Migrate(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS urls (
		id SERIAL PRIMARY KEY,
		short_code VARCHAR(20) UNIQUE NOT NULL,
		original_url TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		click_count INT DEFAULT 0
	);

	CREATE INDEX IF NOT EXISTS idx_short_code ON urls(short_code);
	CREATE INDEX IF NOT EXISTS idx_original_url ON urls(original_url);
	`

	_, err := db.Exec(query)
	return err
}
