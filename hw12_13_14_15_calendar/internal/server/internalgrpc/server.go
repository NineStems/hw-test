package internalgrpc

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	v1 "github.com/hw-test/hw12_13_14_15_calendar/api/v1"
	"github.com/hw-test/hw12_13_14_15_calendar/common"
	"github.com/hw-test/hw12_13_14_15_calendar/domain"
	"github.com/hw-test/hw12_13_14_15_calendar/internal/config"
)

type ServerGRPC struct {
	v1.UnimplementedCalendarServer
	cfg *config.Server
	log common.Logger
	app domain.Application
	ln  net.Listener
}

func NewServer(log common.Logger, cfg *config.Server, app domain.Application) *ServerGRPC {
	return &ServerGRPC{
		cfg: cfg,
		log: log,
		app: app,
	}
}

func (s *ServerGRPC) Start(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	lis, err := net.Listen("tcp", s.cfg.Grpc.Host+":"+s.cfg.Grpc.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	gs := grpc.NewServer()
	v1.RegisterCalendarServer(gs, s)

	if err := gs.Serve(lis); err != nil {
		return err
	}

	return nil
}