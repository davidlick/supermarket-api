package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

// server holds configurations and services used by the HTTP server and handlers.
type server struct {
	port        int
	logger      *logrus.Logger
	environment string
	server      *http.Server
}

// NewServer initializes a new server with the required configurations.
func NewServer(port int, logger *logrus.Logger, environment string) *server {
	return &server{
		port:        port,
		logger:      logger,
		environment: environment,
		server: &http.Server{
			Addr:         fmt.Sprintf(":%d", port),
			ReadTimeout:  60 * time.Second,
			WriteTimeout: 60 * time.Second,
		},
	}
}

// Run builds the routes and starts the server listening on the configured port.
func (s *server) Run() error {
	s.server.Handler = s.buildRoutes()
	s.logger.Infof("starting server on port: %d", s.port)
	return s.server.ListenAndServe()
}

// Shutdown calls Shutdown on the http.Server which attempts to gracefully shutdown. If the context deadline is exceeded any cotnext errors are returned.
func (s *server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

// buildRoutes configures a router with middleware, headers, and routes.
func (s *server) buildRoutes() http.Handler {
	r := s.configureRouter()

	r.Get("/health", s.handleHealth)
	r.Get("/ready", s.handleReadiness)

	r.Group(func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			// r.Group(s.produceGroup)
		})
	})

	return r
}

// configureRouter returns a configured chi.Router with sane defaults.
func (s *server) configureRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(setResponseHeaders(map[string]string{
		"Content-Type":                 "application/json",
		"Allow-Access-Control-Origin":  "*",
		"Allow-Access-Control-Method":  "OPTIONS, GET, POST, DELETE",
		"Allow-Access-Control-Headers": "Origin, Content-Type, Accept",
		"Access-Control-Max-Age":       "600",
	}))

	return r
}
