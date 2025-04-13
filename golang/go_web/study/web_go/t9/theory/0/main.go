package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	webApp := fiber.New()

	webApp.Use(logger.New(logger.Config{
		Format:     "${time} ${method} ${path} - ${status} - ${latency}\n",
		TimeFormat: "2006-01-02 15:04:05.000000",
	}))
	webApp.Get("/", func(c *fiber.Ctx) error {
		// Создаем искусственную задержку, чтобы проверить логирование.
		time.Sleep(300 * time.Millisecond)

		return c.SendString("OK")
	})

	logrus.Fatal(webApp.Listen(":80"))
}
