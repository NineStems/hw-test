package sqlstorage

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/hw-test/hw12_13_14_15_calendar/domain"
	"github.com/hw-test/hw12_13_14_15_calendar/internal/pkg/util"
	"github.com/hw-test/hw12_13_14_15_calendar/pkg/errors"
)

type Row interface {
	Scan(dest ...interface{}) (err error)
	Err() error
}

type Rows interface {
	Scan(dest ...interface{}) (err error)
	Next() bool
	Err() error
	Close() error
}

type DB interface {
	Open(ctx context.Context) error
	Close() error
	Query(ctx context.Context, sql string, args ...interface{}) (context.CancelFunc, *sql.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) (context.CancelFunc, *sql.Row)
	Exec(ctx context.Context, sql string, args ...interface{}) error
}

type Storage struct {
	db DB
}

func New(db DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) Create(ctx context.Context, event *domain.Event) (string, error) {
	if err := s.checkDateExists(ctx, event.Date, event.DateEnd); err != nil {
		return "", errors.Wrap(err, "s.checkDateExists")
	}

	id, err := util.GenerateUUID()
	if err != nil {
		return "", errors.Wrap(err, "util.GenerateUUID")
	}
	event.ID = id

	dEvent := eventFromDomain(event)
	query := `INSERT INTO 
			otus.notification (	id, ownerid, title, date, dateend, datenotification, description) 
			VALUES ($1,$2,$3,$4,$5,$6,$7)`

	err = s.db.Exec(
		ctx,
		query,
		dEvent.ID,
		dEvent.OwnerID,
		dEvent.Title,
		dEvent.Date,
		dEvent.DateEnd,
		dEvent.DateNotification,
		dEvent.Description,
	)
	if err != nil {
		return "", errors.Wrap(err, "s.db.Exec")
	}

	return event.ID, nil
}

func (s *Storage) Update(ctx context.Context, event *domain.Event) error {
	if err := s.checkExists(ctx, event.ID); err != nil {
		return errors.Wrap(err, "s.checkExists")
	}

	if err := s.checkDateExists(ctx, event.Date, event.DateEnd); err != nil {
		return errors.Wrap(err, "s.checkDateExists")
	}

	dEvent := eventFromDomain(event)
	query := `UPDATE otus.notification 
	SET ownerid=$2, title=$3, date=$4, dateend=$5, datenotification=$6, description=$7
	WHERE  id=$1`

	err := s.db.Exec(
		ctx,
		query,
		dEvent.ID,
		dEvent.OwnerID,
		dEvent.Title,
		dEvent.Date,
		dEvent.DateEnd,
		dEvent.DateNotification,
		dEvent.Description,
	)
	if err != nil {
		return errors.Wrap(err, "s.db.Exec")
	}

	return nil
}

func (s *Storage) Delete(ctx context.Context, ids []string) error {
	for _, id := range ids {
		if err := s.checkExists(ctx, id); err != nil {
			return errors.Wrap(err, "s.checkExists")
		}
	}

	query := fmt.Sprintf("DELETE FROM otus.notification WHERE id IN ('%s')", strings.Join(ids, "', '"))

	if err := s.db.Exec(ctx, query); err != nil {
		return errors.Wrap(err, "s.db.Exec")
	}

	return nil
}

func (s *Storage) Read(ctx context.Context, date time.Time, condition int) ([]domain.Event, error) {
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

	query := `SELECT id, ownerid, title, date, dateend, datenotification, description 
				FROM otus.notification`
	var args []interface{}
	if !start.IsZero() && !end.IsZero() {
		query += ` WHERE date >= $1 AND dateend <= $2`
		args = append(args, start, end)
	}

	cancel, rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "s.db.Query")
	}
	defer cancel()
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "rows.Err")
	}

	var events []Event
	for rows.Next() {
		var event Event
		if err = rows.Scan(
			&event.ID,
			&event.OwnerID,
			&event.Title,
			&event.Date,
			&event.DateEnd,
			&event.DateNotification,
			&event.Description,
		); err != nil {
			return nil, errors.Wrap(err, "rows.Next")
		}
		events = append(events, event)
	}

	return eventsFromDomain(events), nil
}

func (s *Storage) Connect(ctx context.Context) error {
	if err := s.db.Open(ctx); err != nil {
		return errors.Wrap(err, "s.db.Open")
	}
	return nil
}

func (s *Storage) Close() error {
	if err := s.db.Close(); err != nil {
		return errors.Wrap(err, "s.db.Close")
	}
	return nil
}
