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

// get all user
func (r *UserRepository) GetAllUser() *[]models.User {
	return r.db
}

// get user by email
func (r *UserRepository) GetByEmail(email string) *models.User {
	users := *r.db
	for i := range users {
		if users[i].Email == email {
			return &users[i]
		}
	}
	return &models.User{}
}
