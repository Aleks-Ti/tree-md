package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type User struct {
	ID      int64
	Email   string
	Age     int
	Country string
}

var users = map[int64]User{}

type (
	CreateUserRequest struct {
		// BEGIN (write your solution here)
		ID      int64  `json:"id" validate:"required,min=1"`
		Email   string `json:"email" validate:"required,email"`
		Age     int    `json:"age" validate:"required,min=18,max=130"`
		Country string `json:"country" validate:"required,allowable_country"`
		// END
	}
	// CreateUserRequest struct {
	// 	ID      int64  `json:"id" validate:"required,min=1"`
	// 	Email   string `json:"email" validate:"required,email"`
	// 	Age     int    `json:"age" validate:"required,min=18,max=130"`
	// 	Country string `json:"country" validate:"required,allowable_country"`
	// }
)

func main() {
	webApp := fiber.New(fiber.Config{ReadBufferSize: 16 * 1024})
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("1334414")
	})
	// BEGIN (write your solution here) (write your solution here)
	var authorizedCountry = []string{
		"USA",
		"Germany",
		"France",
	}
	validate := validator.New()
	validate.RegisterValidation("allowable_country", func(fl validator.FieldLevel) bool {
		text := fl.Field().String()
		for _, word := range authorizedCountry {
			if word == text {
				return true
			}
		}

		return false
	})

	webApp.Post("/users", func(ctx *fiber.Ctx) error {
		var req CreateUserRequest
		if err := ctx.BodyParser(&req); err != nil {
			return ctx.SendStatus(400)
		}

		err := validate.Struct(req)
		if err != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
		}
		user := User{
			ID:      req.ID,
			Age:     req.Age,
			Country: req.Country,
			Email:   req.Email,
		}
		users[req.ID] = user
		return ctx.SendStatus(fiber.StatusOK)
	})
	// END
	logrus.Fatal(webApp.Listen(":8080"))
}
