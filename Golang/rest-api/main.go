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

	app.Post("/todo", services.CreateTodo)
	app.Get("/todo", services.GetAllTodo)

	app.Listen("localhost:8080")
}
