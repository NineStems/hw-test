package sqlstorage

import (
	"context"
	"time"

	"github.com/fixme_my_friend/hw12_13_14_15_calendar/domain"
)

type Storage struct { // TODO
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Create(ctx context.Context, event *domain.Event) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) Update(ctx context.Context, event *domain.Event) error {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) Read(ctx context.Context, date time.Time, condition int) ([]domain.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) ReadDay(ctx context.Context, day time.Time) ([]domain.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) ReadWeek(ctx context.Context, weekDay time.Time) ([]domain.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) ReadMonth(ctx context.Context, monthDay time.Time) ([]domain.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) Connect(ctx context.Context) error {
	// TODO
	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	// TODO
	return nil
}
