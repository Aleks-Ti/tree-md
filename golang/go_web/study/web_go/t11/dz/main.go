package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/sirupsen/logrus"
)

type (
	CreateItemRequest struct {
		Name  string `json:"name"`
		Price uint   `json:"price"`
	}

	Item struct {
		Name  string `json:"name"`
		Price uint   `json:"price"`
	}

	ItemsHandler struct {
		storage *ItemStorage
	}

	ItemStorage struct {
		item []Item
	}
)

var (
	items []Item
)

func main() {
	viewsEngine := html.New("./templates", ".tmpl")
	webApp := fiber.New(fiber.Config{
		Views:          viewsEngine,
		ReadBufferSize: 16 * 1024,
	})
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})
	// BEGIN (write your solution here)
	itemStorage := &ItemStorage{item: items}
	itemHandler := &ItemsHandler{storage: itemStorage}
	webApp.Post("/items", itemHandler.AddItem)
	webApp.Get("/items/view", itemHandler.GetItems)
	// END

	logrus.Fatal(webApp.Listen(":8080"))
}

func (h *ItemsHandler) AddItem(c *fiber.Ctx) error {
	req := CreateItemRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.SendStatus(422)
	}

	newData := Item{
		Name:  req.Name,
		Price: req.Price,
	}

	h.storage.item = append(h.storage.item, newData)
	return c.SendStatus(200)
}

func (h *ItemsHandler) GetItems(c *fiber.Ctx) error {
	for _, item := range h.storage.item {
		println(item.Name, item.Price)
	}
	err := c.Render("items", h.storage.item)
	if err != nil {
		fmt.Println(err)
		c.SendStatus(400)
	}
	return c.SendStatus(200)
}
