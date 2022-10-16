package sqlstorage

import (
	"database/sql"
	"time"
)

// TimeToNull преобразует time.Time в sql.NullTime.
func TimeToNull(value time.Time) sql.NullTime {
	return sql.NullTime{Time: value, Valid: !value.IsZero()}
}

// StringToNull преобразует string в sql.NullString.
func StringToNull(value string) sql.NullString {
	return sql.NullString{String: value, Valid: len(value) > 0}
}
