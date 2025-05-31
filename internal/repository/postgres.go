package repository

import (
	"context"
	"database/sql"

	"github.com/leesamuel423/url-shortener/internal/models"
	"github.com/leesamuel423/url-shortener/pkg/errors"
)

type postgresURLRepository struct {
	db *sql.DB
}

func NewPostgresURLRepository(db *sql.DB) URLRepository {
	return &postgresURLRepository{db: db}
}

func (r *postgresURLRepository) Create(ctx context.Context, url *models.URL) error {
	query := `
		INSERT INTO urls (short_code, original_url) 
		VALUES ($1, $2) 
		RETURNING id, created_at`

	err := r.db.QueryRowContext(
		ctx, query, url.ShortCode, url.OriginalURL,
	).Scan(&url.ID, &url.CreatedAt)

	if err != nil {
		return errors.Wrap(err, "failed to create URL")
	}

	return nil
}

func (r *postgresURLRepository) GetByShortCode(ctx context.Context, shortCode string) (*models.URL, error) {
	var url models.URL
	query := `
		SELECT id, short_code, original_url, created_at, click_count 
		FROM urls 
		WHERE short_code = $1`

	err := r.db.QueryRowContext(ctx, query, shortCode).Scan(
		&url.ID, &url.ShortCode, &url.OriginalURL, &url.CreatedAt, &url.ClickCount,
	)

	if err == sql.ErrNoRows {
		return nil, errors.ErrURLNotFound
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to get URL")
	}

	return &url, nil
}

func (r *postgresURLRepository) GetByOriginalURL(ctx context.Context, originalURL string) (*models.URL, error) {
	var url models.URL
	query := `
		SELECT id, short_code, original_url, created_at, click_count 
		FROM urls 
		WHERE original_url = $1`

	err := r.db.QueryRowContext(ctx, query, originalURL).Scan(
		&url.ID, &url.ShortCode, &url.OriginalURL, &url.CreatedAt, &url.ClickCount,
	)

	if err == sql.ErrNoRows {
		return nil, errors.ErrURLNotFound
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to get URL")
	}

	return &url, nil
}

func (r *postgresURLRepository) IncrementClickCount(ctx context.Context, shortCode string) error {
	query := `UPDATE urls SET click_count = click_count + 1 WHERE short_code = $1`
	_, err := r.db.ExecContext(ctx, query, shortCode)
	return errors.Wrap(err, "failed to increment click count")
}

func (r *postgresURLRepository) ShortCodeExists(ctx context.Context, shortCode string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM urls WHERE short_code = $1)`

	err := r.db.QueryRowContext(ctx, query, shortCode).Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "failed to check short code existence")
	}

	return exists, nil
}
