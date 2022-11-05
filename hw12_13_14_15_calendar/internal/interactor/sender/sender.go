package sender

import (
	"context"
	"time"

	v1 "github.com/hw-test/hw12_13_14_15_calendar/api/v1"
	"github.com/hw-test/hw12_13_14_15_calendar/common"
	"github.com/hw-test/hw12_13_14_15_calendar/internal/config"
	"github.com/hw-test/hw12_13_14_15_calendar/pkg/errors"
)

type Rabbit interface {
	Start() error
	Close()
	Read(ctx context.Context, bodies chan []byte) error
}

type App struct {
	logger common.Logger
	client v1.CalendarClient
	rabbit Rabbit
	cfg    config.Sender
}

func New(
	logger common.Logger,
	client v1.CalendarClient,
	rabbit Rabbit,
	cfg config.Sender,
) *App {
	return &App{
		logger: logger,
		client: client,
		rabbit: rabbit,
		cfg:    cfg,
	}
}

func (a *App) Start(ctx context.Context) error {
	if err := a.rabbit.Start(); err != nil {
		return errors.Wrap(err, "a.rabbit.Start")
	}
	defer a.rabbit.Close()

	bodies := make(chan []byte, 1)

	err := a.rabbit.Read(ctx, bodies)
	if err != nil {
		return errors.Wrap(err, "a.rabbit.Read")
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case body := <-bodies:
			a.logger.Infof("consume message %v", string(body))
		default:
			time.Sleep(a.cfg.Pause)
		}
	}

}
