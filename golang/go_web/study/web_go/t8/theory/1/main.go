package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CreatePostRequest struct {
	// Описываем правила валидации в аннотациях полей структуры.
	UserID int64  `json:"user_id" validate:"required,min=0"`
	Text   string `json:"text" validate:"required,max=140,allowable_text"`
}

var forbiddenWords = []string{
	"umbrella",
	"shinra",
}

func main() {
	webApp := fiber.New()

	validate := validator.New()
	vErr := validate.RegisterValidation("allowable_text", func(fl validator.FieldLevel) bool {
		// Проверяем, что текст не содержит запрещенных слов.
		text := fl.Field().String()
		for _, word := range forbiddenWords {
			if strings.Contains(strings.ToLower(text), word) {
				return false
			}
		}

		return true
	})
	if vErr != nil {
		log.Fatal("register validation ", vErr)
	}

	webApp.Post("/posts", func(ctx *fiber.Ctx) error {
		// Парсинг JSON-строки из тела запроса в объект.
		var req CreatePostRequest
		if err := ctx.BodyParser(&req); err != nil {
			return fmt.Errorf("body parser: %w", err)
		}

		// Проверка запроса на корректность.
		err := validate.Struct(req)
		if err != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
		}

		// @TODO Сохранение поста в хранилище.

		return ctx.SendStatus(fiber.StatusOK)
	})

	logrus.Fatal(webApp.Listen(":8080"))
}
