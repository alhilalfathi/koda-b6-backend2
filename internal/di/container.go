package di

import (
	"koda-b6-backend2/internal/handlers"
	"koda-b6-backend2/internal/models"
	"koda-b6-backend2/internal/repository"
	"koda-b6-backend2/internal/service"

	"github.com/jackc/pgx/v5"
)

type Container struct {
	user        *pgx.Conn
	userRepo    *repository.UserRepository
	userService *service.UserService
	userHandler *handlers.UserHandler

	product        *[]models.Product
	productRepo    *repository.ProductRepository
	productService *service.ProductService
	productHandler *handlers.ProductHandler
}

func NewContainer(db *pgx.Conn) *Container {

	var DataProduct []models.Product

	container := Container{
		user:    db,
		product: &DataProduct,
	}

	container.initDependencies()

	return &container
}

func (c *Container) initDependencies() {
	c.userRepo = repository.NewUserRepository(c.user)
	c.userService = service.NewUserService(c.userRepo)
	c.userHandler = handlers.NewUserHandler(c.userService)

	c.productRepo = repository.NewProductRepository(c.product)
	c.productService = service.NewProductService(c.productRepo)
	c.productHandler = handlers.NewProductHandler(c.productService)
}

func (c *Container) UserHandler() *handlers.UserHandler {
	return c.userHandler
}
func (c *Container) ProductHandler() *handlers.ProductHandler {
	return c.productHandler
}
