package main

import (
	"slices"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type (
	BinarySearchRequest struct {
		Numbers []int `json:"numbers"`
		Target  int   `json:"target"`
	}

	BinarySearchResponse struct {
		TargetIndex int    `json:"target_index"`
		Error       string `json:"error,omitempty"`
	}
)

const targetNotFound = -1

func main() {
	webApp := fiber.New(fiber.Config{
		ReadBufferSize: 16 * 1024})
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	// BEGIN (write your solution here)
	webApp.Post("/search", func(c *fiber.Ctx) error {
		var request BinarySearchRequest
		var response BinarySearchResponse

		err := c.BodyParser(&request)
		if err != nil {
			response.TargetIndex = targetNotFound
			response.Error = "Invalid JSON"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		result_index := slices.Index(request.Numbers, request.Target)
		if result_index == targetNotFound {
			response.TargetIndex = targetNotFound
			response.Error = "Target was not found"
			return c.Status(fiber.StatusNotFound).JSON(response)
		}

		response.TargetIndex = result_index
		return c.Status(fiber.StatusOK).JSON(response)
	})
	// END

	logrus.Fatal(webApp.Listen(":8080"))
}
