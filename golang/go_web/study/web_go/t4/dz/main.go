package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var postLikes = make(map[string]int64)

func main() {
	webApp := fiber.New()

	var patrParametrs string = "post_id"
	webApp.Get("/likes/:post_id", func(c *fiber.Ctx) error {
		param := c.Params(patrParametrs, "null")
		likesCounter, ok := postLikes[param]
		if !ok {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.SendString(strconv.FormatInt(likesCounter, 10))
	})
	webApp.Post("/likes/:post_id", func(c *fiber.Ctx) error {
		param := c.Params(patrParametrs, "null")
		_, ok := postLikes[param]
		if !ok {
			postLikes[param] = 1
			return c.Status(fiber.StatusCreated).SendString("1")
		}
		postLikes[param] += 1
		result := postLikes[param]
		return c.SendString(strconv.FormatInt(result, 10))
	})

	logrus.Fatal(webApp.Listen(":80"))
}
