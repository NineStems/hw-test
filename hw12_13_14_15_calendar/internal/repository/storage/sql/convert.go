package sqlstorage

import (
	"github.com/hw-test/hw12_13_14_15_calendar/domain"
)

// eventFromDomain конвертирует доменный тип в тип репозитория.
func eventFromDomain(in *domain.Event) *Event {
	return &Event{
		ID:               in.ID,
		OwnerID:          in.OwnerID,
		Title:            in.Title,
		Date:             in.Date,
		DateEnd:          in.DateEnd,
		DateNotification: TimeToNull(in.DateNotification),
		Description:      StringToNull(in.Description),
	}
}

// eventToDomain конвертирует тип репозитория к доменному.
func eventToDomain(in *Event) *domain.Event {
	return &domain.Event{
		ID:               in.ID,
		OwnerID:          in.OwnerID,
		Title:            in.Title,
		Date:             in.Date,
		DateEnd:          in.DateEnd,
		DateNotification: in.DateNotification.Time,
		Description:      in.Description.String,
	}
}

// eventsFromDomain конвертирует тип репозитория к доменному.
func eventsFromDomain(in []Event) []domain.Event {
	list := make([]domain.Event, 0, len(in))
	for i := range in {
		list = append(list, *eventToDomain(&in[i]))
	}
	return list
}
