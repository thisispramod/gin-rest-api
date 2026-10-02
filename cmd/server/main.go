package main

import (
	"github.com/gin-gonic/gin"

	"gin-rest-api/internal/handler"
	"gin-rest-api/internal/repository"
	"gin-rest-api/internal/routes"
	"gin-rest-api/internal/service"
)

func main() {
	router := gin.Default()

	productRepository := repository.NewProductRepository()

	productService := service.NewProductService(
		productRepository,
	)

	productHandler := handler.NewProductHandler(
		productService,
	)

	routes.SetupRoutes(
		router,
		productHandler,
	)

	router.Run(":8080")
}
