package storage

import (
	"fmt"
	"log"

	"github.com/dobrbi2000/rest_api_films/models"
)

type UserRepository struct {
	storage *Storage
}

var (
	tableUser string = "users"
)

// create user
func (ur *UserRepository) Create(u *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (login, password) VALUES ($1, $2) RETURNING user_id", tableUser)
	if err := ur.storage.db.QueryRow(query, u.Login, u.Password).Scan(&u.ID); err != nil {
		return nil, err
	}
	fmt.Println(u)
	return u, nil

}

// find
func (ur *UserRepository) FindByLogin(login string) (*models.User, bool, error) {
	users, err := ur.SelectAll()
	var founded bool
	if err != nil {
		return nil, founded, err
	}
	var useFinded *models.User
	for _, u := range users {
		if u.Login == login {
			useFinded = u
			founded = true
			break
		}
	}

	return useFinded, founded, nil
}

// select all users
func (ur *UserRepository) SelectAll() ([]*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableUser)
	rows, err := ur.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		u := models.User{}
		err := rows.Scan(&u.ID, &u.Login, &u.Password)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, &u)
	}

	return users, nil
}
