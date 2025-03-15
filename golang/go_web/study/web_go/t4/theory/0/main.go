package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	webApp := fiber.New()

	// Настраиваем обработчик по пути "/about"
	webApp.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("The best school for Software Engineers")
	})

	// Настраиваем обработчик по пути "/courses"
	webApp.Get("/courses", func(c *fiber.Ctx) error {
		return c.SendString("Java, Go, Python")
	})

	logrus.Fatal(webApp.Listen(":80"))
}
