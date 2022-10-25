package mocks

import (
	"context"
	"time"

	"github.com/hw-test/hw12_13_14_15_calendar/domain"
)

type MockApp struct{}

func (m *MockApp) CreateEvent(ctx context.Context, event *domain.Event) (string, error) {
	return "some-id", nil
}
func (m *MockApp) UpdateEvent(ctx context.Context, event *domain.Event) error { return nil }
func (m *MockApp) DeleteEvent(ctx context.Context, id string) error           { return nil }
func (m *MockApp) ReadEvents(ctx context.Context, date time.Time, condition int) ([]domain.Event, error) {
	return []domain.Event{
		{},
		{},
	}, nil
}
