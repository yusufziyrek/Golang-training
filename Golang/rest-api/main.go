package main

import (
	"github.com/gofiber/fiber/v2"
	"rest_api.go/dal"
	"rest_api.go/database"
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
		Title string
	}

	app.Post("/todo", func(c *fiber.Ctx) error {
		t := new(TodoCreate)

		err := c.BodyParser(t)
		if err != nil {
			return err
		}

		newTodo := dal.Todo{
			Title: t.Title,
		}

		result := database.DB.Create(&newTodo)

		if result.Error != nil {
			return result.Error
		}

		return c.JSON(fiber.Map{
			"message": "Created succesfully",
		})

	})

	app.Listen("localhost:8080")
}
