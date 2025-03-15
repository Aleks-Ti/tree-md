package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type (
	CreateLogEntryRequest struct {
		Message   string `json:"message"`
		Level     string `json:"level"`
		Timestamp int64  `json:"timestamp"`
	}

	CreateLogEntryResponse struct {
		ID string `json:"id"`
	}

	LogEntry struct {
		ID        string
		Message   string
		Level     string
		Timestamp int64
	}
)

var logs []LogEntry

func main() {
	webApp := fiber.New()

	webApp.Post("/logs", func(c *fiber.Ctx) error {
		var request CreateLogEntryRequest
		if err := c.BodyParser(&request); err != nil {
			return fmt.Errorf("body parser: %w", err)
		}

		logEntry := LogEntry{
			ID:        uuid.New().String(),
			Message:   request.Message,
			Level:     request.Level,
			Timestamp: request.Timestamp,
		}

		// Упрощенное хранение в памяти приложения
		logs = append(logs, logEntry)

		return c.JSON(CreateLogEntryResponse{
			ID: logEntry.ID,
		})
	})

	logrus.Fatal(webApp.Listen(":80"))
}
