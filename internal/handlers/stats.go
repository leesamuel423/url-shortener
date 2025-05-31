package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leesamuel423/url-shortener/pkg/errors"
)

func (h *Handler) GetStats(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	stats, err := h.urlService.GetURLStats(r.Context(), shortCode)
	if err != nil {
		if err == errors.ErrURLNotFound {
			http.NotFound(w, r)
			return
		}
		h.respondError(w, http.StatusInternalServerError, "Failed to get stats")
		return
	}

	h.respondJSON(w, http.StatusOK, stats)
}
