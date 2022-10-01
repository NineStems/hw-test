package memorystorage

import (
	"context"
	"testing"
	"time"

	logmock "github.com/fixme_my_friend/hw12_13_14_15_calendar/common/mocks"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/domain"
)

func TestStorage(t *testing.T) { //nolint: gocognit
	log := logmock.MockLogger{}
	ctx := context.Background()
	storage := New(log)

	dateStart := time.Date(
		2022,
		06, //nolint: gofumpt
		28,
		20,
		30,
		50,
		0,
		time.UTC,
	)
	dateEnd := dateStart.Add(time.Hour * 1)

	t.Run("create", func(t *testing.T) {
		event := domain.Event{
			OwnerID:          10,
			Title:            "some title",
			Date:             dateStart,
			DateEnd:          dateEnd,
			DateNotification: 0,
			Description:      "some description",
		}
		id, err := storage.Create(ctx, &event)
		if err != nil {
			t.Error(err)
		}
		if len(id) == 0 {
			t.Errorf("created id is empty")
		}
	})

	t.Run("read", func(t *testing.T) {
		list := []int{
			domain.TakeAllNotification,
			domain.TakeDayPeriodNotification,
			domain.TakeWeekPeriodNotification,
			domain.TakeMonthPeriodNotification,
		}
		for _, condition := range list {
			events, err := storage.Read(ctx, dateStart, condition)
			if err != nil {
				t.Error(err)
			}
			if len(events) != 1 {
				t.Errorf("condition=%d, wait len of events=%d, but got=%d", condition, 1, len(events))
			}
		}
	})

	t.Run("update", func(t *testing.T) {
		events, err := storage.Read(ctx, dateStart, domain.TakeAllNotification)
		if err != nil {
			t.Error(err)
		}
		if len(events) != 1 {
			t.Errorf("condition=%d, wait len of events=%d, but got=%d", domain.TakeAllNotification, 1, len(events))
		}

		event := events[0]
		event.Title = "new title"
		event.OwnerID = 50
		err = storage.Update(ctx, &event)
		if err != nil {
			t.Error(err)
		}

		events, err = storage.Read(ctx, dateStart, domain.TakeAllNotification)
		if err != nil {
			t.Error(err)
		}
		if len(events) != 1 {
			t.Errorf("condition=%d, wait len of events=%d, but got=%d", domain.TakeAllNotification, 1, len(events))
		}
		if events[0].OwnerID != 50 && events[0].Title != "new title" {
			t.Errorf("operation update was bad")
		}
	})

	t.Run("delete", func(t *testing.T) {
		events, err := storage.Read(ctx, dateStart, domain.TakeAllNotification)
		if err != nil {
			t.Error(err)
		}
		if len(events) != 1 {
			t.Errorf("condition=%d, wait len of events=%d, but got=%d", domain.TakeAllNotification, 1, len(events))
		}

		err = storage.Delete(ctx, events[0].ID)
		if err != nil {
			t.Error(err)
		}

		events, err = storage.Read(ctx, dateStart, domain.TakeAllNotification)
		if err != nil {
			t.Error(err)
		}
		if len(events) != 0 {
			t.Errorf("operation delete was bad")
		}
	})
}
