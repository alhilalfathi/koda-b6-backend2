package handlers

import (
	"koda-b6-backend2/internal/models"
	"koda-b6-backend2/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(sr *service.UserService) *UserHandler {
	return &UserHandler{
		service: sr,
	}
}

// get all user
func (h *UserHandler) GetAll(ctx *gin.Context) {
	users := h.service.GetAll()
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "List of all users",
		"results": users,
	})
}

// get user by email
func (h *UserHandler) GetByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	user := h.service.GetByEmail(email)

	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "User not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User found",
		"results": user,
	})
}

// create user
func (h *UserHandler) Create(ctx *gin.Context) {
	var newUser models.CreateUserRequest

	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Create user failed",
		})
		return
	}
	h.service.Create(&newUser)
	ctx.JSON(http.StatusBadRequest, gin.H{
		"success": true,
		"message": "Create user successfuly",
	})
}
