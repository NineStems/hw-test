package domain

import (
	"errors"
	"time"

	"github.com/hw-test/hw12_13_14_15_calendar/internal/pkg/util"
)

const (
	TakeAllNotification = iota
	TakeDayPeriodNotification
	TakeWeekPeriodNotification
	TakeMonthPeriodNotification
)

// TODO: этот блок следует выделить в общий пакет бизнес-ошибок.

var (
	ErrNotDefinedPeriod = errors.New("not defined period for getting notifications")
	ErrDateBusy         = errors.New("selected date already busy")
	ErrNotFound         = errors.New("notification not found")
)

// Event календарное событие.
type Event struct {
	ID               string    // уникальный идентификатор события (можно воспользоваться UUID);
	OwnerID          int       // ИД пользователя, владельца события
	Title            string    // заголовок, короткий текст
	Date             time.Time // дата и время события
	DateEnd          time.Time // дата и время окончания
	DateNotification time.Time // за сколько времени высылать уведомление, записываем время, с которого слать уведомл.
	Description      string    // Описание события, длинный текст
}

func (e Event) NeedNotification(date time.Time) bool {
	if !e.DateNotification.IsZero() {
		conditionBefore := date.After(e.DateNotification) || date.Equal(e.DateNotification)
		conditionAfter := !date.Add(time.Minute * -5).After(e.DateNotification)
		return conditionBefore && conditionAfter

	}
	return (date.After(e.Date) || date.Equal(e.Date)) &&
		!date.Add(time.Minute*-5).After(e.Date)
}

// GetNotification возвращает уведомление на основании события.
func (e Event) GetNotification() Notification {
	return Notification{
		ID:      e.ID,
		OwnerID: e.OwnerID,
		Title:   e.Title,
		Date:    e.Date,
	}
}

// CompareDates проверят, что указанные даты заняты событием.
func (e Event) CompareDates(dateStart, dateEnd time.Time) bool {
	return util.CompareDateRange(e.Date, e.DateEnd, dateStart, dateEnd)
}

// Notification уведомление для отправки пользователю.
type Notification struct {
	ID      string    // ИД события
	OwnerID int       // пользователь, которому отправлять
	Title   string    // заголовок события
	Date    time.Time // дата события
}
