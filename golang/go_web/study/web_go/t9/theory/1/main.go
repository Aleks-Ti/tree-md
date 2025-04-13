package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/sirupsen/logrus"
)

func main() {
	webApp := fiber.New()

	webApp.Use(limiter.New(limiter.Config{
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		Max:        3,
		Expiration: 10 * time.Second,
	}))
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	logrus.Fatal(webApp.Listen(":80"))
}
