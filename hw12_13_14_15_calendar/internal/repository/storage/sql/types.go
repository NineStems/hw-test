package sqlstorage

import (
	"database/sql"
	"time"
)

// Event календарное событие.
type Event struct {
	ID               string         // уникальный идентификатор события (можно воспользоваться UUID);
	OwnerID          int            // ИД пользователя, владельца события
	Title            string         // заголовок, короткий текст
	Date             time.Time      // дата и время события
	DateEnd          time.Time      // Дата и время окончания)
	DateNotification sql.NullTime   //За сколько времени высылать уведомление, опционально
	Description      sql.NullString // Описание события, длинный текст, опционально
}
