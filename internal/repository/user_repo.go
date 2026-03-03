package repository

import "koda-b6-backend2/internal/models"

type UserRepository struct {
	db *[]models.User
}

func NewUserRepository(d *[]models.User) *UserRepository {
	return &UserRepository{
		db: d,
	}
}

func (r *UserRepository) GetAllUser() *[]models.User {
	return r.db
}
