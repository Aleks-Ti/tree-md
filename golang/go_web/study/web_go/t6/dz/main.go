package main

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type (
	CreateLinkRequest struct {
		External string `json:"external"`
		Internal string `json:"internal"`
	}

	GetLinkResponse struct {
		Internal string `json:"internal"`
	}
)

var links = make(map[string]string)

func main() {
	webApp := fiber.New(fiber.Config{
		ReadBufferSize: 16 * 1024})
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})
	// BEGIN (write your solution here)
	webApp.Post("/links", func(c *fiber.Ctx) error {
		var req CreateLinkRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
		}
		if req.External == "" || req.Internal == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
		}
		links[req.External] = req.Internal

		return c.SendStatus(200)
	})
	webApp.Get("/links/:external", func(c *fiber.Ctx) error {
		var externalLink = c.Params("external")
		if externalLink == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Not found param for request")
		}

		decodeExternalLink, err := url.QueryUnescape(externalLink)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).SendString("Invalid link")
		}

		internalLink, ok := links[decodeExternalLink]
		if !ok {
			return c.Status(fiber.StatusNotFound).SendString("Link not found")
		}
		response := GetLinkResponse{
			Internal: internalLink,
		}
		return c.JSON(response)
	})
	// END

	logrus.Fatal(webApp.Listen(":8080"))
}
