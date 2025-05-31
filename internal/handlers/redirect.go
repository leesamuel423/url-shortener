package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leesamuel423/url-shortener/pkg/errors"
)

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]
	log.Printf("Redirect handler called with shortCode: %s", shortCode)

	originalURL, err := h.urlService.RedirectURL(r.Context(), shortCode)
	if err != nil {
		if err == errors.ErrURLNotFound {
			http.NotFound(w, r)
			return
		}
		h.respondError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}
