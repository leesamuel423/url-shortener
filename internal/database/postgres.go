package database

import (
	"database/sql"
	"time"
)

func NewPostgresDB(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// connection pool config
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// test connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
