package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type (
	GetTaskResponse struct {
		ID       int64  `json:"id"`
		Desc     string `json:"description"`
		Deadline int64  `json:"deadline"`
	}

	CreateTaskRequest struct {
		Desc     string `json:"description"`
		Deadline int64  `json:"deadline"`
	}

	CreateTaskResponse struct {
		ID int64 `json:"id"`
	}

	UpdateTaskRequest struct {
		Desc     string `json:"description"`
		Deadline int64  `json:"deadline"`
	}

	Task struct {
		ID       int64
		Desc     string
		Deadline int64
	}
)

var (
	taskIDCounter int64 = 1
	tasks               = make(map[int64]Task)
)

func main() {
	webApp := fiber.New(fiber.Config{
		ReadBufferSize: 16 * 1024})
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	// BEGIN (write your solution here) (write your solution here)
	webApp.Post("/tasks", func(c *fiber.Ctx) error {
		var req CreateTaskRequest
		if err := c.BodyParser(&req); err != nil {
			return c.SendStatus(400)
		}
		newId := taskIDCounter
		taskIDCounter += 1
		task := Task{
			ID:       newId,
			Desc:     req.Desc,
			Deadline: req.Deadline,
		}
		tasks[newId] = task
		return c.JSON(CreateTaskResponse{ID: newId})
	})
	webApp.Get("/tasks/:id", func(c *fiber.Ctx) error {
		param, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Error params")
		}
		task, ok := tasks[param]
		if !ok {
			return c.Status(fiber.StatusNotFound).SendString("Not Found")
		}

		return c.JSON(GetTaskResponse{ID: task.ID, Desc: task.Desc, Deadline: task.Deadline})
	})
	webApp.Patch("/tasks/:id", func(c *fiber.Ctx) error {
		var req UpdateTaskRequest
		if err := c.BodyParser(&req); err != nil {
			return c.SendStatus(400)
		}

		taskId, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Error params")
		}
		task, ok := tasks[taskId]
		if !ok {
			return c.Status(fiber.StatusNotFound).SendString("Not Found")
		}

		if req.Desc != "" {
			task.Desc = req.Desc
		}
		if req.Deadline != 0 {
			task.Deadline = req.Deadline
		}

		tasks[task.ID] = task
		return c.SendStatus(200)
	})
	webApp.Delete("/tasks/:id", func(c *fiber.Ctx) error {
		taskId, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Error params")
		}
		_, ok := tasks[taskId]
		if !ok {
			return c.Status(fiber.StatusNotFound).SendString("Not Found")
		}

		delete(tasks, taskId)
		return c.SendStatus(200)
	})
	// END

	logrus.Fatal(webApp.Listen(":8080"))
}
