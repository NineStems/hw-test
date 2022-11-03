package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/hw-test/hw12_13_14_15_calendar/common"
	"github.com/hw-test/hw12_13_14_15_calendar/domain"
	"github.com/hw-test/hw12_13_14_15_calendar/internal/config"
	internalgrpc "github.com/hw-test/hw12_13_14_15_calendar/internal/server/grpc"
	internalhttp "github.com/hw-test/hw12_13_14_15_calendar/internal/server/http"
)

type Server struct {
	cfg  *config.Config
	log  common.Logger
	rest *internalhttp.ServerHTTP
	grpc *internalgrpc.ServerGRPC
}

func NewServer(log common.Logger, cfg *config.Config, app domain.Application) *Server {
	restNode := internalhttp.NewServer(log, &cfg.Server, app)
	grpcNode := internalgrpc.NewServer(log, &cfg.Server, app)
	return &Server{
		cfg:  cfg,
		log:  log,
		rest: restNode,
		grpc: grpcNode,
	}
}

func (s *Server) Start(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	osSignals := make(chan os.Signal, 1)
	listenErr := make(chan error, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.rest.Start(ctx, osSignals, listenErr); err != nil {
			listenErr <- err
		}
	}()

	go func() {
		if err := s.grpc.Start(ctx, osSignals, listenErr); err != nil {
			listenErr <- err
		}
	}()

	<-ctx.Done()
	return nil
}
