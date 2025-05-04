package main

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"rest_api.go/dal"
	"rest_api.go/database"
)

var validate = validator.New()

func main() {

	database.Connect()
	database.DB.AutoMigrate(&dal.Todo{})

	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"message": "Hello World",
		})
	})

	type TodoCreate struct {
		Title string `validate:"required"`
	}

	app.Post("/todo", func(c *fiber.Ctx) error {
		t := new(TodoCreate)

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

		result := database.DB.Create(&newTodo)

		if result.Error != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Failed to create",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Created succesfully",
		})

	})

	app.Listen("localhost:8080")
}
