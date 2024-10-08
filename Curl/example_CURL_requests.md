# Curl

Офф дока ```https://curl.se```

## Навигация

- [Примеры запросов](#пример-запросов)
- - [POST запрос с Body](#типичный-post-запрос-json-с-body)
- - [GET запрос с query](#get-запрос-с-query-параметрами)
- - [GET запрос с path](#get-запрос-с-использованием-path-параметров)
- - [POST запрос с FormData](#post-запрос-с-использованием-formdata)
- [Лайфхаки](#лайфхаки)
- - [Чтение и сохранение ответов](#чтение-и-сохранение-ответов)
- - [Сохранение сессии (Cookies)](#сохранение-сессии-cookies)
- - [Просмотр детализированных логов запроса](#просмотр-детализированных-логов-запроса)
- - [Проверка времени выполнения запроса](#проверка-времени-выполнения-запроса)
- - [Работа с прокси](#работа-с-прокси)
- - [Загрузка файлов с выводом прогресса](#загрузка-файлов-с-выводом-прогресса)
- - [Повторение запросов](#повторение-запросов)
- - [Автоматическая обработка перенаправлений](#автоматическая-обработка-перенаправлений)
- - [Параллельные запросы с использованием xargs](#параллельные-запросы-с-использованием-xargs)

## Пример запросов

Простые и типичные запросы для API

### Типичный POST запрос json с Body

```rust
curl -X 'POST' \
  'http://127.0.0.1:8000/hello' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "Ivan",
  "surname": "Petrov",
  "age": 18,
  "is_staff": false,
  "education_level": "Среднее образование"
}'
```

### GET запрос с query параметрами

```rust
curl -X 'GET' \
  'http://127.0.0.1:8000/users?name=Ivan&age=18&is_staff=false' \
  -H 'accept: application/json'
```

### GET запрос с использованием path параметров

```rust
curl -X 'GET' \
  'http://127.0.0.1:8000/users/123/profile' \
  -H 'accept: application/json'
```

### POST запрос с использованием FormData

- Поле category со списком id (например, [1, 2, 3]).
- Поле avatar (одна картинка).
- Поле images (несколько картинок).

```rust
curl -X 'POST' \
  'http://127.0.0.1:8000/upload' \
  -H 'accept: application/json' \
  -H 'Content-Type: multipart/form-data' \
  -F 'name=Ivan' \
  -F 'surname=Petrov' \
  -F 'category[]=1' \
  -F 'category[]=2' \
  -F 'category[]=3' \
  -F 'avatar=@/path/to/avatar.jpg' \
  -F 'images[]=@/path/to/image1.jpg' \
  -F 'images[]=@/path/to/image2.jpg'
```

## Лайфхаки

### Чтение и сохранение ответов

Запись ответа в файл:

Cохранить ответ запроса в файл, можно использовать флаг `-o`:

```rust
curl -X 'GET' \
  'http://127.0.0.1:8000/users/123/profile' \
  -o response.json
```

Этот запрос сохранит ответ в файл response.json, вместо вывода в терминал.

Просмотр только заголовков ответа:
Если вам нужно увидеть только заголовки ответа (например, чтобы проверить статус кода или cookies), используйте флаг -I:

```rust
curl -I 'http://127.0.0.1:8000/users/123/profile'
```

### Сохранение сессии (Cookies)

Для взаимодействия с API, которые требуют авторизации и сессий, вы можете сохранить и использовать cookies:

- Сохранение cookies в файл:

```rust
curl -c cookies.txt -X 'POST' \
  'http://127.0.0.1:8000/login' \
  -H 'Content-Type: application/json' \
  -d '{
  "username": "user",
  "password": "password"
}'
```

Флаг -c cookies.txt сохраняет cookies в файл cookies.txt.

- Использование сохранённых cookies:

```rust
curl -b cookies.txt -X 'GET' \
  'http://127.0.0.1:8000/protected-resource'
```

Флаг -b cookies.txt будет отправлять cookies с запросом, сохраняя сессию.

### Просмотр детализированных логов запроса

Чтобы лучше понимать, что происходит во время запроса, можно включить детализированные логи с помощью флага `-v` (verbose). Это поможет отслеживать шаги запроса, видеть все заголовки, устанавливаемые `curl`, и ответы сервера.

```rust
curl -v 'http://127.0.0.1:8000/users/123/profile'
```

Если нужен ещё более детализированный лог, можно использовать флаг `--trace`:

```rust
curl --trace trace.log 'http://127.0.0.1:8000/users/123/profile'
```

Этот флаг записывает полный трейсинг в файл trace.log, включая DNS-запросы, SSL-сессии и т.д.

### Авторизация

Если нужно передать токен для авторизации, можно использовать заголовок `Authorization`:

- JWT токен:

```rust
curl -X 'GET' \
  'http://127.0.0.1:8000/protected-resource' \
  -H 'Authorization: Bearer YOUR_TOKEN'
```

- Базовая авторизация:
- - Для передачи имени пользователя и пароля можно использовать флаг -u:

```rust
curl -u username:password \
  'http://127.0.0.1:8000/protected-resource'
```

### Проверка времени выполнения запроса

Чтобы понять, сколько времени занимает выполнение запроса, можно использовать флаг `-w` (`write-out`). Например, выводим время ответа:

```rust
curl -w "Time: %{time_total}\n" -o /dev/null -s \
  'http://127.0.0.1:8000/users/123/profile'
```

- `-o /dev/null` указывает записывать тело ответа в пустоту,
- `-s` подавляет прогресс бар,
- `%{time_total}` выводит общее время выполнения запроса.

### Работа с прокси

Если вам нужно отправить запрос через прокси-сервер, используйте флаг -x:

```rust
curl -x http://proxy-server:port \
  'http://127.0.0.1:8000/users/123/profile'
```

### Загрузка файлов с выводом прогресса

Если вы загружаете большие файлы и хотите видеть прогресс, используйте команду -O для сохранения файлов с их оригинальными именами и отображения прогресса:

```rust
curl -O http://example.com/largefile.zip
```

### Повторение запросов

Если сервер временно не отвечает, можно настроить повторение запроса:

```rust
curl --retry 5 'http://127.0.0.1:8000/resource'
```

Этот пример заставит curl повторить запрос до 5 раз при сбое.

### Автоматическая обработка перенаправлений

Если сервер возвращает `HTTP редиректы`, можно автоматически следовать за ними с помощью флага `-L`:

```rust
curl -L 'http://example.com/old-url'
```

### Параллельные запросы с использованием xargs

Для отправки нескольких параллельных запросов можно использовать xargs:

```rust
echo 'http://example.com/resource1' 'http://example.com/resource2' | xargs -n 1 -P 2 curl -O
```

Этот запрос запустит два параллельных запроса к разным ресурсам.
