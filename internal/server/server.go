package server

import (
	"context"
)

type Server struct {
}

func New() *Server {
	return new(Server)
}

func (s *Server) Ping(_ context.Context) int {
	return 42
}
