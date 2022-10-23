package domain

import (
	"context"
	"time"
)

type Application interface {
	CreateEvent(ctx context.Context, event *Event) (string, error)
	UpdateEvent(ctx context.Context, event *Event) error
	DeleteEvent(ctx context.Context, id string) error
	ReadEvents(ctx context.Context, date time.Time, condition int) ([]Event, error)
}
