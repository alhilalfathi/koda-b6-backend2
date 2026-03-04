package service

import (
	"errors"
	"fmt"
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

func (s *ProductService) Delete(id string) error {
	user := s.repo.GetProductById(id)

	if user == nil {
		return errors.New("User not found")
	}

	err := s.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}
