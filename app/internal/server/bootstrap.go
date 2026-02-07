package server

import (
	"net/http"

	"github.com/alfattd/skeleton/internal/platform/config"
	"github.com/alfattd/skeleton/internal/platform/logger"
	"github.com/alfattd/skeleton/internal/platform/monitor"
)

func Run() (*config.Config, *http.Server) {
	cfg := config.Load()

	logger.New()

	monitor.Init()

	srv := New(cfg)

	return cfg, srv
}
