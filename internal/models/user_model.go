package models

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UpdateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Product struct {
	ProductName string `json:"product-name"`
	Desc        string `json:"product-desc"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
}
type CreateProductRequest struct {
	ProductName string `json:"product-name"`
	Desc        string `json:"product-desc"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
}
type UpdateProductRequest struct {
	ProductName string `json:"product-name"`
	Desc        string `json:"product-desc"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
}
