package main

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Создание сотрудника
type (
	ListEmployeesResponse struct {
		Employees []EmployeePayload `json:"employees"`
	}

	GetEmployeeResponse struct {
		EmployeePayload
	}

	EmployeePayload struct {
		ID    string `json:"id"`
		Email string `json:"email"`
		Role  string `json:"role"`
	}
)

type (
	CreateEmployeeRequest struct {
		Email string `json:"email"`
		Role  string `json:"role"`
	}

	CreateEmployeeResponse struct {
		ID string `json:"id"`
	}
)
type (
	UpdateEmployeeRequest struct {
		Email string `json:"email"`
		Role  string `json:"role"`
	}
)

// Хранилище
type (
	Employee struct {
		ID    string
		Email string
		Role  string
	}

	EmployeeStorageInMemory struct {
		employees map[string]Employee
	}
)

func (s *EmployeeStorageInMemory) Create(empl Employee) (string, error) {
	// Генерируем ID для сотрудника
	empl.ID = uuid.New().String()

	s.employees[empl.ID] = empl

	return empl.ID, nil
}

func (s *EmployeeStorageInMemory) List() []Employee {
	// Инициализируем массив с размером равным количеству
	// всех сотрудников в хранилище
	employees := make([]Employee, 0, len(s.employees))

	for _, empl := range s.employees {
		employees = append(employees, empl)
	}

	return employees
}

func (s *EmployeeStorageInMemory) Get(id string) (Employee, error) {
	empl, ok := s.employees[id]
	if !ok {
		// Возвращаем ошибку, если сотрудника с таким
		// идентификатором не существует
		return Employee{}, errors.New("employee not found")
	}

	return empl, nil
}

func (s *EmployeeStorageInMemory) Update(id, email, role string) error {
	empl, ok := s.employees[id]
	if !ok {
		// Возвращаем ошибку, если сотрудника с таким
		// идентификатором не существует
		return errors.New("employee not found")
	}

	// Обновляем электронную почту сотрудника,
	// если новое значение было передано
	if email != "" {
		empl.Email = email
	}
	// Обновляем роль сотрудника,
	// если новое значение было передано
	if role != "" {
		empl.Role = role
	}

	s.employees[empl.ID] = empl

	return nil
}
func (s *EmployeeStorageInMemory) Delete(id string) {
	delete(s.employees, id)
}
func main() {
	webApp := fiber.New()

	storage := &EmployeeStorageInMemory{
		employees: make(map[string]Employee),
	}

	webApp.Post("/employees", func(c *fiber.Ctx) error {
		var req CreateEmployeeRequest
		if err := c.BodyParser(&req); err != nil {
			return fmt.Errorf("body parser: %w", err)
		}

		id, err := storage.Create(Employee{
			Email: req.Email,
			Role:  req.Role,
		})
		if err != nil {
			return fmt.Errorf("create in storage: %w", err)
		}

		return c.JSON(CreateEmployeeResponse{ID: id})
	})

	webApp.Get("/employees", func(c *fiber.Ctx) error {
		// Получаем список всех сотрудников из хранилища
		employees := storage.List()

		// Формируем ответ
		resp := ListEmployeesResponse{
			Employees: make([]EmployeePayload, len(employees)),
		}
		for i, empl := range employees {
			resp.Employees[i] = EmployeePayload(empl)
		}

		// Возвращаем список сотрудников JSON-строкой в теле ответа
		return c.JSON(resp)
	})

	// Получение одного сотрудника
	webApp.Get("/employees/:id", func(c *fiber.Ctx) error {
		empl, err := storage.Get(c.Params("id"))
		if err != nil {
			return fiber.ErrNotFound
		}

		// Возвращаем данные сотрудника JSON-строкой в теле ответа
		return c.JSON(GetEmployeeResponse{EmployeePayload(empl)})
	})
	webApp.Patch("/employees/:id", func(c *fiber.Ctx) error {
		// Парсим JSON-тело запроса в объект UpdateEmployeeRequest
		var req UpdateEmployeeRequest
		if err := c.BodyParser(&req); err != nil {
			return fmt.Errorf("body parser: %w", err)
		}

		// Обновляем данные сотрудника в хранилище. Эта функция может вернуть ошибку,
		// если сотрудника с таким идентификатором не существует.
		err := storage.Update(c.Params("id"), req.Email, req.Role)
		if err != nil {
			return fmt.Errorf("update: %w", err)
		}

		return nil
	})
	webApp.Delete("/employees/:id", func(c *fiber.Ctx) error {
		storage.Delete(c.Params("id"))

		// Возвращаем успешный ответ без тела
		return c.SendStatus(fiber.StatusNoContent)
	})
	logrus.Fatal(webApp.Listen(":80"))
}
