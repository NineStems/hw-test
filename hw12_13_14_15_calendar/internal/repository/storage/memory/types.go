package memorystorage

import (
	"time"
)

// Event календартное событие.
type Event struct {
	ID               string        // уникальный идентификатор события (можно воспользоваться UUID);
	OwnerID          int           // ИД пользователя, владельца события
	Title            string        // заголовок, короткий текст
	Date             time.Time     // дата и время события
	DateEnd          time.Time     // Дата и время окончания)
	DateNotification time.Duration //За сколько времени высылать уведомление, опционально
	Description      string        // Описание события, длинный текст, опционально
}
