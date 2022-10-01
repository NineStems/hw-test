package app

import (
	"context"
	"time"

	"github.com/fixme_my_friend/hw12_13_14_15_calendar/common"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/domain"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/pkg/errors"
)

type App struct {
	logger common.Logger
	db     Storage
}

type Storage interface {
	Create(ctx context.Context, event *domain.Event) (string, error)
	Update(ctx context.Context, event *domain.Event) error
	Delete(ctx context.Context, id string) error
	Read(ctx context.Context, date time.Time, condition int) ([]domain.Event, error)
}

func New(logger common.Logger, storage Storage) *App {
	return &App{
		logger: nil, // TODO: на слое usecase нет необходимости добавлять логгер.
		db:     storage,
	}
}

// CreateEvent создаёт уведомление.
func (a *App) CreateEvent(ctx context.Context, event *domain.Event) (string, error) {
	id, err := a.db.Create(ctx, event)
	if err != nil {
		return "", errors.Wrap(err, "a.db.Create")
	}
	return id, nil
}

// UpdateEvent обновляет уведомление.
func (a *App) UpdateEvent(ctx context.Context, event *domain.Event) error {
	err := a.db.Update(ctx, event)
	if err != nil {
		return errors.Wrap(err, "a.db.Update")
	}
	return nil
}

// DeleteEvent удаляет уведомление.
func (a *App) DeleteEvent(ctx context.Context, id string) error {
	err := a.db.Delete(ctx, id)
	if err != nil {
		return errors.Wrap(err, "a.db.Delete")
	}
	return nil
}

// ReadEvents получает уведомления по условию.
func (a *App) ReadEvents(ctx context.Context, date time.Time, condition int) ([]domain.Notification, error) {
	if date.IsZero() {
		condition = domain.TakeAllNotification
	}
	events, err := a.db.Read(ctx, date, condition)
	if err != nil {
		return nil, errors.Wrap(err, "a.db.Read")
	}

	if len(events) == 0 {
		return nil, nil
	}

	notification := make([]domain.Notification, 0, len(events))
	for i := range events {
		notification = append(notification, events[i].GetNotification())
	}

	return notification, nil
}
