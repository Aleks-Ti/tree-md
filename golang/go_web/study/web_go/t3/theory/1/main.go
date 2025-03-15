package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

const profileUnknown = "unknown"

func main() {
	webApp := fiber.New()
	webApp.Get("/profiles", func(c *fiber.Ctx) error {
		profileID := c.Query("profile_id", profileUnknown)
		if profileID == "" {
			profileID = profileUnknown
		}

		if profileID == profileUnknown {
			return c.Status(fiber.StatusUnprocessableEntity).SendString("profile_id is required")
		}

		return c.SendString(fmt.Sprintf("User Profile ID: %s", profileID))
	})

	logrus.Fatal(webApp.Listen(":80"))
}
