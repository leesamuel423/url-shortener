package service

import (
	"context"
	"fmt"
	"log"

	"github.com/leesamuel423/url-shortener/internal/config"
	"github.com/leesamuel423/url-shortener/internal/models"
	"github.com/leesamuel423/url-shortener/internal/repository"
	"github.com/leesamuel423/url-shortener/internal/utils"
	"github.com/leesamuel423/url-shortener/pkg/errors"
)

type urlService struct {
	repo   repository.URLRepository
	config *config.Config
}

func NewURLService(repo repository.URLRepository, cfg *config.Config) URLService {
	return &urlService{
		repo:   repo,
		config: cfg,
	}
}

func (s *urlService) ShortenURL(ctx context.Context, req *models.ShortenRequest) (*models.ShortenResponse, error) {
	// validate URL
	if !utils.IsValidURL(req.URL) {
		return nil, errors.ErrInvalidURL
	}

	// check if we're trying to shorten our own domain
	if utils.IsSameDomain(req.URL, s.config.Domain) {
		return nil, errors.ErrCannotShortenOwnDomain
	}

	var shortCode string

	// handle custom code
	if req.CustomCode != "" {
		exists, err := s.repo.ShortCodeExists(ctx, req.CustomCode)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, errors.ErrShortCodeExists
		}
		shortCode = req.CustomCode
	} else {
		// check if URL already exists
		existingURL, err := s.repo.GetByOriginalURL(ctx, req.URL)
		if err == nil && existingURL != nil {
			return &models.ShortenResponse{
				ShortURL:    fmt.Sprintf("%s/%s", s.config.Domain, existingURL.ShortCode),
				OriginalURL: existingURL.OriginalURL,
			}, nil
		}

		// generate new short code
		shortCode, err = s.generateUniqueShortCode(ctx)
		if err != nil {
			return nil, err
		}
	}

	// create new URL
	url := &models.URL{
		ShortCode:   shortCode,
		OriginalURL: req.URL,
	}

	if err := s.repo.Create(ctx, url); err != nil {
		return nil, err
	}

	return &models.ShortenResponse{
		ShortURL:    fmt.Sprintf("%s/%s", s.config.Domain, shortCode),
		OriginalURL: url.OriginalURL,
	}, nil
}

func (s *urlService) GetURLStats(ctx context.Context, shortCode string) (*models.StatsResponse, error) {
	url, err := s.repo.GetByShortCode(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	return &models.StatsResponse{
		URL:      url,
		ShortURL: fmt.Sprintf("%s/%s", s.config.Domain, url.ShortCode),
	}, nil
}

func (s *urlService) RedirectURL(ctx context.Context, shortCode string) (string, error) {
	log.Printf("RedirectURL called with shortCode: %s", shortCode)
	url, err := s.repo.GetByShortCode(ctx, shortCode)
	if err != nil {
		log.Printf("Error getting URL for shortCode %s: %v", shortCode, err)
		return "", err
	}

	// increment click count asynchronously
	go func() {
		_ = s.repo.IncrementClickCount(context.Background(), shortCode)
	}()

	return url.OriginalURL, nil
}

func (s *urlService) generateUniqueShortCode(ctx context.Context) (string, error) {
	maxAttempts := 10

	for range maxAttempts {
		code := utils.GenerateShortCode()
		exists, err := s.repo.ShortCodeExists(ctx, code)
		if err != nil {
			return "", err
		}
		if !exists {
			return code, nil
		}
	}

	return "", errors.ErrFailedToGenerateCode
}
