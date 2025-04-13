package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type (
	// Структура HTTP-запроса на расчет диапазона дат
	DateRangeRequest struct {
		From Date `json:"from"`
		To   Date `json:"to"`
	}

	// Структура даты, которая хранит формат и значение
	Date struct {
		Value  string `json:"value"`
		Format string `json:"format"`
	}

	// Структура HTTP-ответа на расчет диапазона дат
	// Хранит значение в секундах
	DateRangeResponse struct {
		SecondsRange int64 `json:"seconds_range"`
	}
)

func main() {
	webApp := fiber.New(fiber.Config{
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})
	// Устанавливаем посредника, который будет
	// восстанавливать веб-приложение после паники
	webApp.Use(recover.New())

	webApp.Post("/daterange", func(c *fiber.Ctx) error {
		var req *DateRangeRequest
		c.BodyParser(req)

		from, _ := time.Parse(req.From.Format, req.From.Value)
		to, _ := time.Parse(req.To.Format, req.To.Value)

		return c.JSON(DateRangeResponse{
			SecondsRange: int64(to.Sub(from).Seconds()),
		})
	})

	webApp.Listen(":80")
}
