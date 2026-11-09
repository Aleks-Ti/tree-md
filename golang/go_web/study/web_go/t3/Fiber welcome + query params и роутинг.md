---
tags:
  - go
  - web
  - fiber
  - course
  - t3
up: "[[MOC Go web курс]]"
---

> [!info] Раздел: [[MOC Go web курс]]

## Идея

Fiber (`github.com/gofiber/fiber/v2`) — веб-фреймворк. Создаём приложение `fiber.New()`, регистрируем обработчики по методу и пути (`webApp.Get(path, handler)`), запускаем `Listen`. Обработчик принимает `*fiber.Ctx` и возвращает `error`. Query-параметры (`?a=1`) читаются через `c.Query(name, default)`.

## Hello-world (theory/0)

```go
webApp := fiber.New()
webApp.Get("/address", func(c *fiber.Ctx) error {
    return c.SendString("145 DUNDEE SOUTH SAN FRANCISCO CA 94080-1023. USA")
})
logrus.Fatal(webApp.Listen(":80"))
```

`Listen` блокирует поток и возвращает ошибку только при падении, поэтому её оборачивают в `logrus.Fatal`.

## Query params (theory/1)

```go
const profileUnknown = "unknown"

webApp.Get("/profiles", func(c *fiber.Ctx) error {
    profileID := c.Query("profile_id", profileUnknown)
    if profileID == "" {
        profileID = profileUnknown
    }
    if profileID == profileUnknown {
        return c.Status(fiber.StatusUnprocessableEntity).SendString("profile_id is required")
    }
    return c.SendString(fmt.Sprintf("User Profile ID: %s", profileID))
})
```

Второй аргумент `c.Query` — значение по умолчанию, если параметра нет.

## ДЗ: конвертер валют (dz)

Таблица курсов в `map[string]float64` с ключами вида `"USD/EUR"`; `GET /convert?from=..&to=..`:

```go
from := c.Query("from", "null")
if from == "null" {
    return c.Status(fiber.StatusUnprocessableEntity).SendString("`from` mandatory parameter for the request")
}
// то же для to
pair := from + "/" + to
if value, ok := exchangeRate[pair]; ok {
    return c.Status(fiber.StatusOK).SendString(fmt.Sprintf("%.2f", value))
} else {
    return c.Status(fiber.StatusNotFound).SendString("404 Not Found")
}
```

## Запуск и проверка

```bash
go run .
curl "http://localhost:80/address"                 # theory/0
curl "http://localhost:80/profiles?profile_id=42"  # theory/1 -> User Profile ID: 42
curl -i "http://localhost:80/profiles"             # 422
curl "http://localhost:8080/convert?from=USD&to=EUR"   # dz -> 0.80
curl -i "http://localhost:8080/convert?from=USD"       # 422
curl -i "http://localhost:8080/convert?from=USD&to=RUB" # 404
```

## Подводные камни

- Порт `:80` в theory требует прав/свободного порта; dz слушает `:8080`.
- В dz формат `%.2f` обрезает курс: `JPY/USD` (0.0091) вернётся как `0.01`.
- В dz «нет параметра» определяется sentinel-значением `"null"`, поэтому буквальный `?from=null` тоже считается отсутствующим.
- Нужен `go mod tidy` в папке урока (у каждого урока свой `go.mod`).

## Связанное

- [[MOC Go web курс]]
- [[Go start]]
- [[base зависимости и модули]]
- [[base пакеты, файлы и импорты]]
- [[http protocol]]
- [[example_CURL_requests]]
- Следующий урок: [[руотинг + динамический роутинг]]
