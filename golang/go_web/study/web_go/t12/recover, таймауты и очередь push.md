---
tags:
  - go
  - web
  - fiber
  - course
  - t12
up: "[[MOC Go web курс]]"
---

> [!info] Раздел: [[MOC Go web курс]]

# t12: `recover` и таймауты в Fiber

## Theory: `/daterange`

```go
webApp := fiber.New(fiber.Config{
    ReadTimeout:  3 * time.Second,   // theory/1: добавлены таймауты
    WriteTimeout: 3 * time.Second,
})
webApp.Use(recover.New())            // перехватывает panic -> ответ 500, процесс жив
```

| Версия | Отличие |
| --- | --- |
| `theory/0` | только `recover.New()` |
| `theory/1` | + `ReadTimeout` / `WriteTimeout` |

> [!bug] Намеренная (?) ошибка в примере
> ```go
> var req *DateRangeRequest   // nil-указатель
> c.BodyParser(req)           // ошибка не проверена
> from, _ := time.Parse(req.From.Format, ...)   // nil dereference -> panic
> ```
> Это и показывает `recover`: без него сервер упал бы. В реальном коде: `var req DateRangeRequest` + `BodyParser(&req)` + проверка `err` (см. [[сериализация данных]]).

Формат дат — layout Go (`"2006-01-02"`), см. `time.Parse`.

## DZ: очередь push-уведомлений

```go
webApp.Post("/push/send", func(c *fiber.Ctx) error {
    var req SendPushNotificationRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).SendString("Invalid JSON")
    }
    pushNotificationsQueue = append(pushNotificationsQueue, PushNotification{...})
    if len(pushNotificationsQueue) > 3 {
        panic("Queue is full")      // recover превратит в 500
    }
    return c.SendStatus(fiber.StatusOK)
})
```

```bash
curl -X POST localhost:8080/push/send -H 'Content-Type: application/json' \
     -d '{"message":"hi","user_id":1}'      # 200, с 4-го раза — 500
```

## Заметки

- `panic` для «очередь полна» — для демонстрации `recover`; в реальном API лучше вернуть `503`/`429`.
- Глобальный срез без мьютекса — гонки при параллельных запросах (см. [[хранение данных in memory и мини бек с map]]).
- Очередь никогда не очищается: после 4-го запроса каждый следующий тоже падает.
- `logrus.Fatal(webApp.Listen(...))` — логирование как в [[logrus и логирование]].

---

## Связанное

- [[MOC Go web курс]]
- [[middleware]]
- [[net-http сервер на стандартной библиотеке]]
- [[Go start]]
- [[example_CURL_requests]]
