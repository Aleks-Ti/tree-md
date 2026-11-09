---
tags:
  - go
  - web
  - course
  - t1
up: "[[MOC Go web курс]]"
---

> [!info] Раздел: [[MOC Go web курс]]

# t1: HTTP-сервер на `net/http`

Базовый сервер без фреймворка. Дальше курс переходит на Fiber (см. [[Fiber welcome + query params и роутинг]]).

## Theory: `HandleFunc` + query-параметр

```go
http.HandleFunc("/courses", func(w http.ResponseWriter, r *http.Request) {
    page := r.URL.Query().Get("page")   // "" если параметра нет
    switch page {
    case "":  w.Write([]byte("Введите номер курса!"))
    case "1": w.Write([]byte("..."))
    default:  w.Write([]byte("Курс в разработке..."))
    }
})
http.ListenAndServe(":80", nil)
```

- Тело ответа — `[]byte`; статус по умолчанию `200`.
- `nil` вторым аргументом = глобальный `DefaultServeMux`.

```bash
curl "localhost:80/courses?page=1"
```

## DZ: свой `ServeMux` и `http.Server`

```go
mux := http.NewServeMux()
mux.HandleFunc("/courses/description", CourseDescHandler)

server := &http.Server{
    Addr:              ":8080",
    Handler:           mux,
    ReadHeaderTimeout: 10 * time.Second,
}
err := server.ListenAndServe()
```

Хендлер парсит `course_id` и отвечает кодами:

| Ситуация | Ответ |
| --- | --- |
| `course_id` не число | `400 Bad Request` (`strconv.ParseInt` вернул ошибку) |
| курса нет в `map` | `404 Not Found` |
| найден | `200` + описание |

```bash
curl -i "localhost:8080/courses/description?course_id=2"
```

## Заметки

> [!tip] Почему свой `http.Server`
> У `http.ListenAndServe` нет таймаутов; `ReadHeaderTimeout` защищает от медленных клиентов (slowloris). Тот же мотив — таймауты Fiber в [[recover, таймауты и очередь push]].

- Глобальный `DefaultServeMux` лучше не использовать в проде: любой импортированный пакет может на нём регистрировать хендлеры.
- Порт `:80` требует прав root на Linux — см. [[Порты, и действия с ними]].

---

## Связанное

- [[MOC Go web курс]]
- [[logrus и логирование]]
- [[Go start]]
- [[http protocol]]
- [[example_CURL_requests]]
