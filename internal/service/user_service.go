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

func (s *UserService) GetAll() ([]models.User, error) {
	users, err := s.repo.GetAllUser()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) GetByEmail(email string) (*models.User, error) {

	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Create(req *models.CreateUserRequest) error {

	if req.Email == "" || req.Password == "" {
		return errors.New("Email and password required")
	}

	newUser := models.User{
		Email:    req.Email,
		Password: req.Password,
	}

	err := s.repo.Create(newUser)
	if err != nil {
		return fmt.Errorf("Failed to create user: %w", err)
	}

	return nil
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
		return nil, fmt.Errorf("Failed to update user: %w", err)
	}

	return updatedUser, nil
}

func (s *UserService) Delete(email string) error {

	_, err := s.repo.GetByEmail(email)
	if err != nil {
		return errors.New("User not found")
	}

	err = s.repo.Delete(email)
	if err != nil {
		return fmt.Errorf("Failed to delete user: %w", err)
	}

	return nil
}
