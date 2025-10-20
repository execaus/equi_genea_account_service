-- +goose Up
-- +goose StatementBegin
create table if not exists accounts(
    id uuid primary key, -- Уникальный идентификатор аккаунта
    email varchar(255) not null unique, -- Адрес электронной почты пользователя
    password varchar(255) not null, -- Хэш пароля пользователя
    created_at timestamp not null default now(), -- Дата и время создания аккаунта
    updated_at timestamp not null default now(), -- Дата и время последнего обновления аккаунта
    last_activity_at timestamp not null default now() -- Дата и время последней активности пользователя
);

comment on table accounts is 'Таблица хранения аккаунтов пользователей';

comment on column accounts.id is 'Уникальный идентификатор аккаунта';
comment on column accounts.email is 'Адрес электронной почты пользователя';
comment on column accounts.password is 'Хэш пароля пользователя';
comment on column accounts.created_at is 'Дата и время создания аккаунта';
comment on column accounts.updated_at is 'Дата и время последнего обновления аккаунта';
comment on column accounts.last_activity_at is 'Дата и время последней активности пользователя';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists accounts;
-- +goose StatementEnd