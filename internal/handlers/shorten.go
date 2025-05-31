package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/leesamuel423/url-shortener/internal/models"
	"github.com/leesamuel423/url-shortener/pkg/errors"
)

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req models.ShortenRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	resp, err := h.urlService.ShortenURL(r.Context(), &req)
	if err != nil {
		switch err {
		case errors.ErrInvalidURL:
			h.respondError(w, http.StatusBadRequest, "Invalid URL")
		case errors.ErrShortCodeExists:
			h.respondError(w, http.StatusConflict, "Custom code already in use")
		case errors.ErrCannotShortenOwnDomain:
			h.respondError(w, http.StatusBadRequest, "Cannot shorten own domain")
		default:
			h.respondError(w, http.StatusInternalServerError, "Failed to shorten URL")
		}
		return
	}

	h.respondJSON(w, http.StatusOK, resp)
}
