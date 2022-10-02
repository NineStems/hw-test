package sqlstorage

import (
	"context"
	"time"

	"github.com/calendar/hw12_13_14_15_calendar/domain"
	"github.com/calendar/hw12_13_14_15_calendar/pkg/errors"
)

func (s *Storage) checkExists(ctx context.Context, id string) error {
	sql := `SELECT id FROM otus.notification WHERE id >= $1`
	cancel, rows, err := s.db.Query(ctx, sql, id)
	if err != nil {
		return errors.Wrap(err, "s.db.Query")
	}
	defer cancel()
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return errors.Wrap(err, "rows.Err")
	}

	if !rows.Next() {
		return domain.ErrNotFound
	}
	return nil
}

func (s *Storage) checkDateExists(ctx context.Context, dateStart, dateEnd time.Time) error {
	sql := `SELECT id FROM otus.notification WHERE date >= $1`
	args := []interface{}{dateStart}
	if !dateEnd.IsZero() {
		sql += " AND dateend <= $2"
		args = append(args, dateEnd)
	}

	cancel, rows, err := s.db.Query(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "s.db.Query")
	}
	defer cancel()
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return errors.Wrap(err, "rows.Err")
	}

	if rows.Next() {
		return domain.ErrDateBusy
	}

	return nil
}
