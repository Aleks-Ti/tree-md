---
tags:
  - go
  - web
  - fiber
  - course
  - t10
up: "[[MOC Go web курс]]"
---

> [!info] Раздел: [[MOC Go web курс]]

## Идея

Первый шаг аутентификации: регистрация пользователя в памяти (`map` по email) и логин с выдачей JWT. Здесь токен только выдаётся; проверки токена на роутах ещё нет (она в [[registerr + login + permission middleware fiber]]).

## Структуры

```go
type (
    AuthHandler struct { storage *AuthStorage }
    AuthStorage struct { users map[string]User }
    User struct {
        Email    string
        Name     string
        password string
    }
)

type RegisterRequest struct {
    Email    string `json:"email"`
    Name     string `json:"name"`
    Password string `json:"password"`
}
```

## Роуты и регистрация

```go
app.Post("/register", authHandler.Register)
app.Post("/login", authHandler.Login)
```

```go
if _, exists := h.storage.users[regReq.Email]; exists {
    return errors.New("the user already exists")
}
h.storage.users[regReq.Email] = User{
    Email: regReq.Email, Name: regReq.Name, password: regReq.Password,
}
return c.SendStatus(fiber.StatusCreated)
```

## Логин и JWT

```go
var jwtSecretKey = []byte("very-secret-key")

payload := jwt.MapClaims{
    "sub": user.Email,
    "exp": time.Now().Add(time.Hour * 72).Unix(),
}
token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
t, err := token.SignedString(jwtSecretKey)
return c.JSON(LoginResponse{AccessToken: t})
```

Ответ: `{"access_token": "..."}`.

## Запуск и проверка

```bash
go run .
curl -i -X POST http://localhost:80/register -H "Content-Type: application/json" \
  -d '{"email":"john@doe.com","name":"John","password":"secret"}'   # 201
curl -X POST http://localhost:80/login -H "Content-Type: application/json" \
  -d '{"email":"john@doe.com","password":"secret"}'                 # {"access_token":"..."}
```

Содержимое токена можно посмотреть на jwt.io (payload не шифруется).

## Подводные камни

- Ошибки (`errBadCredentials`, «user already exists») возвращаются как `error`, Fiber отдаёт 500 вместо 409/401.
- Одно и то же сообщение для «нет пользователя» и «неверный пароль» — правильно: не раскрывает, существует ли email.
- Пароль в открытом виде, секрет в коде: только для учёбы.
- Здесь импорт `github.com/golang-jwt/jwt` (без `/v5`), в следующем шаге `jwt/v5`.
- Map без мьютекса: параллельные регистрации небезопасны.

## Связанное

- [[MOC Go web курс]]
- [[Go start]]
- [[base зависимости и модули]]
- [[base пакеты, файлы и импорты]]
- [[http protocol]]
- [[example_CURL_requests]]
- [[Криптография]]
- [[example_best_practice]]
- Обзор: [[jwt]]
- Следующий шаг: [[registerr + login + permission middleware fiber]]
