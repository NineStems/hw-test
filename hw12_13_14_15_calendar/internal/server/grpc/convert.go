package grpc

import (
	"time"

	v1 "github.com/hw-test/hw12_13_14_15_calendar/api/v1"
	"github.com/hw-test/hw12_13_14_15_calendar/domain"
	"github.com/hw-test/hw12_13_14_15_calendar/pkg/errors"
)

// CreateEventToDomain конвертирует данные запроса создания в доменную область.
func CreateEventToDomain(in *v1.CreateRequest) (*domain.Event, error) {
	event := &domain.Event{
		OwnerID:     int(in.OwnerId),
		Title:       in.Title,
		Description: in.Description,
	}

	var err error
	if in.Date != "" {
		event.Date, err = time.Parse(time.RFC3339, in.Date)
		if err != nil {
			return nil, errors.Wrap(err, "time.Parse")
		}
	}

	if in.DateEnd != "" {
		event.DateEnd, err = time.Parse(time.RFC3339, in.DateEnd)
		if err != nil {
			return nil, errors.Wrap(err, "time.Parse")
		}
	}

	if in.DateNotification != "" {
		event.DateNotification, err = time.Parse(time.RFC3339, in.DateNotification)
		if err != nil {
			return nil, errors.Wrap(err, "time.Parse")
		}
	}

	return event, nil
}

// UpdateEventToDomain конвертирует данные запроса обновления в доменную область.
func UpdateEventToDomain(in *v1.UpdateRequest) (*domain.Event, error) {
	event := &domain.Event{
		ID:          in.Id,
		OwnerID:     int(in.OwnerId),
		Title:       in.Title,
		Description: in.Description,
	}

	var err error
	if in.Date != "" {
		event.Date, err = time.Parse(time.RFC3339, in.Date)
		if err != nil {
			return nil, errors.Wrap(err, "time.Parse")
		}
	}

	if in.DateEnd != "" {
		event.DateEnd, err = time.Parse(time.RFC3339, in.DateEnd)
		if err != nil {
			return nil, errors.Wrap(err, "time.Parse")
		}
	}

	if in.DateNotification != "" {
		event.DateNotification, err = time.Parse(time.RFC3339, in.DateNotification)
		if err != nil {
			return nil, errors.Wrap(err, "time.Parse")
		}
	}

	return event, nil
}

func EventFromDomain(in *domain.Event) *v1.Event {
	return &v1.Event{
		Id:               in.ID,
		OwnerId:          int32(in.OwnerID),
		Title:            in.Title,
		Date:             in.Date.Format(time.RFC3339),
		DateEnd:          in.DateEnd.Format(time.RFC3339),
		DateNotification: in.DateNotification.Format(time.RFC3339),
		Description:      in.Description,
	}
}

func EventsFromDomain(in []domain.Event) []*v1.Event {
	list := make([]*v1.Event, 0, len(in))
	for i := range in {
		list = append(list, EventFromDomain(&in[i]))
	}
	return list
}

func EventToDomain(in *v1.Event) (*domain.Event, error) {
	date, err := time.Parse(time.RFC3339, in.GetDate())
	if err != nil {
		return nil, err
	}
	dateEnd, err := time.Parse(time.RFC3339, in.GetDateEnd())
	if err != nil {
		return nil, err
	}
	dateNotification, err := time.Parse(time.RFC3339, in.GetDateNotification())
	if err != nil {
		return nil, err
	}
	return &domain.Event{
		ID:               in.GetId(),
		OwnerID:          int(in.GetOwnerId()),
		Title:            in.GetTitle(),
		Date:             date,
		DateEnd:          dateEnd,
		DateNotification: dateNotification,
		Description:      in.GetDescription(),
	}, nil
}

func EventsToDomain(in []*v1.Event) ([]domain.Event, error) {
	list := make([]domain.Event, 0, len(in))
	for i := range in {
		event, err := EventToDomain(in[i])
		if err != nil {
			return nil, err
		}
		list = append(list, *event)
	}
	return list, nil
}
