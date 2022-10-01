package util

import (
	"time"

	"github.com/google/uuid"
)

// GenerateUUID генерирует уникальный идентификатор.
func GenerateUUID() (string, error) {
	id, err := uuid.NewUUID()
	return id.String(), err
}

// StartDateDay преобразует дату до начала дня.
func StartDateDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// EndDateDay преобразует дату до конца дня.
func EndDateDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, t.Location())
}

// StartDateWeek преобразует дату до начала недели.
func StartDateWeek(t time.Time) time.Time {
	weekday := time.Duration(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	year, month, day := t.Date()
	currentZeroDay := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return currentZeroDay.Add(-1 * (weekday - 1) * 24 * time.Hour)
}

// EndDateWeek преобразует дату до конца недели.
func EndDateWeek(t time.Time) time.Time {
	weekday := time.Duration(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	weekday = 7 - weekday
	year, month, day := t.Date()
	currentZeroDay := time.Date(year, month, day, 23, 59, 59, 0, time.Local)
	return currentZeroDay.Add(1 * (weekday) * 24 * time.Hour)
}

// StartDateMonth преобразует дату до начала месяца.
func StartDateMonth(date time.Time) time.Time {
	return StartDateDay(date.AddDate(0, 0, -date.Day()+1))
}

// EndDateMonth преобразует дату до окончания месяца.
func EndDateMonth(date time.Time) time.Time {
	return StartDateDay(date.AddDate(0, 1, -date.Day()))
}

// CompareDateRange сравнивает диапазон дат.
// Первая группа - динамические даты, которые ищем.
// Вторая группа - статические даты, среди которых ищем.
func CompareDateRange(start, end, baseStart, baseEnd time.Time) bool {
	startCondition := start.After(baseStart) || start.Equal(baseStart)
	endCondition := !baseEnd.IsZero() && (end.Before(baseEnd) || end.Equal(baseEnd))
	return startCondition && endCondition
}
