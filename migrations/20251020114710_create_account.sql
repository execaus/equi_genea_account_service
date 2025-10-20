-- +goose Up
-- +goose StatementBegin
create table if not exists accounts(
    id uuid primary key,
    email varchar(255) not null unique,
    password varchar(255) not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    last_activity_at timestamp not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists accounts;
-- +goose StatementEnd
