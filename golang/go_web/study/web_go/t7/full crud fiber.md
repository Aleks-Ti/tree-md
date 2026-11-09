---
tags:
  - go
  - web
  - fiber
  - course
  - t7
up: "[[MOC Go web курс]]"
---

> [!info] Раздел: [[MOC Go web курс]]

## Идея

CRUD — четыре операции над ресурсом, которые мапятся на HTTP-методы: создать (`POST`), прочитать (`GET`), частично обновить (`PATCH`), удалить (`DELETE`). Хранилище — `map` в памяти, ресурс адресуется по `/:id`.

## Структуры (theory/0, сотрудники)

Отдельные типы под каждый запрос/ответ и модель хранилища:

```go
type CreateEmployeeRequest struct {
    Email string `json:"email"`
    Role  string `json:"role"`
}
type EmployeePayload struct {
    ID    string `json:"id"`
    Email string `json:"email"`
    Role  string `json:"role"`
}
type GetEmployeeResponse struct {
    EmployeePayload
}
```

Хранилище с методами `Create/List/Get/Update/Delete`; `Update` меняет только непустые поля:

```go
if email != "" {
    empl.Email = email
}
if role != "" {
    empl.Role = role
}
```

## Роуты (theory/0)

```go
webApp.Post("/employees", ...)        // создать, ответ {"id": "<uuid>"}
webApp.Get("/employees", ...)         // список {"employees": [...]}
webApp.Get("/employees/:id", ...)     // fiber.ErrNotFound, если нет
webApp.Patch("/employees/:id", ...)   // частичное обновление
webApp.Delete("/employees/:id", func(c *fiber.Ctx) error {
    storage.Delete(c.Params("id"))
    return c.SendStatus(fiber.StatusNoContent)
})
```

Приведение типов между структурами с одинаковыми полями: `EmployeePayload(empl)`.

## ДЗ: задачи (dz)

Целочисленный ID через счётчик, `GET /tasks/:id`, `POST /tasks`, `PATCH /tasks/:id`, `DELETE /tasks/:id`:

```go
var (
    taskIDCounter int64 = 1
    tasks               = make(map[int64]Task)
)

taskId, err := strconv.ParseInt(c.Params("id"), 10, 64)
if err != nil {
    return c.Status(fiber.StatusBadRequest).SendString("Error params")
}
task, ok := tasks[taskId]
if !ok {
    return c.Status(fiber.StatusNotFound).SendString("Not Found")
}
```

## Запуск и проверка

```bash
go run .
# dz (порт 8080)
curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json" \
  -d '{"description":"buy milk","deadline":1700000000}'     # {"id":1}
curl http://localhost:8080/tasks/1
curl -X PATCH http://localhost:8080/tasks/1 -H "Content-Type: application/json" \
  -d '{"description":"buy bread"}'
curl -i -X DELETE http://localhost:8080/tasks/1
curl -i http://localhost:8080/tasks/1                       # 404
# theory/0 (порт 80)
curl -X POST http://localhost:80/employees -H "Content-Type: application/json" \
  -d '{"email":"a@b.c","role":"dev"}'
```

## Подводные камни

- PATCH с нулевыми значениями: через «пустое = не менять» нельзя сбросить поле в `""`/`0`. Для различения нужны указатели в структуре запроса.
- В theory ошибка `Update` (не найден) возвращается через `fmt.Errorf`, т.е. 500 вместо 404.
- `Delete` идемпотентен: удаление несуществующего в theory всё равно даёт 204, в dz — 404.
- Обход `map` не гарантирует порядок, список сотрудников приходит в случайном порядке.
- Map без мьютекса, счётчик ID без атомарности: для реального сервиса нужна синхронизация.

## Связанное

- [[MOC Go web курс]]
- [[Go start]]
- [[base зависимости и модули]]
- [[base пакеты, файлы и импорты]]
- [[http protocol]]
- [[example_CURL_requests]]
- [[example_best_practice]]
- Предыдущий урок: [[хранение данных in memory и мини бек с map]]
- Следующий урок: [[валидация запросов]]
