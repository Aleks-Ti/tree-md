package main

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Запросы и ответы
type (
	CreateOrderRequest struct {
		UserID     int64   `json:"user_id"`
		ProductIDs []int64 `json:"product_ids"`
	}

	CreateOrderResponse struct {
		ID string `json:"id"`
	}

	GetOrderResponse struct {
		ID         string  `json:"id"`
		UserID     int64   `json:"user_id"`
		ProductIDs []int64 `json:"product_ids"`
	}
)

func main() {
	orderHandler := &OrderHandler{
		storage: &OrderStorage{
			orders: make(map[string]Order),
		},
	}

	webApp := fiber.New()
	webApp.Post("/orders", orderHandler.CreateOrder)
	webApp.Get("/orders/:id", orderHandler.GetOrder)

	logrus.Fatal(webApp.Listen(":80"))
}

// Абстрактное хранилище
type OrderCreatorGetter interface {
	CreateOrder(order Order) (string, error)
	GetOrder(id string) (Order, error)
}

// Обработчик
type OrderHandler struct {
	storage OrderCreatorGetter
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var request CreateOrderRequest
	if err := c.BodyParser(&request); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}

	order := Order{
		ID:         uuid.New().String(),
		UserID:     request.UserID,
		ProductIDs: request.ProductIDs,
	}

	id, err := h.storage.CreateOrder(order)
	if err != nil {
		return fmt.Errorf("create order: %w", err)
	}

	return c.JSON(CreateOrderResponse{
		ID: id,
	})
}

func (h *OrderHandler) GetOrder(c *fiber.Ctx) error {
	id := c.Params("id")

	order, err := h.storage.GetOrder(id)
	if err != nil {
		return fmt.Errorf("get order: %w", err)
	}

	return c.JSON(GetOrderResponse(order))
}

// Модель заказа
type Order struct {
	ID         string
	UserID     int64
	ProductIDs []int64
}

// Хранилище
type OrderStorage struct {
	orders map[string]Order
}

func (o *OrderStorage) CreateOrder(order Order) (string, error) {
	o.orders[order.ID] = order

	return order.ID, nil
}

// Ошибки
var (
	errOrderNotFound = errors.New("order not found")
)

func (o *OrderStorage) GetOrder(id string) (Order, error) {
	order, ok := o.orders[id]
	if !ok {
		return Order{}, errOrderNotFound
	}

	return order, nil
}
