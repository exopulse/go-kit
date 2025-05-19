package httpd

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const (
	readHeaderTimeout = 5 * time.Second
	readTimeout       = 15 * time.Second
	idleTimeout       = 30 * time.Second
	writeTimeout      = 15 * time.Second
)

// Server implements an HTTP server.
type Server struct {
	srv      *http.Server
	listener net.Listener
}

// NewServer creates a new server.
func NewServer(cfg Config, handler http.Handler) (*Server, error) {
	ln, err := net.Listen("tcp", net.JoinHostPort(cfg.Interface, cfg.Port))
	if err != nil {
		return nil, errors.Wrap(err, "failed to bind HTTP server")
	}

	return &Server{
		srv: &http.Server{
			Handler:           handler,
			ReadHeaderTimeout: readHeaderTimeout,
			ReadTimeout:       readTimeout,
			WriteTimeout:      writeTimeout,
			IdleTimeout:       idleTimeout,
		},
		listener: ln,
	}, nil
}

// Run runs the server. It blocks until the server is stopped.
func (s *Server) Run() error {
	// ErrServerClosed is returned when the server is stopped.
	// ErrClosed is returned when the listener is closed.
	// We don't want to return an error in these cases.
	if err := s.srv.Serve(s.listener); err != nil && !errors.Is(err, http.ErrServerClosed) && !errors.Is(err, net.ErrClosed) {
		return errors.Wrap(err, "error serving")
	}

	return nil
}

// Stop stops the server gracefully. It blocks until all connections are closed.
// It returns an error if the server is already stopped or if the shutdown timeout is reached.
func (s *Server) Stop(ctx context.Context) error {
	if err := s.srv.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "server shutdown failed")
	}

	return nil
}

// Unbind unbinds the server from the listening address.
// Existing connections are not closed. New connections are rejected.
// It returns an error if the server is already unbound.
func (s *Server) Unbind() error {
	if err := s.listener.Close(); err != nil {
		return errors.Wrap(err, "failed to close listener")
	}

	return nil
}

// Address returns the server address.
func (s *Server) Address() string {
	return s.listener.Addr().String()
}
