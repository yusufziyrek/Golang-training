package main

import (
	"github.com/gofiber/fiber/v2"
	"rest_api.go/dal"
	"rest_api.go/database"
	"rest_api.go/services"
)

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

	app.Post("/todo", services.CreateTodo)

	app.Listen("localhost:8080")
}
