package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"api/product/models"
	"api/product/repositories"
	"api/product/services"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) Create(c *gin.Context) {
	var p models.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	if err := h.service.Create(c.Request.Context(), &p); err != nil {
		if errors.Is(err, repositories.ErrDuplicateName) {
			c.JSON(http.StatusConflict, gin.H{"error": "Product with the same name already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating product"})
		return
	}

	c.JSON(http.StatusCreated, p)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving product"})
		return
	}

	c.JSON(http.StatusOK, product)
}
