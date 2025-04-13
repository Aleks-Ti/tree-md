package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/sirupsen/logrus"
)

// Структура с информацией о фильме
type Film struct {
	Title    string
	IsViewed bool
}

// Для простоты описываем хранилище фильмов в коде
var films = []Film{
	{
		Title:    "The Shawshank Redemption",
		IsViewed: true,
	},
	{
		Title:    "The Godfather",
		IsViewed: true,
	},
	{
		Title:    "The Godfather: Part II",
		IsViewed: false,
	},
}

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

	// Настраиваем обработчик для веб-страницы со списком фильмов
	webApp.Get("/films", func(c *fiber.Ctx) error {
		return c.Render("film-list", films)
	})

	logrus.Fatal(webApp.Listen(":80"))
}
