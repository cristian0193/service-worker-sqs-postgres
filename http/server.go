package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const rootPrefix = "/service-template-golang"

// Server is an instance of Http Server for Rest endpoints.
type Server struct {
	server    *echo.Echo
	startedAt time.Time
	port      int
}

type healthResponse struct {
	Status    string    `json:"status"`
	StartedAt time.Time `json:"started_at"`
}

// NewServer creates an instance of Http Server.
func NewServer(port int) *Server {
	e := echo.New()

	// middleware
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// routes prefix
	path := e.Group(rootPrefix)

	server := &Server{server: e, startedAt: time.Now(), port: port}

	// healthcheck route
	path.GET("/health", server.health)

	return server
}

func (s *Server) health(c echo.Context) error {
	r := &healthResponse{Status: "OK", StartedAt: s.startedAt}
	return c.JSON(http.StatusOK, r)
}

// Start runs an http server.
func (s *Server) Start() error {
	port := fmt.Sprintf(":%v", s.port)

	if err := s.server.Start(port); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("server.Start: %w", err)
	}

	return nil
}

// Stop stops an http server.
func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server.Shutdown: %w", err)
	}

	return nil
}
