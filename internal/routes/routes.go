package routes

import (
	"gin-rest-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	productHandler *handler.ProductHandler,
) {
	api := router.Group("/api/v1") // common prefix

	api.GET("/products", productHandler.GetProducts)

	api.GET("/products/:id", productHandler.GetProductByID)

	api.POST("/products", productHandler.CreateProduct)

}
