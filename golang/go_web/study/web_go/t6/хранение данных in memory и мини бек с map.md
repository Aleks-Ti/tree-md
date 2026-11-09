---
tags:
  - go
  - web
  - fiber
  - course
  - t6
up: "[[MOC Go web курс]]"
---

> [!info] Раздел: [[MOC Go web курс]]

## Идея

Простейшее хранилище — `map` в памяти процесса. Чтобы не связывать обработчики с конкретным хранилищем, выделяют интерфейс и структуру-хэндлер с зависимостью от него. Данные живут, пока жив процесс.

## Слоистая структура (theory/0)

Интерфейс хранилища, обработчик и реализация на `map`:

```go
type OrderCreatorGetter interface {
    CreateOrder(order Order) (string, error)
    GetOrder(id string) (Order, error)
}

type OrderHandler struct {
    storage OrderCreatorGetter
}

type OrderStorage struct {
    orders map[string]Order
}

func (o *OrderStorage) CreateOrder(order Order) (string, error) {
    o.orders[order.ID] = order
    return order.ID, nil
}
```

Сборка в `main` и роуты:

```go
orderHandler := &OrderHandler{
    storage: &OrderStorage{orders: make(map[string]Order)},
}
webApp.Post("/orders", orderHandler.CreateOrder)
webApp.Get("/orders/:id", orderHandler.GetOrder)
```

ID генерируется через `uuid.New().String()`; не найденный заказ — `errOrderNotFound`.

## Mutex (theory/1)

`map` небезопасна для конкурентной записи; защита — `sync.Mutex`:

```go
var (
    mu = sync.Mutex{}
    i  = 0
)

func safeInc() {
    mu.Lock()
    i++
    mu.Unlock()
}
```

## ДЗ: хранилище ссылок (dz)

```go
var links = make(map[string]string)

webApp.Post("/links", ...)          // тело {"external": "...", "internal": "..."}
webApp.Get("/links/:external", ...) // ответ {"internal": "..."}
```

Ключевая деталь: внешняя ссылка приходит в пути URL-кодированной, её надо декодировать:

```go
decodeExternalLink, err := url.QueryUnescape(externalLink)
internalLink, ok := links[decodeExternalLink]
if !ok {
    return c.Status(fiber.StatusNotFound).SendString("Link not found")
}
return c.JSON(GetLinkResponse{Internal: internalLink})
```

## Запуск и проверка

```bash
go run .
curl -X POST http://localhost:8080/links -H "Content-Type: application/json" \
  -d '{"external":"https://example.com/a","internal":"/a"}'
curl "http://localhost:8080/links/https%3A%2F%2Fexample.com%2Fa"   # {"internal":"/a"}
curl -i http://localhost:8080/links/none                          # 404

# theory/0 (порт 80)
curl -X POST http://localhost:80/orders -H "Content-Type: application/json" \
  -d '{"user_id":1,"product_ids":[10,20]}'
curl http://localhost:80/orders/<id>
```

## Подводные камни

- В theory/0 и dz map не защищена мьютексом: при конкурентных запросах возможна гонка/паника (`concurrent map writes`). Хорошая практика — `sync.RWMutex` внутри хранилища.
- Ошибки из `GetOrder` оборачиваются и возвращаются как есть, Fiber отдаст 500, а не 404.
- В dz для сломанного декодирования возвращается `502`, что семантически неверно (скорее 400).
- Данные пропадают при рестарте.
- Для работы с `%2F` в параметре пути нужно именно кодировать слэши, иначе роут не совпадёт.

## Связанное

- [[MOC Go web курс]]
- [[Go start]]
- [[base зависимости и модули]]
- [[base пакеты, файлы и импорты]]
- [[http protocol]]
- [[example_CURL_requests]]
- [[example_best_practice]]
- Предыдущий урок: [[сериализация данных]]
- Следующий урок: [[full crud fiber]]
