package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/andreyxaxa/calendar/config"
	"github.com/andreyxaxa/calendar/internal/controller/restapi"
	"github.com/andreyxaxa/calendar/internal/repo/inmemory"
	"github.com/andreyxaxa/calendar/internal/usecase/events"
	"github.com/andreyxaxa/calendar/pkg/httpserver"
	"github.com/andreyxaxa/calendar/pkg/logger"
)

// Run -.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	inmem := inmemory.New()

	// Use-Case
	eventsUseCase := events.New(inmem)

	// HTTP Server
	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port))
	restapi.NewRouter(httpServer.App, cfg, eventsUseCase, l)

	// Start server
	httpServer.Start()

	// Waiting Signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: %s", s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	err := httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
