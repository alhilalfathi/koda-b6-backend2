package main

import (
	"context"
	"fmt"
	"koda-b6-backend2/internal/di"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	connConfig, _ := pgx.ParseConfig("")

	conn, err := pgx.Connect(context.Background(), connConfig.ConnString())

	if err != nil {
		fmt.Println("Failed to connecting db")
	}

	r := gin.Default()

	container := di.NewContainer(conn)

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
