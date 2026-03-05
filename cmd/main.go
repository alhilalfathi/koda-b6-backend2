package main

import (
	"koda-b6-backend2/internal/di"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	container := di.NewContainer()

	userHandler := container.UserHandler()
	productHandler := container.ProductHandler()

	users := r.Group("/users")
	{
		users.GET("", userHandler.GetAll)
		users.GET("/:email", userHandler.GetByEmail)
		users.POST("", userHandler.Create)
		users.PUT("/:email", userHandler.Update)
		users.DELETE("/:email", userHandler.Delete)
	}

	products := r.Group("/products")
	{
		products.GET("", productHandler.GetAllProduct)
		products.GET("/:id", productHandler.GetProductById)
		products.POST("", productHandler.CreateProduct)
		products.PUT("/:id", productHandler.UpdateProduct)
		products.DELETE("", productHandler.Delete)
	}

	r.Run("localhost:8888")
}
