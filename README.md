## Переменные окружения

Для корректной работы проекта необходимо установить следующие переменные окружения:

```env
GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://admin:admin@localhost:5432/admin_db
GOOSE_MIGRATION_DIR=./migrations
DATABASE_PASSWORD=1234
```

### Описание переменных

* `GOOSE_DRIVER` — драйвер базы данных, используемый Goose (например, `postgres`).
* `GOOSE_DBSTRING` — строка подключения к базе данных в формате `postgres://user:password@host:port/dbname`.
* `GOOSE_MIGRATION_DIR` — путь к директории с миграциями базы данных.
* `DATABASE_PASSWORD` — пароль для базы данных.

# Конфигурация

Конфигурация приложения задаётся в файле `config/config.yaml`. В этом файле можно указать параметры сервера и базы данных.

## Пример файла конфигурации

```yaml
server:
  port: "8080"

database:
  host: "localhost"
  port: 5432
  user: "dbuser"
  name: "dbname"
```

## Описание параметров

- `server.port` — порт, на котором будет запущен сервер.
- `database.host` — адрес сервера базы данных.
- `database.port` — порт базы данных.
- `database.user` — имя пользователя базы данных.
- `database.name` — имя базы данных.
