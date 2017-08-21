package httpapi

import (
	"net/http"
	"time"

	"github.com/albert-widi/transaction_example/log"
	"github.com/albert-widi/transaction_example/service/httpapi/apiv1"
	"github.com/pressly/chi"
	"github.com/prometheus/client_golang/prometheus"
)

// Options of
type Config struct {
	ListenAddress string
}

// Handler struct
type Handler struct {
	config Config
	router *chi.Mux
	birth  time.Time
}

// New web handler
func New(conf Config) *Handler {
	r := chi.NewRouter()
	h := &Handler{
		config: conf,
		router: r,
		birth:  time.Now(),
	}
	// grouping the api
	r.Route("/api/v1", registerAPI())
	r.Handle("/metrics", prometheus.Handler())
	return h
}

func registerAPI() func(chi.Router) {
	return func(r chi.Router) {
		apiv1.New().Register(r)
	}
}

// Run web service
func (h *Handler) Run() error {
	log.Printf("Starting web service on %s", h.config.ListenAddress)
	return http.ListenAndServe(h.config.ListenAddress, h.router)
}
