# HTTP-сервер "Календарь событий"

[Быстрый старт](https://github.com/andreyxaxa/calendar/tree/main?tab=readme-ov-file#%D0%B7%D0%B0%D0%BF%D1%83%D1%81%D0%BA)

## Обзор

- Документация API - Swagger - http://localhost:8080/swagger
- Конфиг - [config/config.go](https://github.com/andreyxaxa/calendar/blob/main/config/config.go). Читается из `.env` файла.
- Логгер - [pkg/logger/logger.go](https://github.com/andreyxaxa/calendar/blob/main/pkg/logger/logger.go). Интерфейс позволяет подменить логгер.
- Graceful shutdown - [internal/app/app.go](https://github.com/andreyxaxa/calendar/blob/main/internal/app/app.go).
- Удобная и гибкая конфигурация HTTP сервера - [pkg/httpserver/options.go](https://github.com/andreyxaxa/calendar/blob/main/pkg/httpserver/options.go).
  Позволяет конфигурировать сервер в конструкторе таким образом:
  ```go
  httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port))
  ```
- В слое хэндлеров применяется версионирование - [internal/controller/http/v1](https://github.com/andreyxaxa/calendar/tree/main/internal/controller/restapi/v1).
  Для версии v2 нужно будет просто добавить папку `restapi/v2` с таким же содержимым, в файле [internal/controller/restapi/router.go](https://github.com/andreyxaxa/calendar/blob/main/internal/controller/restapi/router.go) добавить строку:
  ```go
  {
      v1.NewEventsRoutes(apiV1Group, e, l)
  }

  {
      v2.NewEventsRoutes(apiV1Group, e, l)
  }
  ```

## Запуск

Клонируем репозиторий, выполняем:
```
make compose-up
```

Можно запустить без докера:
```
make run
```

## Тесты

Запустить тесты:
```
make test
```

## Прочие `make` команды
Зависимости:
```
make deps
```
docker compose down:
```
make compose-down
```

## API

### POST http://localhost:8080/v1/create_event
request:
```json
{
    "user_id": 1,
    "date": "2026-01-08",
    "text": "event"
}
```
response:
```json
{
    "result": {
        "user_id": 1,
        "uid": "bb52a762-f283-48ad-8cb5-cfe8e5bfa8eb",
        "date": "2026-01-08",
        "text": "event"
    }
}
```

### POST http://localhost:8080/v1/update_event
request:
```json
{
    "user_id": 1,
    "date": "2026-01-10",
    "text": "new text",
    "uid": "bb52a762-f283-48ad-8cb5-cfe8e5bfa8eb"
}
```
response:
```json
{
    "result": {
        "user_id": 1,
        "uid": "bb52a762-f283-48ad-8cb5-cfe8e5bfa8eb",
        "date": "2026-01-10",
        "text": "new text"
    }
}
```

### POST http://localhost:8080/v1/delete_event
request:
```json
{
    "user_id": 1,
    "uid": "bb52a762-f283-48ad-8cb5-cfe8e5bfa8eb"
}
```
response:
OK(200)

***После удаления не забудьте создать новое событие(я) для тестирования следующих методов.

### GET http://localhost:8080/v1/events_for_day?user_id=1&date=2026-01-08
response:
```json
[
    {
        "result": {
            "user_id": 1,
            "uid": "62dad1b6-9f83-4e2d-9047-aa5bf5641e11",
            "date": "2026-01-08",
            "text": "one more event"
        }
    },
    {
        "result": {
            "user_id": 1,
            "uid": "cee8027f-d1ba-424d-85ce-44ba201fd9d3",
            "date": "2026-01-08",
            "text": "событие"
        }
    }
]
```

### GET http://localhost:8080/v1/events_for_week?user_id=1&date=2026-01-05
response:
```json
[
    {
        "result": {
            "user_id": 1,
            "uid": "cee8027f-d1ba-424d-85ce-44ba201fd9d3",
            "date": "2026-01-08",
            "text": "событие"
        }
    },
    {
        "result": {
            "user_id": 1,
            "uid": "62dad1b6-9f83-4e2d-9047-aa5bf5641e11",
            "date": "2026-01-08",
            "text": "one more event"
        }
    }
]
```

### GET http://localhost:8080/v1/events_for_month?user_id=1&date=2026-01-01
response:
```json
[
    {
        "result": {
            "user_id": 1,
            "uid": "cee8027f-d1ba-424d-85ce-44ba201fd9d3",
            "date": "2026-01-08",
            "text": "событие"
        }
    },
    {
        "result": {
            "user_id": 1,
            "uid": "62dad1b6-9f83-4e2d-9047-aa5bf5641e11",
            "date": "2026-01-08",
            "text": "one more event"
        }
    }
]
```
