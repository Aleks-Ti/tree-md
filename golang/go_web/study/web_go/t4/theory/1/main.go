package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var counter int64 = 0

func main() {
	webApp := fiber.New()

	webApp.Get("/counter", func(c *fiber.Ctx) error {
		return c.SendString(strconv.FormatInt(counter, 10))
	})

	webApp.Post("/counter", func(c *fiber.Ctx) error {
		counter++

		return c.SendStatus(fiber.StatusOK)
	})

	logrus.Fatal(webApp.Listen(":80"))
}
