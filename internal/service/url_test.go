package service

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/leesamuel423/url-shortener/internal/config"
	"github.com/leesamuel423/url-shortener/internal/models"
	"github.com/leesamuel423/url-shortener/pkg/errors"
)

// Mock repository for testing
type mockURLRepository struct {
	urls map[string]*models.URL
}

func newMockRepository() *mockURLRepository {
	return &mockURLRepository{
		urls: make(map[string]*models.URL),
	}
}

func (m *mockURLRepository) Create(ctx context.Context, url *models.URL) error {
	url.ID = len(m.urls) + 1
	url.CreatedAt = time.Now()
	m.urls[url.ShortCode] = url
	return nil
}

func (m *mockURLRepository) GetByShortCode(ctx context.Context, shortCode string) (*models.URL, error) {
	url, exists := m.urls[shortCode]
	if !exists {
		return nil, errors.ErrURLNotFound
	}
	return url, nil
}

func (m *mockURLRepository) GetByOriginalURL(ctx context.Context, originalURL string) (*models.URL, error) {
	for _, url := range m.urls {
		if url.OriginalURL == originalURL {
			return url, nil
		}
	}
	return nil, errors.ErrURLNotFound
}

func (m *mockURLRepository) IncrementClickCount(ctx context.Context, shortCode string) error {
	if url, exists := m.urls[shortCode]; exists {
		url.ClickCount++
		return nil
	}
	return errors.ErrURLNotFound
}

func (m *mockURLRepository) ShortCodeExists(ctx context.Context, shortCode string) (bool, error) {
	_, exists := m.urls[shortCode]
	return exists, nil
}

func TestShortenURL(t *testing.T) {
	repo := newMockRepository()
	cfg := &config.Config{Domain: "http://test.com"}
	service := NewURLService(repo, cfg)

	tests := []struct {
		name    string
		req     *models.ShortenRequest
		wantErr error
	}{
		{
			name: "Valid URL",
			req: &models.ShortenRequest{
				URL: "https://example.com/test",
			},
			wantErr: nil,
		},
		{
			name: "Invalid URL",
			req: &models.ShortenRequest{
				URL: "not-a-url",
			},
			wantErr: errors.ErrInvalidURL,
		},
		{
			name: "Own domain",
			req: &models.ShortenRequest{
				URL: "http://test.com/page",
			},
			wantErr: errors.ErrCannotShortenOwnDomain,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := service.ShortenURL(context.Background(), tt.req)

			if tt.wantErr != nil {
				if err != tt.wantErr {
					t.Errorf("Expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if resp.OriginalURL != tt.req.URL {
				t.Errorf("Expected original URL %s, got %s", tt.req.URL, resp.OriginalURL)
			}
		})
	}
}

func TestShortenURLWithCustomCode(t *testing.T) {
	repo := newMockRepository()
	cfg := &config.Config{Domain: "http://test.com"}
	service := NewURLService(repo, cfg)

	// First request with custom code
	req1 := &models.ShortenRequest{
		URL:        "https://example.com",
		CustomCode: "custom123",
	}

	resp1, err := service.ShortenURL(context.Background(), req1)
	if err != nil {
		t.Fatalf("Failed to shorten URL: %v", err)
	}

	if !strings.Contains(resp1.ShortURL, "custom123") {
		t.Errorf("Expected short URL to contain custom code")
	}

	// Second request with same custom code should fail
	req2 := &models.ShortenRequest{
		URL:        "https://other.com",
		CustomCode: "custom123",
	}

	_, err = service.ShortenURL(context.Background(), req2)
	if err != errors.ErrShortCodeExists {
		t.Errorf("Expected ErrShortCodeExists, got %v", err)
	}
}
