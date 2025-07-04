package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"product-app/common/app"
	"product-app/common/postgresql"
	"product-app/controller"
	"product-app/repository"
	"product-app/service"
)

func main() {
	ctx := context.Background()
	e := echo.New()

	configurationManager := app.NewConfigurationManager()

	dbPoll := postgresql.GetConnectionPool(ctx, configurationManager.PostgreSqlConfig)
	productRepository := repository.NewProductRepository(dbPoll)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	productController.RegisterRoutes(e)
	e.Start("localhost:8080")

}
