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
