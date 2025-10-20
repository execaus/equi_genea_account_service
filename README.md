## Переменные окружения

Для корректной работы проекта необходимо установить следующие переменные окружения:

```env
GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://admin:admin@localhost:5432/admin_db
GOOSE_MIGRATION_DIR=./migrations
```

### Описание переменных

* `GOOSE_DRIVER` — драйвер базы данных, используемый Goose (например, `postgres`).
* `GOOSE_DBSTRING` — строка подключения к базе данных в формате `postgres://user:password@host:port/dbname`.
* `GOOSE_MIGRATION_DIR` — путь к директории с миграциями базы данных.
