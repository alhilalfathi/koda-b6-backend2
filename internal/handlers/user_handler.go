package handlers

import (
	"koda-b6-backend2/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHanler(sr *service.UserService) *UserHandler {
	return &UserHandler{
		service: sr,
	}
}

func (h *UserHandler) GetAll(ctx *gin.Context) {
	users := h.service.GetAll()
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "List of all users",
		"results": users,
	})
}
