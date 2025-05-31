package service

import (
	"context"

	"github.com/leesamuel423/url-shortener/internal/models"
)

type URLService interface {
	ShortenURL(ctx context.Context, req *models.ShortenRequest) (*models.ShortenResponse, error)
	GetURLStats(ctx context.Context, shortCode string) (*models.StatsResponse, error)
	RedirectURL(ctx context.Context, shortCode string) (string, error)
}
