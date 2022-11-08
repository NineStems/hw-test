package grpc

import (
	"context"
	"log"
	"net"
	"os"

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

func (s *ServerGRPC) Start(ctx context.Context, osSignals chan os.Signal, listenErr chan error) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var err error
	s.ln, err = net.Listen("tcp", net.JoinHostPort(s.cfg.Grpc.Host, s.cfg.Grpc.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	gs := grpc.NewServer()
	v1.RegisterCalendarServer(gs, s)

	go func() {
		s.log.Infof("grpc server started on %s", s.ln.Addr())
		listenErr <- gs.Serve(s.ln)
	}()

	for {
		select {
		case err = <-listenErr:
			s.log.Errorf("grpc server stopped error:v", err)
			return err
		case <-osSignals:
			s.log.Info("grpc server stopped")
			gs.GracefulStop()
			s.log.Info("grpc server exited properly")
		}
	}
}
