package internalhttp

import (
	"net/http"
)

func (s *Server) CreateMux() {
	mux := http.NewServeMux()
	mux.Handle("/", s.loggingMiddleware(s.defaultRoute()))
	s.handler = mux
}
