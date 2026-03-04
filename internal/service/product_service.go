package service

import (
	"koda-b6-backend2/internal/models"
	"koda-b6-backend2/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(rp *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: rp,
	}
}
func (s *ProductService) GetAll() []models.Product {
	return *s.repo.GetAllProduct()
}
func (s *ProductService) GetProductById(id string) *models.Product {
	return s.repo.GetProductById(id)
}
