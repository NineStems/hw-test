package server

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/hw-test/hw12_13_14_15_calendar/common"
	"github.com/hw-test/hw12_13_14_15_calendar/domain"
	"github.com/hw-test/hw12_13_14_15_calendar/internal/config"
	internalhttp "github.com/hw-test/hw12_13_14_15_calendar/internal/server/http"
	"github.com/hw-test/hw12_13_14_15_calendar/internal/server/internalgrpc"
)

type Server struct {
	cfg  *config.Config
	log  common.Logger
	app  domain.Application
	ln   net.Listener
	rest internalhttp.ServerHTTP
	grpc internalgrpc.ServerGRPC
}

func NewServer(log common.Logger, cfg *config.Config, app domain.Application) *Server {
	return &Server{
		cfg: cfg,
		log: log,
		app: app,
	}
}

func (s *Server) Start(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	osSignals := make(chan os.Signal, 1)
	listenErr := make(chan error, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	if err := s.rest.Start(ctx, osSignals, listenErr); err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}

func StartHTTP(ctx context.Context) error {

	return nil
}

func StartGRPC(ctx context.Context) error {

	return nil
}
