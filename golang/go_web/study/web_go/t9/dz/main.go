package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/sirupsen/logrus"
)

func main() {

	file, err := os.OpenFile(".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	webApp := fiber.New(fiber.Config{
		ReadBufferSize: 16 * 1024})
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	// BEGIN (write your solution here) (write your solution here)
	webApp.Use(requestid.New())
	webApp.Use(logger.New(logger.Config{
		Format:     "${locals:requestid}: ${method} ${path} - ${status}\n",
		TimeFormat: "2006-01-02 15:04:05.000000",
		Output:     file,
	}))
	webApp.Use(limiter.New(limiter.Config{
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		Max:        1,
		Expiration: 2 * time.Second,
	}))
	// END

	webApp.Get("/foo", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	webApp.Get("/bar", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	logrus.Fatal(webApp.Listen(":8080"))
}
