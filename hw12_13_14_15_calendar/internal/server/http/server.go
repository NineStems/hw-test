package internalhttp

import (
	"context"
	"net"
	"time"

	"github.com/fixme_my_friend/hw12_13_14_15_calendar/common"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/domain"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/config"
)

type Server struct { // TODO
	cfg *config.Config
	log common.Logger
	ln  net.Listener
	app Application
}

type Application interface { // TODO
	CreateEvent(ctx context.Context, notification *domain.Event) (string, error)
	UpdateEvent(ctx context.Context, notification *domain.Event) error
	DeleteEvent(ctx context.Context, id string) error
	ReadEvents(ctx context.Context, date time.Time, condition int) ([]domain.Notification, error)
}

func NewServer(log common.Logger, app Application, cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
		log: log,
		app: app,
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
