package service

import (
	"koda-b6-backend2/internal/models"
	"koda-b6-backend2/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(rp *repository.UserRepository) *UserService {
	return &UserService{
		repo: rp,
	}
}

func (s *UserService) GetAll() []models.User {
	return *s.repo.GetAllUser()
}

func (s *UserService) GetByEmail(email string) *models.User {
	return s.repo.GetByEmail(email)
}

func (s *UserService) Create(req *models.CreateUserRequest) {
	newUser := models.User{
		Email:    req.Email,
		Password: req.Password,
	}
	s.repo.Create(newUser)
}
