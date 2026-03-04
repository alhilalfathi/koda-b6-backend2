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

// get all product
func (r *ProductRepository) GetAllProduct() *[]models.Product {
	return r.p
}

// get product by id
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

// delete product
func (r *ProductRepository) Delete(id string) error {
	data := *r.p
	strId, _ := strconv.Atoi(id)
	// data = append(data[:index], data[index+1:]...)
	for i, v := range data {
		if v.Id == strId {
			data = append(data[:i], data[i+1:]...)
		}
	}
	return nil
}
