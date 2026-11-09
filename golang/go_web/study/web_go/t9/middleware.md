---
tags:
  - go
  - web
  - fiber
  - course
  - t9
up: "[[MOC Go web курс]]"
---

> [!info] Раздел: [[MOC Go web курс]]

## Идея

Middleware — обработчик, который выполняется до (и/или после) основного хэндлера: логирование, лимиты, request id, авторизация. В Fiber подключается через `webApp.Use(...)`. Порядок `Use` важен: middleware срабатывают в порядке регистрации и применяются к роутам, зарегистрированным после них.

## Логгер (theory/0)

```go
webApp.Use(logger.New(logger.Config{
    Format:     "${time} ${method} ${path} - ${status} - ${latency}\n",
    TimeFormat: "2006-01-02 15:04:05.000000",
}))
webApp.Get("/", func(c *fiber.Ctx) error {
    time.Sleep(300 * time.Millisecond) // искусственная задержка
    return c.SendString("OK")
})
```

Пакет: `github.com/gofiber/fiber/v2/middleware/logger`.

## Лимитер запросов (theory/1)

```go
webApp.Use(limiter.New(limiter.Config{
    KeyGenerator: func(c *fiber.Ctx) string {
        return c.IP()
    },
    Max:        3,
    Expiration: 10 * time.Second,
}))
```

Не более 3 запросов с одного IP за 10 секунд, дальше `429 Too Many Requests`.

## ДЗ: requestid + logger в файл + limiter (dz)

```go
file, err := os.OpenFile(".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

webApp.Use(requestid.New())
webApp.Use(logger.New(logger.Config{
    Format:     "${locals:requestid}: ${method} ${path} - ${status}\n",
    TimeFormat: "2006-01-02 15:04:05.000000",
    Output:     file,
}))
webApp.Use(limiter.New(limiter.Config{
    KeyGenerator: func(c *fiber.Ctx) string { return c.IP() },
    Max:        1,
    Expiration: 2 * time.Second,
}))
```

`requestid` должен стоять перед `logger`, иначе `${locals:requestid}` будет пустым.

## Запуск и проверка

```bash
go run .
curl -i http://localhost:80/              # theory/0, в консоли строка лога с latency ~300ms
for i in 1 2 3 4; do curl -s -o /dev/null -w "%{http_code}\n" http://localhost:80/; done   # theory/1: 200 200 200 429
# dz (порт 8080)
curl -i http://localhost:8080/foo         # 200, заголовок X-Request-Id
curl -i http://localhost:8080/bar         # сразу после: 429
cat .log                                  # строки вида <requestid>: GET /foo - 200
```

## Подводные камни

- Порядок `Use`: logger перед limiter логирует и отклонённые запросы (429); если поставить logger после limiter, отказанные запросы в лог не попадут.
- Лимитер по `c.IP()` за прокси/балансировщиком видит IP прокси; нужен корректный `ProxyHeader`/заголовок реального IP.
- Хранилище лимитера по умолчанию in-memory, при нескольких инстансах лимиты не общие.
- `Output: file` — файл надо открыть и закрыть самому (`defer file.Close()`), логгер этого не делает.
- Файл `.log` лежит в папке задания (`t9/dz/.log`), в git его лучше не коммитить.

## Связанное

- [[MOC Go web курс]]
- [[Go start]]
- [[base зависимости и модули]]
- [[base пакеты, файлы и импорты]]
- [[http protocol]]
- [[example_CURL_requests]]
- [[example_best_practice]]
- Предыдущий урок: [[валидация запросов]]
- Следующий урок: [[jwt]]
