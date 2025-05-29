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

	app.Get("/todo", services.GetAllTodo)
	app.Get("/todo/:id", services.GetById)
	app.Post("/todo", services.CreateTodo)

	app.Listen("localhost:8080")
}
