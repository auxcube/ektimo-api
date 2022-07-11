package httpserver

import (
	"fmt"
	"net"
	"time"
)

type Option func(*Server)

// Port overrides the default port
func Port(port int) Option {
	return func(s *Server) {
		s.server.Addr = net.JoinHostPort("", fmt.Sprint(port))
	}
}

// ReadTimeout overrides the default read timeout
func ReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.ReadTimeout = timeout
	}
}

// WriteTimeout overrides the default write timeout
func WriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.WriteTimeout = timeout
	}
}

// ShutdownTimeout overrides the default shutdown timeout
func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}
