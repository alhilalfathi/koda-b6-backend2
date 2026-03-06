package repository

import (
	"context"
	"koda-b6-backend2/internal/models"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn
	// db *[]models.User
}

func NewUserRepository(d *pgx.Conn) *UserRepository {
	return &UserRepository{
		db: d,
	}
}

// get all user
//
//	func (r *UserRepository) GetAllUser() *pgx.Conn {
//		return r.db
//	}
func (r *UserRepository) GetAllUser() ([]models.User, error) {

	rows, err := r.db.Query(context.Background(), `SELECT id,email,password FROM "USER"`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.Id,
			&user.Email,
			&user.Password,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// get user by email
//
//	func (r *UserRepository) GetByEmail(email string) *pgx.Conn{
//		users := *r.db
//		for i := range users {
//			if users[i].Email == email {
//				return &users[i]
//			}
//		}
//		return &models.User{}
//	}
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {

	var user models.User

	query := `SELECT id, email, password FROM "USER" WHERE email=$1`

	err := r.db.QueryRow(context.Background(), query, email).
		Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// create user
//
//	func (r *UserRepository) Create(user models.User) {
//		*r.db = append(*r.db, user)
//	}
func (r *UserRepository) Create(user models.User) error {

	query := `INSERT INTO "USER" (email, password) VALUES ($1,$2)`

	_, err := r.db.Exec(context.Background(), query, user.Email, user.Password)

	return err
}

// update user
// func (r *UserRepository) Update(email string, user *models.User) (*models.User, error) {
// 	if r.db == nil {
// 		return nil, errors.New("no data")
// 	}

// 	for i, v := range *r.db {
// 		if v.Email == email {
// 			(*r.db)[i].Password = user.Password

//				return &(*r.db)[i], nil
//			}
//		}
//		return nil, errors.New("User not found")
//	}
func (r *UserRepository) Update(email string, user *models.User) (*models.User, error) {

	query := `UPDATE "USER"	SET password=$1	WHERE email=$2	RETURNING id, email, password`

	var updatedUser models.User

	err := r.db.QueryRow(context.Background(), query, user.Password, email).
		Scan(&updatedUser.Id, &updatedUser.Email, &updatedUser.Password)

	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

// delete user
//
//	func (r *UserRepository) Delete(email string) error {
//		data := *r.db
//		// data = append(data[:index], data[index+1:]...)
//		for i, v := range data {
//			if v.Email == email {
//				data = append(data[:i], data[i+1:]...)
//			}
//		}
//		return nil
//	}
func (r *UserRepository) Delete(email string) error {

	query := `DELETE FROM "USER" WHERE email=$1`

	_, err := r.db.Exec(context.Background(), query, email)

	if err != nil {
		return err
	}

	return nil
}
