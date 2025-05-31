package models

import (
	"time"
)

type URL struct {
	ID          int       `json:"id" db:"id"`
	ShortCode   string    `json:"short_code" db:"short_code"`
	OriginalURL string    `json:"original_url" db:"original_url"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	ClickCount  int       `json:"click_count" db:"click_count"`
}

type ShortenRequest struct {
	URL        string `json:"url" validate:"required,url"`
	CustomCode string `json:"custom_code,omitempty" validate:"omitempty,alphanum,min=3,max=20"`
}

type ShortenResponse struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type StatsResponse struct {
	*URL
	ShortURL string `json:"short_url"`
}
