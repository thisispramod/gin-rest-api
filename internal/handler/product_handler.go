package handler

import (
	"gin-rest-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(
	service *service.ProductService,
) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products := h.service.GetAllProducts()

	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})

}
