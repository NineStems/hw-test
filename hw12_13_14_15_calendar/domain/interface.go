package domain

import (
	"context"
	"time"
)

type Application interface {
	CreateEvent(ctx context.Context, notification *Event) (string, error)
	UpdateEvent(ctx context.Context, notification *Event) error
	DeleteEvent(ctx context.Context, id string) error
	ReadEvents(ctx context.Context, date time.Time, condition int) ([]Notification, error)
}
