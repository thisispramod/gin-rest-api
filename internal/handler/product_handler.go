package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"gin-rest-api/internal/model"
	"gin-rest-api/internal/service"
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
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product model.Product

	err := c.ShouldBindJSON(&product)

	if err != nil {

		if err != nil {
			var validationErrors validator.ValidationErrors

			if errors.As(err, &validationErrors) {
				errorMessages := make(map[string]string)

				for _, fieldError := range validationErrors {
					fieldName := fieldError.Field()

					switch fieldError.Tag() {
					case "required":
						errorMessages[fieldName] = fieldName + " is required"

					case "min":
						errorMessages[fieldName] = fieldName + " must be at least 2 characters"

					case "gt":
						errorMessages[fieldName] = fieldName + " must be greater than 0"
					}
				}

				c.JSON(http.StatusBadRequest, gin.H{
					"error":  "validation failed",
					"fields": errorMessages,
				})
				return
			}
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	createProduct := h.service.CreateProduct(product)

	c.JSON(http.StatusCreated, gin.H{
		"data": createProduct,
	})
}
