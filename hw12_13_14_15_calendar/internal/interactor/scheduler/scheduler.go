package scheduler

import (
	"context"
	"time"

	v1 "github.com/hw-test/hw12_13_14_15_calendar/api/v1"
	"github.com/hw-test/hw12_13_14_15_calendar/common"
	"github.com/hw-test/hw12_13_14_15_calendar/internal/config"
	"github.com/hw-test/hw12_13_14_15_calendar/internal/server/grpc"
	"github.com/hw-test/hw12_13_14_15_calendar/pkg/errors"
)

type Rabbit interface {
	Start() error
	Close()
	Publish(ctx context.Context, message interface{}) error
}

type App struct {
	logger    common.Logger
	client    v1.CalendarClient
	rabbit    Rabbit
	scheduler config.Scheduler
}

func New(
	logger common.Logger,
	client v1.CalendarClient,
	rabbit Rabbit,
	scheduler config.Scheduler,
) *App {
	return &App{
		logger:    logger,
		client:    client,
		rabbit:    rabbit,
		scheduler: scheduler,
	}
}

func (a *App) Start(ctx context.Context) error {
	if err := a.rabbit.Start(); err != nil {
		return errors.Wrap(err, "a.rabbit.Start")
	}
	defer a.rabbit.Close()

	for {
		time.Sleep(a.scheduler.Pause)
		select {
		case <-ctx.Done():
			return nil
		default:
			a.clearOldEvents(ctx)
			a.readAndPushEvents(ctx)
		}
	}
}

func (a *App) readAndPushEvents(ctx context.Context) {
	now := time.Now()
	date := now.Format(time.RFC3339)
	condition := int32(1)
	request := v1.ReadRequest{
		Date:      date,
		Condition: condition,
	}
	response, err := a.client.ReadEvents(ctx, &request)
	if err != nil {
		return
	}
	if len(response.Events) == 0 {
		a.logger.Infof("Not found any events for time='%v' condition='%v'", date, condition)
		return
	}

	events := grpc.EventsToDomain(response.Events)
	for _, event := range events {
		if !event.NeedNotification(now) {
			continue
		}
		notification := event.GetNotification()
		publishCtx := context.Background()
		if err = a.rabbit.Publish(publishCtx, &notification); err != nil {
			a.logger.Errorf("Publish notification: error='%v'", err)
			continue
		}
	}
}

func (a *App) clearOldEvents(ctx context.Context) {
	now := time.Now()
	request := v1.ReadRequest{
		Date:      "2006-01-02T15:04:05+03:00",
		Condition: 0,
	}
	response, err := a.client.ReadEvents(ctx, &request)
	if err != nil {
		a.logger.Errorf("DeleteEvent: error='%v'", err)
		return
	}
	if len(response.Events) == 0 {
		return
	}
	events := grpc.EventsToDomain(response.Events)
	for i := range events {
		if !events[i].Date.Before(now.AddDate(-1, 0, 0)) {
			continue
		}
		if _, err = a.client.DeleteEvent(ctx, &v1.DeleteRequest{Id: events[i].ID}); err != nil {
			a.logger.Errorf("DeleteEvent: error='%v'", err)
		}
	}

}
