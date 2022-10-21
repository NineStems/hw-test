package internalhttp

import (
	"net/http"
)

func (s *ServerHTTP) createMux() {
	mux := http.NewServeMux()
	mux.Handle("/", s.defaultRoute())
	s.handler = mux
}
