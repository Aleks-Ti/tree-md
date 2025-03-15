package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	webApp := fiber.New()
	// Обозначаем, что на GET запрос по пути /address нужно вернуть строку с адресом
	webApp.Get("/address", func(c *fiber.Ctx) error {
		return c.SendString("145 DUNDEE SOUTH SAN FRANCISCO CA 94080-1023. USA")
	})

	// Запускаем веб-приложение на порту 80
	// Оборачиваем в функцию логирования, чтобы видеть ошибки, если они возникнут
	logrus.Fatal(webApp.Listen(":80"))
}
