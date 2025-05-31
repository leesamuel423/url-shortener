package errors

import (
	"errors"
	"fmt"
)

var (
	ErrURLNotFound            = errors.New("url not found")
	ErrInvalidURL             = errors.New("invalid url")
	ErrShortCodeExists        = errors.New("short code already exists")
	ErrCannotShortenOwnDomain = errors.New("cannot shorten own domain")
	ErrFailedToGenerateCode   = errors.New("failed to generate unique short code")
)

func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", message, err)
}
