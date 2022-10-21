package internalhttp

import (
	"context"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	v1 "github.com/hw-test/hw12_13_14_15_calendar/api/v1"
	"github.com/hw-test/hw12_13_14_15_calendar/common"
	"github.com/hw-test/hw12_13_14_15_calendar/domain"
	"github.com/hw-test/hw12_13_14_15_calendar/internal/config"
)

type ServerHTTP struct {
	v1.UnimplementedCalendarServer
	cfg     *config.Server
	log     common.Logger
	app     domain.Application
	ln      net.Listener
	handler http.Handler
}

func NewServer(log common.Logger, cfg *config.Server, app domain.Application) *ServerHTTP {
	return &ServerHTTP{
		cfg: cfg,
		log: log,
		app: app,
	}
}

func (s *ServerHTTP) Start(ctx context.Context, osSignals chan os.Signal, listenErr chan error) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var err error
	s.ln, err = net.Listen("tcp4", s.cfg.Http.Host+":"+s.cfg.Http.Port)
	if err != nil {
		return err
	}

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = v1.RegisterCalendarHandlerFromEndpoint(ctx, mux, s.cfg.Grpc.Host+":"+s.cfg.Grpc.Port, opts)
	if err != nil {
		return err
	}

	server := http.Server{
		Handler: s.loggingMiddleware(mux),
	}

	go func() {
		s.log.Infof("rest server started on %s", s.ln.Addr())
		listenErr <- server.Serve(s.ln)
	}()

	for {
		select {
		case err = <-listenErr:
			return err
		case <-osSignals:
			s.log.Info("rest server stopped")
			if err := server.Shutdown(ctx); err != nil {
				return err
			}
			s.log.Info("rest server exited properly")
		}
	}
}
