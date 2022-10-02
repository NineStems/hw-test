create table otus.notification
(
    ID               varchar      not null,
    OwnerID          integer      not null,
    Title            varchar(100) not null,
    Date             timestamp    not null,
    DateEnd          timestamp    not null,
    DateNotification timestamp,
    Description      varchar(1000)
);

comment on table otus.notification is 'таблица уведомлений';
comment on column otus.notification.ID is 'уникальный идентификатор события';
comment on column otus.notification.OwnerID is 'ИД пользователя, владельца события';
comment on column otus.notification.Title is 'заголовок, короткий текст';
comment on column otus.notification.Date is 'дата и время события';
comment on column otus.notification.DateEnd is 'дата и время окончания';
comment on column otus.notification.DateNotification is 'за сколько времени высылать уведомление';
comment on column otus.notification.Description is 'описание события, длинный текст';