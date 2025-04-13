package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/sirupsen/logrus"
)

func main() {
	// Инициализируем стандартный Go-шаблонизатор
	// Указываем папку, в которой хранятся шаблоны
	// Указываем расширение файлов шаблонов
	viewsEngine := html.New("./template", ".tmpl")

	// Создаем новый экземпляр Fiber веб-приложения,
	// указывая наш шаблонизатор
	webApp := fiber.New(fiber.Config{
		Views: viewsEngine,
	})

	// Настраиваем обработчик для веб-страницы аккаунта пользователя
	webApp.Get("/profile", func(c *fiber.Ctx) error {
		return c.Render("profile", fiber.Map{
			"name":  "John",
			"email": "john@doe.com",
		})
	})

	logrus.Fatal(webApp.Listen(":80"))
}
