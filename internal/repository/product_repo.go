package repository

import "koda-b6-backend2/internal/models"

type ProductRepository struct {
	p *[]models.Product
}

func NewProductRepository(products *[]models.Product) *ProductRepository {
	return &ProductRepository{
		p: products,
	}
}
func (r *ProductRepository) GetAllProduct() *[]models.Product {
	return r.p
}
