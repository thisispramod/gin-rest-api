package handler

import (
	"gin-rest-api/internal/service"
	"net/http"
	"strconv"

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

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid product id",
		})
	}

	product, found := h.service.GetProductById(id)

	if !found {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product Not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}
