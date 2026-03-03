package service

import (
	"errors"
	"fmt"
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

func (s *UserService) Update(email string, u *models.UpdateUserRequest) (*models.User, error) {
	if u.Password == "" {
		return nil, errors.New("Password cannot blank")
	}

	user := &models.User{
		Password: u.Password,
	}

	updatedUser, err := s.repo.Update(email, user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return updatedUser, nil
}
