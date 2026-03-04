package repository

import (
	"koda-b6-backend2/internal/models"
	"strconv"
)

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

func (r *ProductRepository) GetProductById(id string) *models.Product {
	data := *r.p
	strId, _ := strconv.Atoi(id)
	for i := range data {
		if data[i].Id == strId {
			return &data[i]
		}
	}
	return &models.Product{}
}
