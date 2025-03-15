package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var counters = make(map[string]int64)

const requestParamKeyEvent = "event"

func main() {
	webApp := fiber.New()

	webApp.Get("/counter/:event", func(c *fiber.Ctx) error {
		event := c.Params(requestParamKeyEvent, "")
		if event == "" {
			return c.SendStatus(fiber.StatusUnprocessableEntity)
		}

		eventCounter, ok := counters[event]
		if !ok {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.SendString(strconv.FormatInt(eventCounter, 10))
	})

	webApp.Post("/counter/:event", func(c *fiber.Ctx) error {
		event := c.Params(requestParamKeyEvent, "")
		if event == "" {
			return c.SendStatus(fiber.StatusUnprocessableEntity)
		}

		counters[event] += 1

		return c.SendStatus(fiber.StatusOK)
	})

	logrus.Fatal(webApp.Listen(":80"))
}
