---
tags:
  - go
  - web
  - fiber
  - course
  - t11
up: "[[MOC Go web курс]]"
---

> [!info] Раздел: [[MOC Go web курс]]

## Идея

Fiber может отдавать не JSON, а HTML, собранный из шаблона. Подключается движок `github.com/gofiber/template/html` (обёртка над стандартным `html/template`): указываем папку и расширение, затем в хэндлере вызываем `c.Render(имя_шаблона, данные)`.

## Подключение и рендер (theory/0)

```go
viewsEngine := html.New("./template", ".tmpl")

webApp := fiber.New(fiber.Config{
    Views: viewsEngine,
})

webApp.Get("/profile", func(c *fiber.Ctx) error {
    return c.Render("profile", fiber.Map{
        "name":  "John",
        "email": "john@doe.com",
    })
})
```

Имя шаблона — имя файла без расширения (`template/profile.tmpl`).

## Цикл и условие в шаблоне (theory/1)

Данные — срез структур:

```go
type Film struct {
    Title    string
    IsViewed bool
}
webApp.Get("/films", func(c *fiber.Ctx) error {
    return c.Render("film-list", films)
})
```

Шаблон `template/film-list.tmpl`:

```html
{{ range $film := . }}
<div>
    <h2>{{ $film.Title }}</h2>
    <p>
        {{ if $film.IsViewed }}Просмотрено{{ else }}Не просмотрено{{ end }}
    </p>
</div>
{{ end }}
```

(в оригинале внутри ещё эмодзи-отметки.) Точка `.` — переданные данные.

## ДЗ: список товаров (dz)

Папка шаблонов `./templates`, файл `items.tmpl`:

```go
viewsEngine := html.New("./templates", ".tmpl")
webApp.Post("/items", itemHandler.AddItem)
webApp.Get("/items/view", itemHandler.GetItems)
```

```go
func (h *ItemsHandler) GetItems(c *fiber.Ctx) error {
    err := c.Render("items", h.storage.item)
    ...
}
```

В шаблоне — `{{ range $item := . }}` с `{{ $item.Name }}` и `{{ $item.Price }}`.

## Запуск и проверка

Запускать из папки урока, чтобы относительный путь к шаблонам нашёлся.

```bash
go run .
curl http://localhost:80/profile            # theory/0
curl http://localhost:80/films              # theory/1
# dz (порт 8080)
curl -X POST http://localhost:8080/items -H "Content-Type: application/json" \
  -d '{"name":"Book","price":10}'
curl http://localhost:8080/items/view       # или открыть в браузере
```

## Подводные камни

- `html.New("./template", ...)` — относительный путь от рабочей директории процесса, а не от исходника. Запуск из другой папки даёт ошибку «template not found».
- Шаблон `profile.tmpl` в theory/0 содержит статичный текст (John, john@doe.com), переданные в `fiber.Map` данные в нём не подставляются: нужны `{{ .name }}` и `{{ .email }}`.
- В dz после `c.Render` в `GetItems` вызывается ещё `c.SendStatus(200)`, что затирает статус и лишний раз пишет ответ; стоит возвращать результат `Render` напрямую. Также в нём остаётся отладочный `println`.
- Разные версии Fiber/template в `go.mod` (в dz fiber v2.44.0, template v1.8.1): `go mod tidy` перед запуском.
- `html/template` экранирует вывод (защита от XSS), не нужно подменять на `text/template`.

## Связанное

- [[MOC Go web курс]]
- [[Go start]]
- [[base зависимости и модули]]
- [[base пакеты, файлы и импорты]]
- [[http protocol]]
- [[example_CURL_requests]]
- [[example_best_practice]]
- Предыдущий урок: [[jwt]]
