package services

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"rest_api.go/dal"
	"rest_api.go/types"
)

var validate = validator.New()

func GetById(c *fiber.Ctx) error {
	todoId := c.Params("id")

	d := types.TodoResponseDTO{}

	result := dal.GetTodoById(&d, todoId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"message": "Todo not found",
			})
		}

		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to get todo",
		})

	}

	return c.JSON(d)
}

func GetAllTodo(c *fiber.Ctx) error {
	todos := []types.TodoResponseDTO{}

	result := dal.GetAllTodo(&todos)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to get todos",
		})
	}

	return c.JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
	t := new(types.TodoCreateDTO)

	err := c.BodyParser(t)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	if err := validate.Struct(t); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	newTodo := dal.Todo{
		Title: t.Title,
	}

	result := dal.CreateTodo(&newTodo)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to create todo",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Todo Created succesfully",
	})

}
