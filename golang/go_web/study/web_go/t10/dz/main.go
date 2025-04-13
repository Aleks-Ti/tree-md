package main

import (
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type (
	SignUpRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	SignInRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	SignInResponse struct {
		JWTToken string `json:"jwt_token"`
	}

	ProfileResponse struct {
		Email string `json:"email"`
	}

	User struct {
		Email    string
		password string
	}
)

var (
	webApiPort = ":8080"

	users = map[string]User{}

	secretKey = []byte("qwerty123456")

	contextKeyUser = "user"
)

func main() {
	webApp := fiber.New(fiber.Config{
		ReadBufferSize: 16 * 1024})
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	// BEGIN (write your solution here) (write your solution here)
	authStorage := &AuthStorage{users}
	authHandler := &AuthHandler{storage: authStorage}

	publicGroup := webApp.Group("")
	publicGroup.Post("/signup", authHandler.Register)
	publicGroup.Post("/signin", authHandler.Login)
	authorizedGroup := webApp.Group("")
	authorizedGroup.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: secretKey,
		},
		ContextKey: contextKeyUser,
	}))
	authorizedGroup.Get("/profile", authHandler.Profile)

	// END

	logrus.Fatal(webApp.Listen(webApiPort))
}

type (
	AuthHandler struct {
		storage *AuthStorage
	}

	AuthStorage struct {
		users map[string]User
	}
)

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	regReq := SignUpRequest{}
	if err := c.BodyParser(&regReq); err != nil {
		return c.SendStatus(422)
	}

	if _, exists := h.storage.users[regReq.Email]; exists {
		return c.SendStatus(fiber.StatusConflict)
	}

	h.storage.users[regReq.Email] = User{
		Email:    regReq.Email,
		password: regReq.Password,
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	regReq := SignInRequest{}
	if err := c.BodyParser(&regReq); err != nil {
		return c.SendStatus(422)
	}

	user, exists := h.storage.users[regReq.Email]
	if !exists {
		return c.SendStatus(422)
	}

	if user.password != regReq.Password {
		return c.Status(422).SendString("invalid password or email")
	}

	payload := jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString(secretKey)
	if err != nil {
		logrus.WithError(err).Error("JWT token signing")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(SignInResponse{JWTToken: t})
}

func (h *AuthHandler) Profile(c *fiber.Ctx) error {
	jwtPayload, ok := jwtPayloadFromRequest(c)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	userInfo, ok := h.storage.users[jwtPayload["sub"].(string)]
	if !ok {
		return c.SendStatus(404)
	}

	return c.JSON(ProfileResponse{
		Email: userInfo.Email,
	})
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
