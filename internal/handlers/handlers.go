package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leesamuel423/url-shortener/internal/config"
	"github.com/leesamuel423/url-shortener/internal/service"
)

type Handler struct {
	urlService service.URLService
	config     *config.Config
}

func NewHandler(urlService service.URLService, cfg *config.Config) *Handler {
	return &Handler{
		urlService: urlService,
		config:     cfg,
	}
}

func (h *Handler) SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Middleware
	r.Use(h.loggingMiddleware)
	r.Use(h.corsMiddleware)

	// API routes
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/shorten", h.ShortenURL).Methods("POST")
	api.HandleFunc("/stats/{shortCode}", h.GetStats).Methods("GET")

	// Serve static files for specific paths
	r.HandleFunc("/", h.serveHome).Methods("GET")
	r.HandleFunc("/stats.html", h.serveStats).Methods("GET")

	// Redirect route - must come after static routes
	r.HandleFunc("/{shortCode}", h.Redirect).Methods("GET")

	return r
}

func (h *Handler) serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/static/index.html")
}

func (h *Handler) serveStats(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/static/stats.html")
}

func (h *Handler) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
