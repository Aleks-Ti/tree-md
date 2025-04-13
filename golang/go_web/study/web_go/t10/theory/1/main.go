package main

import (
	"errors"
	"fmt"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

const (
	contextKeyUser = "user"
)

func main() {
	app := fiber.New()

	authStorage := &AuthStorage{map[string]User{}}
	authHandler := &AuthHandler{storage: authStorage}
	userHandler := &UserHandler{storage: authStorage}

	// Группа обработчиков, которые доступны неавторизованным пользователям
	publicGroup := app.Group("")
	publicGroup.Post("/register", authHandler.Register)
	publicGroup.Post("/login", authHandler.Login)

	// Группа обработчиков, которые требуют авторизации
	authorizedGroup := app.Group("")
	authorizedGroup.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: jwtSecretKey,
		},
		ContextKey: contextKeyUser,
	}))
	authorizedGroup.Get("/profile", userHandler.Profile)

	logrus.Fatal(app.Listen(":80"))
}

type (
	// Обработчик HTTP-запросов на регистрацию и аутентификацию пользователей
	AuthHandler struct {
		storage *AuthStorage
	}

	// Хранилище зарегистрированных пользователей
	// Данные хранятся в оперативной памяти
	AuthStorage struct {
		users map[string]User
	}

	// Структура данных с информацией о пользователе
	User struct {
		Email    string
		Name     string
		password string
	}
)

// Структура HTTP-запроса на регистрацию пользователя
type RegisterRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// Обработчик HTTP-запросов на регистрацию пользователя
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	regReq := RegisterRequest{}
	if err := c.BodyParser(&regReq); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}

	// Проверяем, что пользователь с таким email еще не зарегистрирован
	if _, exists := h.storage.users[regReq.Email]; exists {
		return errors.New("the user already exists")
	}

	// Сохраняем в память нового зарегистрированного пользователя
	h.storage.users[regReq.Email] = User{
		Email:    regReq.Email,
		Name:     regReq.Name,
		password: regReq.Password,
	}

	return c.SendStatus(fiber.StatusCreated)
}

// Структура HTTP-запроса на вход в аккаунт
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Структура HTTP-ответа на вход в аккаунт
// В ответе содержится JWT-токен авторизованного пользователя
type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

var (
	errBadCredentials = errors.New("email or password is incorrect")
)

// Секретный ключ для подписи JWT-токена
// Необходимо хранить в безопасном месте
var jwtSecretKey = []byte("very-secret-key")

// Обработчик HTTP-запросов на вход в аккаунт
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	regReq := LoginRequest{}
	if err := c.BodyParser(&regReq); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}

	// Ищем пользователя в памяти приложения по электронной почте
	user, exists := h.storage.users[regReq.Email]
	// Если пользователь не найден, возвращаем ошибку
	if !exists {
		return errBadCredentials
	}
	// Если пользователь найден, но у него другой пароль, возвращаем ошибку
	if user.password != regReq.Password {
		return errBadCredentials
	}

	// Генерируем JWT-токен для пользователя,
	// который он будет использовать в будущих HTTP-запросах

	// Генерируем полезные данные, которые будут храниться в токене
	payload := jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	// Создаем новый JWT-токен и подписываем его по алгоритму HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString(jwtSecretKey)
	if err != nil {
		logrus.WithError(err).Error("JWT token signing")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(LoginResponse{AccessToken: t})
}

// Обработчик HTTP-запросов, которые связаны с пользователем
type UserHandler struct {
	storage *AuthStorage
}

// Структура HTTP-ответа с информацией о пользователе
type ProfileResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func jwtPayloadFromRequest(c *fiber.Ctx) (jwt.MapClaims, bool) {
	jwtToken, ok := c.Context().Value(contextKeyUser).(*jwt.Token)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"jwt_token_context_value": c.Context().Value(contextKeyUser),
		}).Error("wrong type of JWT token in context")
		return nil, false
	}

	payload, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"jwt_token_claims": jwtToken.Claims,
		}).Error("wrong type of JWT token claims")
		return nil, false
	}

	return payload, true
}

// Обработчик HTTP-запросов на получение информации о пользователе
func (h *UserHandler) Profile(c *fiber.Ctx) error {
	jwtPayload, ok := jwtPayloadFromRequest(c)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	userInfo, ok := h.storage.users[jwtPayload["sub"].(string)]
	if !ok {
		return errors.New("user not found")
	}

	return c.JSON(ProfileResponse{
		Email: userInfo.Email,
		Name:  userInfo.Name,
	})
}
