package memorystorage

import (
	"time"

	"github.com/hw-test/hw12_13_14_15_calendar/domain"
	"github.com/hw-test/hw12_13_14_15_calendar/internal/pkg/util"
)

// checkTimeExist проверяет, что событие можно создать на указанном времени.
func (s *Storage) checkTimeExist(event *domain.Event) error {
	events := make([]Event, 0, len(s.data))
	for _, value := range s.data {
		if event.OwnerID == value.OwnerID {
			events = append(events, value)
		}
	}

	for i := range events {
		if event.CompareDates(events[i].Date, events[i].DateEnd) {
			return domain.ErrDateBusy
		}
	}

	return nil
}

// read вычитывает события на основании временных ограничений.
func (s *Storage) read(start, end time.Time) ([]Event, error) { //nolint:unparam
	list := make([]Event, 0, len(s.data))
	for _, value := range s.data {
		if (start.IsZero() && end.IsZero()) || util.CompareDateRange(value.Date, value.DateEnd, start, end) {
			list = append(list, value)
		}
	}
	return list, nil
}
