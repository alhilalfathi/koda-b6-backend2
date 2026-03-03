package repository

import (
	"errors"
	"koda-b6-backend2/internal/models"
)

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

// create user
func (r *UserRepository) Create(user models.User) {
	*r.db = append(*r.db, user)
}

// update user
func (r *UserRepository) Update(email string, user *models.User) (*models.User, error) {
	if r.db == nil {
		return nil, errors.New("no data")
	}

	for i, v := range *r.db {
		if v.Email == email {
			(*r.db)[i].Password = user.Password

			return &(*r.db)[i], nil
		}
	}
	return nil, errors.New("User not found")
}

// delete user
func (r *UserRepository) Delete(email string) error {
	data := *r.db
	// data = append(data[:index], data[index+1:]...)
	for i, v := range data {
		if v.Email == email {
			data = append(data[:i], data[i+1:]...)
		}
	}
	return nil
}
