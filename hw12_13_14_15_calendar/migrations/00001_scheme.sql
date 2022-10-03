-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SCHEMA otus;
CREATE TABLE otus.notification
(
    ID               varchar      not null,
    OwnerID          integer      not null,
    Title            varchar(100) not null,
    Date             timestamp    not null,
    DateEnd          timestamp    not null,
    DateNotification timestamp,
    Description      varchar(1000)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE otus.notification;
DROP SCHEMA otus CASCADE ;
-- +goose StatementEnd
