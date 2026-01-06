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
