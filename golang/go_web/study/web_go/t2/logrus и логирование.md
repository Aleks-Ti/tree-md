---
tags:
  - go
  - web
  - course
  - logging
  - t2
up: "[[MOC Go web курс]]"
---

> [!info] Раздел: [[MOC Go web курс]]

# t2: логирование через `logrus`

```bash
go get github.com/sirupsen/logrus
```

## Theory

```go
logrus.WithFields(logrus.Fields{"port": port}).Info("Starting a web-server on port")

_, err := w.Write([]byte("Welcome to Hexlet"))
if err != nil {
    logrus.WithError(err).Error("write hello hexlet")
}
logrus.Fatal(http.ListenAndServe(":"+port, nil))
```

| Метод | Назначение |
| --- | --- |
| `WithFields(Fields{...})` | структурные поля (ключ-значение) |
| `WithError(err)` | добавить ошибку в запись |
| `Info / Warn / Error` | уровни |
| `Fatal` | записать и `os.Exit(1)` |
| `SetFormatter(&TextFormatter{FullTimestamp: true})` | формат (в коде закомментирован) |

`ListenAndServe` блокирует и возвращает ошибку только при падении, поэтому оборачивается в `Fatal`.

## DZ: `/sum` с логом в файл

- Свой логгер: `logger := logrus.New()`; вывод в файл `.log` в текущей папке:

```go
file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
logger.SetOutput(file)
```

- `/sum?x=1&y=2`: параметры парсятся `strconv.Atoi`, при ошибке — `400`.
- Переполнение `int` проверяется **до** сложения:

```go
if (y > 0 && x > math.MaxInt-y) || (y < 0 && x < math.MinInt-y) {
    w.Write([]byte("-1"))
    logger.WithFields(logrus.Fields{"x": x, "y": y}).Warn("Sum overflows int")
    return
}
```

```bash
curl "localhost:8080/sum?x=5&y=7"      # 12
curl "localhost:8080/sum?x=9223372036854775807&y=1"   # -1 + запись в .log
```

## Заметки

> [!warning] Файл `.log` попал в репозиторий
> В `t2/dz/.log` и `t9/dz/.log` лежат логи. Добавь `*.log` / `.log` в `.gitignore`.

- Конечные `logrus.WithFields(...).Fatal(server.ListenAndServe())` в dz пишет в **стандартный** логгер (stderr), а не в файловый `logger` — это два разных логгера.
- `0755` для лог-файла избыточен: достаточно `0644` ([[chmod]]).

---

## Связанное

- [[MOC Go web курс]]
- [[net-http сервер на стандартной библиотеке]]
- [[middleware]] — логирование запросов middleware
- [[Go start]]
- [[Nginx logs]]
