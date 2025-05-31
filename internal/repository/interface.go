package repository

import (
	"context"

	"github.com/leesamuel423/url-shortener/internal/models"
)

type URLRepository interface {
	Create(ctx context.Context, url *models.URL) error
	GetByShortCode(ctx context.Context, shortCode string) (*models.URL, error)
	GetByOriginalURL(ctx context.Context, originalURL string) (*models.URL, error)
	IncrementClickCount(ctx context.Context, shortCode string) error
	ShortCodeExists(ctx context.Context, shortCode string) (bool, error)
}
