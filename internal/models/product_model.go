package models

type Product struct {
	Id          int    `json:"product-id"`
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
