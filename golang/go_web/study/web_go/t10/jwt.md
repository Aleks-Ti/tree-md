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

JWT (JSON Web Token) — подписанный токен вида `header.payload.signature`. Сервер при логине выдаёт токен, клиент присылает его в заголовке `Authorization: Bearer <token>`, middleware проверяет подпись и срок действия. Состояние сессии на сервере не хранится. Подпись HS256 — HMAC с общим секретом (см. [[Криптография]]).

Подробнее по шагам: [[регистрация in memory]] (регистрация и выдача токена) и [[registerr + login + permission middleware fiber]] (защищённые роуты).

## ДЗ: signup / signin / profile (t10/dz)

Роуты разделены на две группы: публичную и требующую токена.

```go
publicGroup := webApp.Group("")
publicGroup.Post("/signup", authHandler.Register)
publicGroup.Post("/signin", authHandler.Login)

authorizedGroup := webApp.Group("")
authorizedGroup.Use(jwtware.New(jwtware.Config{
    SigningKey: jwtware.SigningKey{Key: secretKey},
    ContextKey: contextKeyUser,
}))
authorizedGroup.Get("/profile", authHandler.Profile)
```

Выдача токена при входе:

```go
payload := jwt.MapClaims{
    "sub": user.Email,
    "exp": time.Now().Add(time.Hour * 72).Unix(),
}
token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
t, err := token.SignedString(secretKey)
return c.JSON(SignInResponse{JWTToken: t})
```

Коды ответов: `/signup` — 200, 409 если email занят, 422 при плохом теле; `/signin` — 422 при неверных email/пароле; `/profile` — 401 без токена (отдаёт middleware).

Импорты: `jwtware "github.com/gofiber/contrib/jwt"` и `jwt "github.com/golang-jwt/jwt/v5"`.

## Запуск и проверка

```bash
go run .
curl -X POST http://localhost:8080/signup -H "Content-Type: application/json" \
  -d '{"email":"a@b.com","password":"pass"}'
curl -X POST http://localhost:8080/signin -H "Content-Type: application/json" \
  -d '{"email":"a@b.com","password":"pass"}'      # {"jwt_token":"..."}
curl -i http://localhost:8080/profile -H "Authorization: Bearer <jwt_token>"   # {"email":"a@b.com"}
curl -i http://localhost:8080/profile                                          # 401
```

## Подводные камни

- Секрет в коде (`"qwerty123456"`) допустим только для учёбы; в реальности читать из окружения/секрет-хранилища.
- Пароли хранятся открытым текстом в памяти; в проде нужен хэш (bcrypt/argon2).
- Payload токена только закодирован (base64), не зашифрован: чувствительные данные туда не класть.
- В dz и theory/1 используется `golang-jwt/jwt/v5`, а в theory/0 старая `github.com/golang-jwt/jwt` (v3/v4): не смешивать, иначе `*jwt.Token` не совпадёт по типу.
- `ContextKey` в конфиге и ключ в `c.Context().Value(...)` должны совпадать, иначе токен из контекста не достать.
- Типовое приведение `jwtPayload["sub"].(string)` паникует, если claim не строка.

## Связанное

- [[MOC Go web курс]]
- [[Go start]]
- [[base зависимости и модули]]
- [[base пакеты, файлы и импорты]]
- [[http protocol]]
- [[example_CURL_requests]]
- [[Криптография]]
- [[example_best_practice]]
- Предыдущий урок: [[middleware]]
- Следующий урок: [[шаблонизатор - fiber + html]]
- Шаги: [[регистрация in memory]], [[registerr + login + permission middleware fiber]]
