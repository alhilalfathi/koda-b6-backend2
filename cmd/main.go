package main

import (
	"koda-b6-backend2/internal/di"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	container := di.NewContainer()

	userHandler := container.UserHandler()

	users := r.Group("/users")
	{
		users.GET("", userHandler.GetAll)
		users.GET("/:email", userHandler.GetByEmail)
		users.POST("", userHandler.Create)
	}

	products := r.Group("/products")
	{
		products.GET("")
	}

	r.Run("localhost:8888")
}
