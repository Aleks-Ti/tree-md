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

Продолжение [[регистрация in memory]]: те же `/register` и `/login`, плюс группа роутов, закрытая JWT-middleware. Без валидного токена `/profile` недоступен. Данные пользователя достаются из claims токена (`sub` — email).

## Группы роутов и middleware

```go
publicGroup := app.Group("")
publicGroup.Post("/register", authHandler.Register)
publicGroup.Post("/login", authHandler.Login)

authorizedGroup := app.Group("")
authorizedGroup.Use(jwtware.New(jwtware.Config{
    SigningKey: jwtware.SigningKey{Key: jwtSecretKey},
    ContextKey: contextKeyUser,
}))
authorizedGroup.Get("/profile", userHandler.Profile)
```

`Use` на группе защищает только роуты этой группы. Импорты: `jwtware "github.com/gofiber/contrib/jwt"`, `jwt "github.com/golang-jwt/jwt/v5"`.

## Чтение токена в хэндлере

```go
func jwtPayloadFromRequest(c *fiber.Ctx) (jwt.MapClaims, bool) {
    jwtToken, ok := c.Context().Value(contextKeyUser).(*jwt.Token)
    if !ok { /* лог + return nil, false */ }
    payload, ok := jwtToken.Claims.(jwt.MapClaims)
    if !ok { /* лог + return nil, false */ }
    return payload, true
}

func (h *UserHandler) Profile(c *fiber.Ctx) error {
    jwtPayload, ok := jwtPayloadFromRequest(c)
    if !ok {
        return c.SendStatus(fiber.StatusUnauthorized)
    }
    userInfo, ok := h.storage.users[jwtPayload["sub"].(string)]
    // ...
    return c.JSON(ProfileResponse{Email: userInfo.Email, Name: userInfo.Name})
}
```

`AuthStorage` общий для `AuthHandler` и `UserHandler`.

## Запуск и проверка

```bash
go run .
curl -X POST http://localhost:80/register -H "Content-Type: application/json" \
  -d '{"email":"john@doe.com","name":"John","password":"secret"}'
TOKEN=$(curl -s -X POST http://localhost:80/login -H "Content-Type: application/json" \
  -d '{"email":"john@doe.com","password":"secret"}' | sed 's/.*"access_token":"\([^"]*\)".*/\1/')
curl -i http://localhost:80/profile -H "Authorization: Bearer $TOKEN"   # {"email":..,"name":..}
curl -i http://localhost:80/profile                                     # 401
```

## Подводные камни

- Заголовок должен быть строго `Authorization: Bearer <token>`.
- Ключ подписи в middleware и в `Login` должен быть одним и тем же (`jwtSecretKey`).
- Регистрировать защищённые роуты нужно в группе с middleware, иначе они останутся открытыми.
- Токен живёт 72 часа и не отзывается до истечения: «logout» на сервере без чёрного списка невозможен.
- `jwtPayload["sub"].(string)` паникует при неверном типе claim.
- Секрет и пароли — как в предыдущем шаге, только для учёбы.

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
- Предыдущий шаг: [[регистрация in memory]]
- Следующий урок: [[шаблонизатор - fiber + html]]
