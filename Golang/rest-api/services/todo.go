package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"rest_api.go/dal"
	"rest_api.go/types"
)

var validate = validator.New()

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

func GetAllTodo(c *fiber.Ctx) error {
	return c.JSON("")
}
