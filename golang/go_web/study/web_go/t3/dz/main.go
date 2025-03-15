package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var exchangeRate = map[string]float64{
	"USD/EUR": 0.8,
	"EUR/USD": 1.25,
	"USD/GBP": 0.7,
	"GBP/USD": 1.43,
	"USD/JPY": 110,
	"JPY/USD": 0.0091,
}

func main() {
	// BEGIN (write your solution here)
	webApp := fiber.New()
	webApp.Get("/convert", func(c *fiber.Ctx) error {
		from := c.Query("from", "null")
		if from == "null" {
			return c.Status(fiber.StatusUnprocessableEntity).SendString("`from` mandatory parameter for the request")
		}
		to := c.Query("to", "null")
		if to == "null" {
			return c.Status(fiber.StatusUnprocessableEntity).SendString("`to` mandatory parameter for the request")
		}
		pair := from + "/" + to
		if value, ok := exchangeRate[pair]; ok {
			return c.Status(fiber.StatusOK).SendString(fmt.Sprintf("%.2f", value))
		} else {
			return c.Status(fiber.StatusNotFound).SendString("404 Not Found")
		}
	})
	// END
	logrus.Fatal(webApp.Listen(":8080"))
}
