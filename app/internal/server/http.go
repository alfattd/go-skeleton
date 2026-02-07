package server

import (
	"net/http"

	"github.com/alfattd/skeleton/internal/platform/config"
	"github.com/alfattd/skeleton/internal/platform/monitor"
)

func New(cfg *config.Config) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", monitor.Health)
	mux.HandleFunc("/version", monitor.Version(cfg.ServiceName, cfg.ServiceVersion))
	mux.Handle("/metrics", monitor.MetricsHandler())

	handlerWithMetrics := MetricsMiddleware(mux)

	return &http.Server{
		Addr:    ":80",
		Handler: handlerWithMetrics,
	}
}
