package internalhttp

import (
	"context"
	"net"

	"github.com/fixme_my_friend/hw12_13_14_15_calendar/common"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/config"
)

type Server struct { // TODO
	logger common.Logger
	cfg    config.Config
	log    common.Logger
	ln     net.Listener
}

type Application interface { // TODO
}

func NewServer(logger common.Logger, app Application) *Server {
	return &Server{
		logger: logger,
	}
}

func (s *Server) Start(ctx context.Context) error {
	// TODO
	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	// TODO
	return nil
}

// TODO
