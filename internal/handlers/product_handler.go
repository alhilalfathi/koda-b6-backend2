package handlers

import (
	"koda-b6-backend2/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(sr *service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: sr,
	}
}

func (h *ProductHandler) GetAll(ctx *gin.Context) {
	products := h.service.GetAll()
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "List of all products",
		"results": products,
	})
}

func (h *ProductHandler) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	data := h.service.GetProductById(id)

	if data == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Product not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Product found",
		"results": data,
	})
}

func (h *ProductHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"success": true,
		"message": "Product deleted",
	})
}
