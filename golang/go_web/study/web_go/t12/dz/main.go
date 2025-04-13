package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
)

type (
	SendPushNotificationRequest struct {
		Message string `json:"message"`
		UserID  int64  `json:"user_id"`
	}

	PushNotification struct {
		Message string `json:"message"`
		UserID  int64  `json:"user_id"`
	}
)

var pushNotificationsQueue []PushNotification

func main() {
	// BEGIN (write your solution here)
	webApp := fiber.New(fiber.Config{
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})
	webApp.Use(recover.New())

	// END
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	webApp.Post("/push/send", func(c *fiber.Ctx) error {
		var req SendPushNotificationRequest
		if err := c.BodyParser(&req); err != nil {
			// BEGIN (write your solution here)
			return c.Status(400).SendString("Invalid JSON")
			// END
		}

		pushNotificationsQueue = append(pushNotificationsQueue, PushNotification{
			Message: req.Message,
			UserID:  req.UserID,
		})
		if len(pushNotificationsQueue) > 3 {
			panic("Queue is full")
		}

		return c.SendStatus(fiber.StatusOK)
	})

	logrus.Fatal(webApp.Listen(":8080"))
}
