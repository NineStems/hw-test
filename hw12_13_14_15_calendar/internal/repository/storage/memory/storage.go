package memorystorage

import (
	"context"
	"sync"
	"time"

	"github.com/fixme_my_friend/hw12_13_14_15_calendar/common"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/domain"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/pkg/util"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/pkg/errors"
)

const (
	actionCreate = "CREATE"
	actionUpdate = "UPDATE"
	actionDelete = "DELETE"
	actionRead   = "READ"

	ctxEventId   = "event-id"
	ctxDateRead  = "date-read"
	ctxCondition = "condition-read"
	ctxDateStart = "date-start"
	ctxDateEnd   = "date-end"
)

type Storage struct {
	data map[string]Event
	mu   sync.RWMutex
	log  common.Logger
}

func (s *Storage) Create(ctx context.Context, event *domain.Event) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id, err := util.GenerateUUID()
	if err != nil {
		return "", errors.Wrap(err, "util.GenerateUUID")
	}
	event.ID = id

	err = s.checkTimeExist(ctx, event)
	if err != nil {
		return "", errors.Wrap(err, "s.checkTimeExist")
	}

	sEvent := eventFromDomain(event)
	dateEnd := &sEvent.DateEnd
	if sEvent.DateEnd.IsZero() {
		dateEnd = nil
	}
	s.log.Debugw(actionCreate, ctxEventId, id, ctxDateStart, sEvent.Date, ctxDateEnd, dateEnd)

	s.data[sEvent.ID] = *sEvent
	return sEvent.ID, nil
}

func (s *Storage) Update(ctx context.Context, event *domain.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := s.checkTimeExist(ctx, event)
	if err != nil {
		return errors.Wrap(err, "s.checkTimeExist")
	}

	dateEnd := &event.DateEnd
	if event.DateEnd.IsZero() {
		dateEnd = nil
	}
	s.log.Debugw(actionUpdate, ctxEventId, event.ID, ctxDateStart, event.Date, ctxDateEnd, dateEnd)
	s.data[event.ID] = *eventFromDomain(event)
	return nil
}

func (s *Storage) Delete(ctx context.Context, id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, id)
	s.log.Debugw(actionDelete, ctxEventId, id)
	return nil
}

func (s *Storage) Read(ctx context.Context, date time.Time, condition int) ([]domain.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var start, end time.Time
	switch condition {
	case domain.TakeAllNotification:
		start, end = time.Time{}, time.Time{}
	case domain.TakeDayPeriodNotification:
		start, end = util.StartDateDay(date), util.EndDateDay(date)
	case domain.TakeWeekPeriodNotification:
		start, end = util.StartDateDay(date), util.EndDateDay(date)
	case domain.TakeMonthPeriodNotification:
		start, end = util.StartDateDay(date), util.EndDateDay(date)
	default:
		return nil, domain.ErrNotDefinedPeriod
	}

	events, err := s.read(ctx, start, end)
	if err != nil {
		return nil, errors.Wrap(err, "s.read")
	}

	s.log.Debugw(actionRead, ctxDateRead, date, ctxCondition, condition)
	return eventsToDomain(events), nil
}

func New(log common.Logger) *Storage {
	return &Storage{
		data: map[string]Event{},
		mu:   sync.RWMutex{},
		log:  log,
	}
}
